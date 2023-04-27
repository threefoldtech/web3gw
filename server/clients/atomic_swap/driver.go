package atomicswap

import (
	"context"
	"crypto/sha256"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	goethclient "github.com/threefoldtech/web3_proxy/server/clients/eth"
	"github.com/threefoldtech/web3_proxy/server/clients/nostr"
	stellargoclient "github.com/threefoldtech/web3_proxy/server/clients/stellar"
)

type (
	// Driver for atomic swaps
	Driver struct {
		nostr   *nostr.Client
		eth     *goethclient.Client
		stellar *stellargoclient.Client

		saleId string
		swapId string

		stage DriverStage

		// amount of TFT to swap, this is initialized in a sell order to the maximum available
		swapAmount uint
		// amount of other token to pay
		swapPrice uint

		msges <-chan nostr.NostrEvent
	}

	DriverStage = int

	MsgBuy struct {
		Id     string `json:"id"`
		Amount uint   `json:"amount"`
	}

	MsgAccept struct {
		EthAddress     common.Address `json:"ethAddress"`
		StellarAddress string         `json:"stellarAddress"`
	}

	MsgInitiateEth struct {
		SharedSecret   [sha256.Size]byte `json:"sharedSecret"`
		EthAddress     common.Address    `json:"ethAddress"`
		StellarAddress string            `json:"stellarAddress"`
	}

	MsgParticipateStellar struct {
		HoldingAccount string `json:"holdingAccount`
		RefundTx       string `json:"refundTx"`
	}

	MsgRedeemed struct {
		Secret [32]byte `json:"secret"`
	}
)

const (
	DriverStageOpenSale DriverStage = iota
	DriverStageStartBuy
	DriverStageAcceptedBuy
	DriverStageSetupSwap
	DriverStageParticipateSwap
	DriverStageClaimSwap
	DriverStageDone
)

func initDriver(nostr *nostr.Client, eth *goethclient.Client, stellar *stellargoclient.Client) *Driver {
	return &Driver{
		nostr:   nostr,
		eth:     eth,
		stellar: stellar,

		swapId: uuid.NewString(),
	}
}

// Buy flow for the driver
func (d *Driver) Buy(ctx context.Context, seller string, sale nostr.Product, amount uint) error {
	d.saleId = sale.Id
	d.stage = DriverStageStartBuy
	msgChan, err := d.nostr.SubscribeDirectMessagesDirect(sale.Id)
	if err != nil {
		return errors.Wrap(err, "could not subscribe to direct messages")
	}
	d.msges = msgChan

	go handleMessage(d)

	msg := MsgBuy{
		Id:     sale.Id,
		Amount: amount,
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "could not encode buy message")
	}
	return d.nostr.PublishDirectMessage(ctx, seller, []string{"s", sale.Id}, string(data))
}

// OpenSale on the driver
func (d *Driver) OpenSale(sale nostr.Product) error {
	d.saleId = sale.Id
	d.stage = DriverStageOpenSale
	msgChan, err := d.nostr.SubscribeDirectMessagesDirect(sale.Id)
	if err != nil {
		return errors.Wrap(err, "could not subscribe to direct messages")
	}
	d.msges = msgChan

	// TODO
	return errors.New("TODO")
}

func (d *Driver) handleBuyRequest(ctx context.Context, buyer string, req MsgBuy) {
	if req.Id != d.saleId {
		log.Debug().Msg("Ignore message which is not intended for this swap")
		return
	}
	// set swap amount
	if req.Amount > d.swapAmount {
		log.Debug().Msg("Buyer wants more TFT than we have, ignore")
		return
	}
	d.swapAmount = req.Amount

	msg := MsgAccept{
		EthAddress:     d.eth.AddressFromKey(),
		StellarAddress: d.stellar.Address(),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Error().Err(err).Msg("Can not encode accept msg")
	}

	if err := d.nostr.PublishDirectMessage(ctx, buyer, []string{"s", d.saleId}, string(data)); err != nil {
		log.Error().Err(err).Msg("Can not send buy accpeted message")
		return
	}

	d.stage = DriverStageAcceptedBuy
}

func handleMessage(driver *Driver) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for evt := range driver.msges {
		var reply string
		switch driver.stage {
		case DriverStageOpenSale:
			msg := MsgBuy{}
			if err := json.Unmarshal([]byte(evt.Content), &msg); err != nil {
				log.Debug().Err(err).Msg("could not decode message in atomic swap driver")
				continue
			}
			driver.handleBuyRequest(ctx, evt.PubKey, msg)
		}
	}
}
