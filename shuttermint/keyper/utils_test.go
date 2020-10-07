package keyper

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestBatchHash(t *testing.T) {
	var zeroHash common.Hash
	tx1 := []byte{0xaa, 0xaa}
	tx2 := []byte{0xbb, 0xbb}

	// computed with Python contract test contract utils
	bh1 := common.HexToHash("15455e7675bd6d7a73a2cb8eb54a354f8180881d4c6b410556b505814b94cda1")
	bh2 := common.HexToHash("5081bbf18eac5fd3c87fb408bd54dca6008ec51aca62e0e4ca4631b96aa43f58")
	bh12 := common.HexToHash("bebf2f754c45520216ca08d65c4fef1b2dc98762f85dbabc3ce6abe23086ddbc")
	bh21 := common.HexToHash("2cb6f9ae6401f16792872bdeabc5f29cb0e52a949e35399e322a5af8bddd1489")

	require.Equal(t, ComputeBatchHash([][]byte{}), zeroHash)
	require.Equal(t, ComputeBatchHash([][]byte{tx1}), bh1)
	require.Equal(t, ComputeBatchHash([][]byte{tx2}), bh2)
	require.Equal(t, ComputeBatchHash([][]byte{tx1, tx2}), bh12)
	require.Equal(t, ComputeBatchHash([][]byte{tx2, tx1}), bh21)
}

func TestDecryptionKeyEncoding(t *testing.T) {
	keys := []*ecdsa.PrivateKey{}
	for i := 0; i < 5; i++ {
		key, err := crypto.GenerateKey()
		require.Nil(t, err)
		keys = append(keys, key)
	}

	bs := make([]byte, 32)
	copy(bs[31:], []byte{1})
	oneKey, err := crypto.ToECDSA(bs)
	require.Nil(t, err)
	keys = append(keys, oneKey)

	for _, key := range keys {
		encoded := DecryptionKeyToBytes(key)
		require.True(t, len(encoded) <= 32)

		recoveredKey, err := crypto.ToECDSA(encoded)
		require.Nil(t, err)
		require.Equal(t, recoveredKey, key)
	}
}

func TestComputeCipherDecryptionSigning(t *testing.T) {
	batcherContractAddress := common.HexToAddress("1cd2a8349c8508756f75c5addb1aa01b8df60799")
	cipherBatchHash := common.HexToHash("c67a3d98c6950912688d10218a54576af8734b6f648fe0c7f288af47439730ac")
	decryptionKeyBytes, _ := hex.DecodeString("5bc5d4a9a24ca47a19f7c6d4ed04bd864dc7eabc721286e06b266bd4faf66546")
	decryptionKey, _ := crypto.ToECDSA(decryptionKeyBytes)
	batchHash := common.HexToHash("13ab69a28f2a53b636c3886a748c41cdf911ee66e025a10317d0f11320388f4c")

	signingKeyBytes, _ := hex.DecodeString("a154d1e21620e97eaac10355eef72a488d04c95ffb0d289507fa48ba2d6653b8")
	signingKey, _ := crypto.ToECDSA(signingKeyBytes)

	expectedSignature, _ := hex.DecodeString("bc0af048df027eb11893647d1833d0b4a2e27df6a4fe1c8bf5a0fc16f77912d37282ac5c0a5e24288c9365007982199106cf308150c85253b76623457ff9d7f41b")

	signature, err := ComputeCipherDecryptionSignature(signingKey, batcherContractAddress, cipherBatchHash, decryptionKey, batchHash)
	fmt.Printf("%x %d\n", signature, len(signature))
	require.Nil(t, err)
	require.Equal(t, signature, expectedSignature)
}
