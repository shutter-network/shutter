package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kr/pretty"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/types"

	"github.com/brainbot-com/shutter/shuttermint/keyper"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

var version string = "(unknown)"

func makeMessage() *shmsg.Message {
	return &shmsg.Message{
		Payload: &shmsg.Message_PublicKeyCommitment{
			PublicKeyCommitment: &shmsg.PublicKeyCommitment{
				BatchIndex: 1,
				Commitment: []byte("foobar"),
			},
		},
	}
}

func printEvents(events []abcitypes.Event) {
	for _, ev := range events {
		x, err := keyper.MakeEvent(ev)
		if err != nil {
			fmt.Println(err)
		} else {
			pretty.Println(ev.Type, "=>", x)
		}
	}
}

func txsearch(cl client.Client) {
	query := "shutter.batch-config.StartBatchIndex>=0"
	res, err := cl.TxSearch(query, false, 0, 50, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("transaction count", res.TotalCount)
	for _, tx := range res.Txs {
		fmt.Printf("=== tx height=%d\n", tx.Height)
		printEvents(tx.TxResult.GetEvents())
	}
}

func status(cl client.StatusClient) {
	st, err := cl.Status()
	if err != nil {
		panic(err)
	}
	pretty.Println("Status:", st)
}

func subscribe(cl client.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	query := "tx.height > 3"

	txs, err := cl.Subscribe(ctx, "test-client", query)
	if err != nil {
		panic(err)
	}

	go func() {
		for e := range txs {
			d := e.Data.(types.EventDataTx)
			events := d.TxResult.Result.Events
			printEvents(events)
		}
	}()
	time.Sleep(time.Hour)

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		panic(err)
	}
	msg := makeMessage()
	signedMessage, err := shmsg.SignMessage(msg, privateKey)
	if err != nil {
		panic(err)
	}

	var tx types.Tx = types.Tx(base64.RawURLEncoding.EncodeToString(signedMessage))
	res, err := cl.BroadcastTxCommit(tx)

	fmt.Println("Msg:", base64.RawURLEncoding.EncodeToString(signedMessage))

	if err != nil {
		panic(err)
	}
	pretty.Println("Res", res)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Printf("Starting testclient version %s", version)

	var cl client.Client
	cl, err := http.New("http://localhost:26657", "/websocket")
	if err != nil {
		panic(err)
	}
	err = cl.Start()
	if err != nil {
		panic(err)
	}

	defer func() {
		err = cl.Stop()
		if err != nil {
			panic(err)
		}
	}()

	status(cl)
	txsearch(cl)
	subscribe(cl)
}
