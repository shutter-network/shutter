package main

import (
	"bytes"
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

	"github.com/shutter-network/shutter/shuttermint/cmd/shversion"
	"github.com/shutter-network/shutter/shuttermint/keyper/shutterevents"
	"github.com/shutter-network/shutter/shuttermint/shmsg"
)

func makeMessage() *shmsg.MessageWithNonce {
	msg := &shmsg.Message{
		Payload: &shmsg.Message_CheckIn{
			CheckIn: &shmsg.CheckIn{
				ValidatorPublicKey:  bytes.Repeat([]byte("x"), 32),
				EncryptionPublicKey: bytes.Repeat([]byte("y"), 33),
			},
		},
	}
	return &shmsg.MessageWithNonce{
		RandomNonce: uint64(0),
		Msg:         msg,
	}
}

func printEvents(events []abcitypes.Event, height int64) {
	for _, ev := range events {
		x, err := shutterevents.MakeEvent(ev, height)
		if err != nil {
			fmt.Println(err)
		} else {
			pretty.Println(ev.Type, "=>", x)
		}
	}
}

func txsearch(cl client.Client) {
	query := "shutter.batch-config.StartBatchIndex>=0"
	page := 1
	perPage := 50
	res, err := cl.TxSearch(context.Background(), query, false, &page, &perPage, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("transaction count", res.TotalCount)
	for _, tx := range res.Txs {
		fmt.Printf("=== tx height=%d\n", tx.Height)
		printEvents(tx.TxResult.GetEvents(), tx.Height)
	}
}

func status(cl client.StatusClient) {
	st, err := cl.Status(context.Background())
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
			printEvents(events, d.Height)
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
	res, err := cl.BroadcastTxCommit(context.Background(), tx)

	fmt.Println("Msg:", base64.RawURLEncoding.EncodeToString(signedMessage))

	if err != nil {
		panic(err)
	}
	pretty.Println("Res", res)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Printf("Starting testclient version %s", shversion.Version())

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
