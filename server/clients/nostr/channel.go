package nostr

import (
	"context"
	"encoding/json"

	"github.com/nbd-wtf/go-nostr"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type (
	Channel struct {
		Name    string `json:"name"`
		About   string `json:"about"`
		Picture string `json:"picture"`
	}

	RelayChannel struct {
		Channel
		Relay string `json:"relay"`
		Id    string `json:"id"`
	}

	ChannelMessage struct {
		// Content of the message
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		// MessageID is used for replies
		MessageID string `json:"message_id"`
		// PublicKey of author to reply to
		PublicKey string `json:"public_key"`
	}

	RelayChannelMessage struct {
		Content string     `json:"content"`
		Tags    [][]string `json:"tags"`
		Relay   string     `json:"relay"`
		Id      string     `json:"id"`
	}
)

const (
	// kindCreateChannel creates a channel
	kindCreateChannel = 40
	// kindSetChannelMetadata updates channel metadata
	kindSetChannelMetadata = 41
	// kindCreateChannelMessage creates a message in a channel
	kindCreateChannelMessage = 42
	// kindHideChannelMessage hides a message in the channel
	kindHideChannelMessage = 43
	// kindMuteChannelUser mutes a channel user for the sending user
	kindMuteChanneluser = 44
)

// CreateChannel creates a new channel
func (c *Client) CreateChannel(ctx context.Context, tags []string, content Channel) (string, error) {
	if content.Name == "" {
		return "", errors.New("Channel must have a name")
	}
	marshalledContent, err := json.Marshal(content)
	if err != nil {
		return "", errors.Wrap(err, "could not encode metadata")
	}

	return c.publishEventToRelays(ctx, kindCreateChannel, [][]string{tags}, string(marshalledContent))
}

// UpdateChannelMetadata updates the channel metdata. ChannelID is the event ID of the create channel event used to create the channel to update
func (c *Client) UpdateChannelMetadata(ctx context.Context, tags []string, channelID string, content Channel) error {
	if content.Name == "" {
		return errors.New("Channel must have a name")
	}
	marshalledContent, err := json.Marshal(content)
	if err != nil {
		return errors.Wrap(err, "could not encode metadata")
	}

	if _, err := c.publishEventToRelays(ctx, kindSetChannelMetadata, [][]string{tags, {"e", channelID}}, string(marshalledContent)); err != nil {
		return err
	}

	return nil
}

// CreateChannelRootMessage creates a message in channel. If replyTo is the empty string, it is marked as a root
func (c *Client) CreateChannelRootMessage(ctx context.Context, message ChannelMessage) (string, error) {
	if message.Content == "" {
		return "", errors.New("Refusing to submit empty message")
	}

	tags := [][]string{}
	if message.ChannelID != "" {
		tags = append(tags, []string{"e", message.ChannelID, "", "root"})
	}

	if message.MessageID != "" {
		tags = append(tags, []string{"e", message.MessageID, "", "reply"})
	}

	if message.PublicKey != "" {
		tags = append(tags, []string{"p", message.PublicKey})
	}

	return c.publishEventToRelays(ctx, kindCreateChannelMessage, tags, message.Content)
}

// HideMessage marks a message as hidden for the user. It should be noted that properly handling this is mostly up to the clients
func (c *Client) HideMessage(ctx context.Context, tags []string, messageID string, content string) error {
	if _, err := c.publishEventToRelays(ctx, kindHideChannelMessage, [][]string{tags, {"e", messageID}}, content); err != nil {
		return err
	}

	return nil
}

// MuteUser marks a user as muted for the current user. It should be noted that properly handling this is mostly up to the clients.
// The user to mute is identified by it's pubkey
func (c *Client) MuteUser(ctx context.Context, tags []string, user string, content string) error {
	if _, err := c.publishEventToRelays(ctx, kindMuteChanneluser, [][]string{tags, {"p", user}}, content); err != nil {
		return err
	}

	return nil
}

func (c *Client) SubscribeChannelCreation() (string, error) {
	filters := []nostr.Filter{{
		Kinds: []int{nostr.KindChannelCreation},
		Limit: DEFAULT_LIMIT,
	}}

	return c.subscribeWithFiler(filters)
}

// SubscribeChannelMessages subsribes to messages which are a reply to the given chanMessageId
func (c *Client) SubscribeChannelMessages(chanMessageId string) (string, error) {
	filters := []nostr.Filter{{
		Kinds: []int{nostr.KindChannelMessage},
		Limit: DEFAULT_LIMIT,
		Tags:  nostr.TagMap{"e": []string{chanMessageId}},
	}}

	return c.subscribeWithFiler(filters)
}

func (c *Client) FetchChannelCreation() ([]RelayChannel, error) {
	filters := []nostr.Filter{{
		Kinds: []int{nostr.KindChannelCreation},
		Limit: DEFAULT_LIMIT,
	}}

	channelCreationEvents, err := c.fetchEventsWithFilter(filters)
	if err != nil {
		return nil, err
	}

	rc := make([]RelayChannel, 0, len(channelCreationEvents))

	for _, cce := range channelCreationEvents {
		var c Channel
		if err := json.Unmarshal([]byte(cce.Event.Content), &c); err != nil {
			log.Warn().Err(err).Msg("could not decode channel create message")
			continue
		}
		rc = append(rc, RelayChannel{
			Channel: c,
			Id:      cce.Event.ID,
			Relay:   cce.Relay,
		})
	}

	return rc, nil
}

// SubscribeChannelMessages subsribes to messages which are a reply to the given chanMessageId
func (c *Client) FetchChannelMessages(channelID string) ([]RelayChannelMessage, error) {
	filters := []nostr.Filter{{
		Kinds: []int{nostr.KindChannelMessage},
		Limit: DEFAULT_LIMIT,
		Tags:  nostr.TagMap{"e": []string{channelID}},
	}}

	channelMessageEvents, err := c.fetchEventsWithFilter(filters)
	if err != nil {
		return nil, err
	}

	rm := make([]RelayChannelMessage, 0, len(channelMessageEvents))

	for _, cme := range channelMessageEvents {
		log.Debug().Msgf("incoming channel message event: %+v", cme)
		// var m ChannelMessage
		// if err := json.Unmarshal([]byte(cme.Event.Content), &c); err != nil {
		// 	log.Warn().Err(err).Msg("could not decode channel message")
		// 	continue
		// }
		rm = append(rm, RelayChannelMessage{
			Content: cme.Event.Content,
			Tags:    getTags(cme.Event.Tags),
			Id:      cme.Event.ID,
			Relay:   cme.Relay,
		})
	}

	return rm, nil
}

func getTags(tags nostr.Tags) [][]string {
	ret := [][]string{}
	for _, tag := range tags {
		ret = append(ret, tag)
	}

	return ret
}
