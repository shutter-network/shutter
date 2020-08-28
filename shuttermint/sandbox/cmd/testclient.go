package main

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
	"github.com/brainbot-com/shutter/shuttermint/smclient"
	"github.com/ethereum/go-ethereum/crypto"
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
	cl, err := smclient.Dial("http://localhost:26657")
	if err != nil {
		panic(err)
	}
	fmt.Println("got a client:", cl)
	st, err := cl.GetStatus(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(st))

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		panic(err)
	}
	msg := makeMessage()
	signedMessage, err := shmsg.SignMessage(msg, privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Msg:", base64.RawURLEncoding.EncodeToString(signedMessage))
	res, err := cl.BroadcastTXCommit(context.Background(), signedMessage)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Res: %+v\n", res)
}
