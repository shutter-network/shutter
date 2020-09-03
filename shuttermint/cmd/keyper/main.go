package main

import (
	"log"

	"github.com/brainbot-com/shutter/shuttermint/keyper"
	"github.com/ethereum/go-ethereum/crypto"
)

var version string = "(unknown)"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Printf("Starting keyper version %s", version)

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		panic(err)
	}
	k := keyper.NewKeyper(privateKey, "http://localhost:26657")
	err = k.Run()
	if err != nil {
		panic(err)
	}
}
