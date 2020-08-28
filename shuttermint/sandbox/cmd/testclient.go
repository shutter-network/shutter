package main

import (
	"encoding/base64"
	"fmt"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/types"
)

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

func main() {
	cl, err := http.New("http://localhost:26657", "/websocket")
	if err != nil {
		panic(err)
	}
	// if !cl.IsRunning() {
	//	panic("tendermint not running")
	// }

	fmt.Println("got a client:", cl)
	st, err := cl.Status()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %+v\n", st)

	// st, err := cl.GetStatus(context.Background())
	// if err != nil {
	//	panic(err)
	// }
	// fmt.Println(string(st))

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		panic(err)
	}
	msg := makeMessage()
	signedMessage, err := shmsg.SignMessage(msg, privateKey)
	if err != nil {
		panic(err)
	}

	var tx types.Tx
	tx = types.Tx(base64.RawURLEncoding.EncodeToString(signedMessage))
	res, err := cl.BroadcastTxCommit(tx)

	fmt.Println("Msg:", base64.RawURLEncoding.EncodeToString(signedMessage))

	// res, err := cl.BroadcastTXCommit(context.Background(), signedMessage)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Res: %+v\n", res)
}
