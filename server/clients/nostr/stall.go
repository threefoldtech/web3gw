package nostr

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

// Code to support NIP-15

type (
	// Stall of products for sale
	Stall struct {
		Id          string     `json:"id"`
		Name        string     `json:"name"`
		Description string     `json:"description,omitempty"`
		Currency    string     `json:"currency"`
		Shipping    []Shipping `json:"shipping"`
	}

	// Shipping information for a Stall
	Shipping struct {
		Id        string   `json:"id"`
		Name      string   `json:"name,omitempty"`
		Cost      float64  `json:"cost"`
		Countries []string `json:"countries"`
	}

	// Product for sale in a stall
	Product struct {
		Id          string   `json:"id"`
		StallId     string   `json:"stall_id"`
		Name        string   `json:"name"`
		Description string   `json:"description,omitempty"`
		Images      []string `json:"images,omitempty"`
		Currency    string   `json:"currency"`
		Price       float64  `json:"price"`
		Quantity    uint     `json:"quantity"`
		// Specs is an array of key value pairs
		Specs [][]string `json:"specs"`
	}
)

const (
	// kindSetStall creates or updates a stall
	kindSetStall = 30017
	// kindSetProduct creates or updates a product
	kindSetProduct = 30018
)

var (
	// TagTftAtomicSwapSale is the searchable tag for atomic swaps in tft
	TagTftAtomicSwapSale = []string{"t", "tft_atomic_swap_sale_order"}
)

// PublishStall to connected relays. If a stall with the given ID was already published, conforming relays should update it
func (c *Client) PublishStall(ctx context.Context, tags []string, content Stall) error {
	// Force user to set the ID, even for new stalls
	if content.Id == "" {
		return errors.New("stall must have an ID")
	}

	if content.Name == "" {
		return errors.New("stall name must not be empty")
	}

	if content.Currency == "" {
		return errors.New("stall currency must not be empty")
	}

	// TODO: shipping validation?

	marshalledContent, err := json.Marshal(content)
	if err != nil {
		return errors.Wrap(err, "could not encode metadata")
	}
	return c.publishEventToRelays(ctx, kindSetStall, [][]string{{"d", content.Id}, tags}, string(marshalledContent))
}

// PublishProduct to connected relays. If a product with the given ID was already published, conforming relays should update it
func (c *Client) PublishProduct(ctx context.Context, tags []string, content Product) error {
	// Force user to set the ID, even for new products
	if content.Id == "" {
		return errors.New("product must have an ID")
	}

	if content.StallId == "" {
		return errors.New("product must be part of a stall")
	}

	if content.Name == "" {
		return errors.New("product name must not be empty")
	}

	if content.Currency == "" {
		return errors.New("product currency must not be empty")
	}

	if content.Price < 0 {
		return errors.New("price for a product can't be less than 0")
	}

	for _, spec := range content.Specs {
		if len(spec) != 2 {
			return errors.New("specs must be a list of key value pairs")
		}
	}

	marshalledContent, err := json.Marshal(content)
	if err != nil {
		return errors.Wrap(err, "could not encode metadata")
	}
	return c.publishEventToRelays(ctx, kindSetProduct, [][]string{{"d", content.Id}, tags}, string(marshalledContent))
}
