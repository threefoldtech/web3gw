package ipfs

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/LeeSmet/go-jsonrpc"
	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-cid"
	"github.com/rs/zerolog/log"
)

const (
	// IpfsID is the ID for state of a ipfs client in the connection state.
	IpfsID = "ipfs"
)

type (
	// Client exposes ipfs related functionality
	Client struct {
		peer *ipfslite.Peer
	}
	// state managed by ipfs client
	ipfsState struct {
		cids map[string][]byte
	}
)

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *ipfsState {
	raw, exists := conState[IpfsID]
	if !exists {
		ns := &ipfsState{
			cids: make(map[string][]byte),
		}
		conState[IpfsID] = ns
		return ns
	}
	ns, ok := raw.(*ipfsState)
	if !ok {
		// This means the invariant is violated, so panic here is ok
		panic("Invalid saved state for ipfs")
	}
	return ns
}

func NewClient(peer *ipfslite.Peer) *Client {
	return &Client{peer: peer}
}

// ListCids lists all CIDs stored in the ipfs client
func (c *Client) ListCids(ctx context.Context, conState jsonrpc.State) ([]string, error) {
	log.Debug().Msg("IPFS: listing file cids")

	state := State(conState)
	var cids []string
	for cid := range state.cids {
		cids = append(cids, cid)
	}

	return cids, nil
}

// StoreFile stores a file in the ipfs client
func (c *Client) StoreFile(ctx context.Context, conState jsonrpc.State, data []byte) (string, error) {
	node, err := c.peer.AddFile(ctx, bytes.NewReader(data), &ipfslite.AddParams{})
	if err != nil {
		return "", err
	}

	log.Debug().Msgf("IPFS: stored file with contentId: %s", node.Cid().String())

	state := State(conState)
	state.cids[node.Cid().String()] = data

	return node.Cid().String(), nil
}

// GetFile gets a file from the ipfs client
func (c *Client) GetFile(ctx context.Context, conState jsonrpc.State, contentId string) ([]byte, error) {
	log.Debug().Msgf("IPFS: trying to get file with contentId: %s", contentId)

	cId, err := cid.Decode(contentId)
	if err != nil {
		return nil, err
	}

	// Check if we have the content in our state
	state := State(conState)
	_, found := state.cids[contentId]
	if !found {
		return nil, errors.New("contentId not found in state")
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

// RemoveFile removes a file from the ipfs client
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

// RemoveAllFiles removes all files from the ipfs client
func (c *Client) RemoveAllFiles(ctx context.Context, conState jsonrpc.State) error {
	state := State(conState)
	for id := range state.cids {
		cId, err := cid.Decode(id)
		if err != nil {
			return err
		}

		err = c.peer.Remove(ctx, cId)
		if err != nil {
			return err
		}
	}
	return nil
}
