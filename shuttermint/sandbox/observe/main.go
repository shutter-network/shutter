package main

import (
	"context"

	"github.com/kr/pretty"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"

	"github.com/brainbot-com/shutter/shuttermint/keyper/observe"
)

func main() {
	var cl client.Client
	cl, err := http.New("http://localhost:26657", "/websocket")
	if err != nil {
		panic(err)
	}

	s := observe.NewShutter()

	err = s.SyncToHead(context.Background(), cl)
	if err != nil {
		panic(err)
	}
	pretty.Println("Synced:", s)
}
