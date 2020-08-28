// smclient provides a client for the shuttermint app via JSONRPC

package smclient

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	c *rpc.Client
}

func NewClient(c *rpc.Client) *Client {
	return &Client{c}
}

func (c *Client) Close() {
	c.c.Close()
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	return DialContext(context.Background(), rawurl)
}

func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

func (c *Client) GetStatus(ctx context.Context) (json.RawMessage, error) {
	var raw json.RawMessage
	err := c.c.CallContext(ctx, &raw, "status")
	if err != nil {
		return nil, err
	}
	return raw, nil
}

type BroadcastTXCommit struct {
	Hash   string `json:"hash,omitempty"`
	Height string `json:"height,omitempty"`
}

func (c *Client) BroadcastTXCommit(ctx context.Context, signedMessage []byte) (BroadcastTXCommit, error) {
	var raw json.RawMessage
	encoded := []byte(base64.RawURLEncoding.EncodeToString(signedMessage))
	err := c.c.CallContext(ctx, &raw, "broadcast_tx_commit", encoded)
	if err != nil {
		return BroadcastTXCommit{}, err
	}
	var res BroadcastTXCommit
	err = json.Unmarshal(raw, &res)
	return res, err
}
