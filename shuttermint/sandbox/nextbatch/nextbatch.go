package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/brainbot-com/shutter/shuttermint/contracts/configcontract"
)

func main() {
	cl, err := ethclient.Dial("ws://localhost:8545/ws")
	if err != nil {
		panic(err)
	}
	addr := common.HexToAddress("0x07a457d878BF363E0Bb5aa0B096092f941e19962")
	cc, err := configcontract.NewConfigContract(addr, cl)
	if err != nil {
		panic(err)
	}
	header, err := cl.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("HEAD: %+v\n", header)
	headBlockNumber := header.Number.Uint64()
	fmt.Printf("head block #%d\n", headBlockNumber)
	next, err := cc.NextBatchIndex(headBlockNumber)
	if err != nil {
		panic(err)
	}
	fmt.Printf("next batch: %+v\n", next)
	bp, err := cc.QueryBatchParams(nil, next)
	if err != nil {
		panic(err)
	}
	fmt.Printf("batch params: %+v\n", bp)
}
