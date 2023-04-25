package ipfs

import (
	"bytes"
	"context"
	"io"

	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-cid"
	"github.com/rs/zerolog/log"
)

type Client struct {
	peer *ipfslite.Peer
}

func NewClient(peer *ipfslite.Peer) *Client {
	return &Client{peer: peer}
}

func (c *Client) StoreFile(ctx context.Context, data []byte) (string, error) {
	node, err := c.peer.AddFile(ctx, bytes.NewReader(data), &ipfslite.AddParams{})
	if err != nil {
		return "", err
	}

	log.Debug().Msgf("IPFS: stored file with contentId: %s", node.Cid().String())

	return node.Cid().String(), nil
}

func (c *Client) GetFile(ctx context.Context, contentId string) ([]byte, error) {
	log.Debug().Msgf("IPFS: trying to get file with contentId: %s", contentId)

	cId, err := cid.Decode(contentId)
	if err != nil {
		return nil, err
	}

	node, err := c.peer.GetFile(ctx, cId)
	if err != nil {
		return nil, err
	}

	defer node.Close()
	content, err := io.ReadAll(node)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (c *Client) RemoveFile(ctx context.Context, contentId string) (bool, error) {
	log.Debug().Msgf("IPFS: trying to remove file with contentId: %s", contentId)

	cId, err := cid.Decode(contentId)
	if err != nil {
		return false, err
	}

	err = c.peer.Remove(ctx, cId)
	if err != nil {
		return false, err
	}

	return true, nil
}
