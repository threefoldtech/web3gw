package nostr

import "context"

// Support for NIP-09 - event deletion

const (
	// kindDelete deletes a prodcut or stall
	kindDelete = 5
)

// PublishEventDeletion to connected relays. If an event with the given ID was already published, conforming relays should delete it
func (c *Client) PublishEventDeletion(ctx context.Context, tags []string, id string) error {
	return c.publishEventToRelays(ctx, kindDelete, [][]string{{"e", id}, tags}, "")
}
