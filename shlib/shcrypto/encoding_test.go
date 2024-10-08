package shcrypto

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	blst "github.com/supranational/blst/bindings/go"
	"gotest.tools/v3/assert"
)

func encryptedMessage() *EncryptedMessage {
	blocks := []Block{}
	for i := 0; i < 3; i++ {
		s := bytes.Repeat([]byte{byte(i)}, 32)
		var b Block
		copy(b[:], s)
		blocks = append(blocks, b)
	}

	return &EncryptedMessage{
		C1: makeTestG2(5),
		C2: blocks[0],
		C3: blocks[1:],
	}
}

func TestMarshalUnmarshal(t *testing.T) {
	m1 := encryptedMessage()
	m2 := &EncryptedMessage{}
	err := m2.Unmarshal(m1.Marshal())
	assert.NilError(t, err)
	assert.DeepEqual(t, m1, m2, g2Comparer)
}

func TestUnmarshalBroken(t *testing.T) {
	d := encryptedMessage().Marshal()
	m := EncryptedMessage{}

	err := m.Unmarshal(d[:16])
	assert.Assert(t, err != nil)

	err = m.Unmarshal(d[:32])
	assert.Assert(t, err != nil)

	err = m.Unmarshal(d[:65])
	assert.Assert(t, err != nil)

	err = m.Unmarshal(d[:len(d)-1])
	assert.Assert(t, err != nil)

	v := d[:]
	v[0]++
	err = m.Unmarshal(v)
	assert.Assert(t, err != nil)
}

func TestMarshal(t *testing.T) {
	ask := (*EonSecretKeyShare)(big.NewInt(123))
	ashM := ask.Marshal()
	askD := new(EonSecretKeyShare)
	assert.NilError(t, askD.Unmarshal(ashM))
	assert.Check(t, ask.Equal(askD))

	apks := (*EonPublicKeyShare)(makeTestG2(5))
	apksM := apks.Marshal()
	apksD := new(EonPublicKeyShare)
	assert.NilError(t, apksD.Unmarshal(apksM))
	assert.Check(t, apksD.Equal(apks))

	apk := (*EonPublicKey)(makeTestG2(6))
	apkM := apk.Marshal()
	apkD := new(EonPublicKey)
	assert.NilError(t, apkD.Unmarshal(apkM))
	assert.Check(t, apkD.Equal(apk))

	esks := (*EpochSecretKeyShare)(makeTestG1(7))
	esksM := esks.Marshal()
	esksD := new(EpochSecretKeyShare)
	assert.NilError(t, esksD.Unmarshal(esksM))
	assert.Check(t, esksD.Equal(esks))

	esk := (*EpochSecretKey)(makeTestG1(8))
	eskM := esk.Marshal()
	eskD := new(EpochSecretKey)
	assert.NilError(t, eskD.Unmarshal(eskM))
	assert.Check(t, eskD.Equal(esk))
}

func TestIdentifyVersion(t *testing.T) {
	d := encryptedMessage().Marshal()
	assert.Assert(t, IdentifyVersion(d) == VersionIdentifier)

	// legacy version
	assert.Assert(t, IdentifyVersion(d[1:]) != VersionIdentifier)
	assert.Assert(t, IdentifyVersion(d[1:]) == 0x00)
}

func TestMarshalGammasEmpty(t *testing.T) {
	g := &Gammas{}
	m := g.Marshal()
	assert.Assert(t, len(m) == 0)
}

func TestMarshalGammasError(t *testing.T) {
	validGammaEncoding := makeTestG2(1).Compress()
	inputs := [][]byte{
		{0x00},
		bytes.Repeat([]byte{0xaa}, 96),
		common.FromHex(
			"87f481803120be4e565dc88cdbb1ae4c1ddfa249bd34cdc43982d926278535e84ee584aa9ae3553f56d02d3aa842b3941058f0dcabacbc551dca1d04ba7647c806acf7f960809438993359338dc4858aadcbce50f9a370986c74053303ab4449",
		),
		validGammaEncoding[:len(validGammaEncoding)-1],
		append([]byte{0x00}, validGammaEncoding...),
	}
	for _, input := range inputs {
		g := &Gammas{}
		err := g.Unmarshal(input)
		assert.Assert(t, err != nil)
	}
}

func TestMarshalGammasRoundtrip(t *testing.T) {
	gammaValues := []*Gammas{
		{},
		{makeTestG2(1)},
		{makeTestG2(2), makeTestG2(2), makeTestG2(3), makeTestG2(4), makeTestG2(5), makeTestG2(6)},
	}
	for _, gammas := range gammaValues {
		m := gammas.Marshal()
		assert.Assert(t, len(m) == len([]*blst.P2Affine(*gammas))*96)
		g := &Gammas{}
		assert.NilError(t, g.Unmarshal(m))
		assert.DeepEqual(t, gammas, g, g2Comparer)

		mText, err := gammas.MarshalText()
		assert.NilError(t, err)
		gText := &Gammas{}
		assert.NilError(t, gText.UnmarshalText(mText))
		assert.DeepEqual(t, gammas, gText, g2Comparer)
	}
}
