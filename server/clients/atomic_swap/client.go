package atomicswap

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/web3_proxy/server/clients/nostr"
)

type (
	// Client for atomic swaps
	Client struct {
		nostr  *nostr.Client
		stalls []nostr.Stall
	}
)

const (
	tagTftStall     = "TFT_ATOMIC_SWAP_STALL"
	tagTftSaleOrder = "TFT_ATOMIC_SWAP_SALE_ORDER"
)

var (
	// List of allowed currency strings
	knownCurrencies = map[string]struct{}{"ETH": {}}

	ErrCurrencyNotAllowed = errors.New("currency not allowed")
	ErrNoSalesFound       = errors.New("no sales found for the current currency and price")
)

// NewClient for atomic swaps
func NewClient(ctx context.Context, nostr *nostr.Client) (*Client, error) {
	client := &Client{nostr: nostr}

	if err := client.loadOwnStalls(ctx); err != nil {
		return nil, errors.Wrap(err, "could not fetch owned stalls")
	}

	return client, nil
}

// PlaceSellOrder on nostr relays. A sell order is always for stellar based TFT. The buying currency,
// as well as the price to buy 1 TFT in that currency is specified. Amount is expressed in whole TFT
// (= 10_000_000 stropes of TFT). Price is expressed as the smallest possible unit of the target currency.
func (c *Client) PlaceSellOrder(ctx context.Context, amount uint, currency string, price uint) (*Driver, error) {
	// check if we allow this currency
	if _, allowed := knownCurrencies[currency]; !allowed {
		return nil, ErrCurrencyNotAllowed
	}
	// check if we have a stall
	stallId := ""
	for _, stall := range c.stalls {
		if stall.Currency == currency {
			stallId = stall.Id
			break
		}
	}

	// Create stall first if we don't have one yet
	if stallId == "" {
		stallId = uuid.NewString()
		stall := nostr.Stall{
			Id:       stallId,
			Name:     fmt.Sprintf("TFT_%s_SALE_LISTING", currency),
			Currency: currency,
		}
		c.nostr.PublishStall(ctx, []string{"t", tagTftStall}, stall)
	}
	product := nostr.Product{
		Id:       uuid.NewString(),
		StallId:  stallId,
		Name:     fmt.Sprintf("TFT_%s_SALE_ORDER", currency),
		Currency: currency,
		Price:    float64(price),
		Quantity: amount,
	}

	if err := c.nostr.PublishProduct(ctx, []string{"t", tagTftSaleOrder}, product); err != nil {
		return nil, errors.Wrap(err, "could not publish sale")
	}

	driver := initDriver(c.nostr)
	if err := driver.OpenSale(); err != nil {
		return nil, errors.Wrap(err, "could not start sale driver")
	}

	return initDriver(c.nostr), nil
}

// Attempt to buy from an already existing swap
// TODO: in the future this should be changed to keep listening for new sell orders
func (c *Client) AttemptBuy(ctx context.Context, amount uint, currency string, maxPrice uint) (*Driver, error) {
	// check if we allow this currency
	if _, allowed := knownCurrencies[currency]; !allowed {
		return nil, ErrCurrencyNotAllowed
	}

	openSales, err := c.loadSaleOrders(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not load existing sales")
	}

	filteredSales := []nostr.Product{}
	for _, sale := range openSales {
		if sale.Currency != currency {
			continue
		}

		if sale.Price > float64(maxPrice) {
			continue
		}

		filteredSales = append(filteredSales, sale)
	}

	// filteredSales is a list of all sales in the given currency
	// sort by by price
	sort.Slice(filteredSales, func(i, j int) bool {
		return filteredSales[i].Price < filteredSales[j].Price
	})

	// if we actually have a sale open, attempt to drive it
	if len(filteredSales) > 0 {
		driver := initDriver(c.nostr)
		// TODO
		driver.Buy(filteredSales[0], amount)
		return driver, nil
	}

	return nil, ErrNoSalesFound
}

// load all existing stalls on connected relays
func (c *Client) loadOwnStalls(ctx context.Context) error {
	subId, err := c.nostr.SubscribeStallCreation(tagTftStall)
	if err != nil {
		return errors.Wrap(err, "could not subscribe to stall creation events")
	}

	// Wait to fetch events
	log.Debug().Msg("Waiting to fetch stored events")
	time.Sleep(time.Second * 5)
	events := c.nostr.GetSubscriptionEvents(subId)

	stalls := []nostr.Stall{}
	for _, evt := range events {
		stall := nostr.Stall{}
		// we only care about our own stalls
		if evt.PubKey != c.nostr.PublicKey() {
			continue
		}
		if err := json.Unmarshal([]byte(evt.Content), &stall); err != nil {
			log.Debug().Err(err).Msg("unexpected content in event")
		}
		stalls = append(stalls, stall)
	}

	c.stalls = stalls

	c.nostr.CloseSubscription(subId)

	return nil
}

func (c *Client) loadSaleOrders(ctx context.Context) ([]nostr.Product, error) {
	subId, err := c.nostr.SubscribeProductCreation(tagTftSaleOrder)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to product creation events")
	}

	// Wait to fetch events
	log.Debug().Msg("Waiting to fetch stored sale events")
	time.Sleep(time.Second * 5)
	events := c.nostr.GetSubscriptionEvents(subId)

	sales := []nostr.Product{}
	for _, evt := range events {
		sale := nostr.Product{}
		if err := json.Unmarshal([]byte(evt.Content), &sale); err != nil {
			log.Debug().Err(err).Msg("unexpected content in event")
		}
		sales = append(sales, sale)
	}

	c.nostr.CloseSubscription(subId)

	return sales, nil
}
