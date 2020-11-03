package shmsg

import (
	"crypto/rand"
	"testing"

	bn256cf "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func randomG1() *bn256cf.G1 {
	_, g1, err := bn256cf.RandomG1(rand.Reader)
	if err != nil {
		panic(err)
	}
	return g1
}

func randomG2() *bn256cf.G2 {
	_, g2, err := bn256cf.RandomG2(rand.Reader)
	if err != nil {
		panic(err)
	}
	return g2
}

func randomGT() *bn256cf.GT {
	g1 := randomG1()
	g2 := randomG2()
	return bn256cf.Pair(g1, g2)
}

func TestG1Marshal(t *testing.T) {
	g := randomG1()
	msg := G1{}
	msg.Set(g)
	marshaled, err := proto.Marshal(&msg)
	require.Nil(t, err)

	umsg := G1{}
	err = proto.Unmarshal(marshaled, &umsg)
	require.Nil(t, err)
	ug, err := umsg.Get()
	require.Nil(t, err)
	require.Equal(t, g, ug)
}

func TestG2Marshal(t *testing.T) {
	g := randomG2()
	msg := G2{}
	msg.Set(g)
	marshaled, err := proto.Marshal(&msg)
	require.Nil(t, err)

	umsg := G2{}
	err = proto.Unmarshal(marshaled, &umsg)
	require.Nil(t, err)
	ug, err := umsg.Get()
	require.Nil(t, err)
	require.Equal(t, g, ug)
}

func TestGTMarshal(t *testing.T) {
	g := randomGT()
	msg := GT{}
	msg.Set(g)
	marshaled, err := proto.Marshal(&msg)
	require.Nil(t, err)

	umsg := GT{}
	err = proto.Unmarshal(marshaled, &umsg)
	require.Nil(t, err)
	ug, err := umsg.Get()
	require.Nil(t, err)
	require.Equal(t, g, ug)
}
