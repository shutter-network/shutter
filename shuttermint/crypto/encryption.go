package crypto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

// EncryptedMessage represents the full output of the encryption procedure.
type EncryptedMessage struct {
	C1 *bn256.G2
	C2 Block
	C3 []Block
}

// Block represents a block of data.
type Block [BlockSize]byte

// BlockSize is the size in bytes of the blocks into which a message is split up before encryption.
const BlockSize = 32

// HashBytesToBlock hashes the given byte slice and returns the result as a block.
func HashBytesToBlock(d []byte) Block {
	h := crypto.Keccak256(d)
	var b Block
	copy(b[:], h)
	return b
}

// HashBlockToInt hashes a block and returns the result as an integer in Z_q.
func HashBlockToInt(d Block) *big.Int {
	h := crypto.Keccak256(d[:])
	i := new(big.Int).SetBytes(h)
	i.Mod(i, bn256.Order) // TODO: check if this is fine
	return i
}

// HashGTToBlock hashes an element of GT and returns the result as a block.
func HashGTToBlock(gt *bn256.GT) Block {
	b := gt.Marshal()
	return HashBytesToBlock(b)
}

// XORBlocks xors the two blocks and returns the result.
func XORBlocks(b1 Block, b2 Block) Block {
	var b Block
	for i := 0; i < BlockSize; i++ {
		b[i] = b1[i] ^ b2[i]
	}
	return b
}

// RandomSigma returns a random value to be used during encryption.
func RandomSigma(r io.Reader) (Block, error) {
	data := make([]byte, BlockSize)
	_, err := r.Read(data)
	if err != nil {
		return Block{}, nil
	}
	var b Block
	copy(b[:], data)
	return b, nil
}

// Encrypt encrypts a message for the epoch given by its id. It uses the eon public key and randomness
// provided in sigma.
func Encrypt(message []byte, eonPublicKey *EonPublicKey, epochID *EpochID, sigma Block) *EncryptedMessage {
	messageBlocks := PadMessage(message)
	r := computeR(sigma)
	result := EncryptedMessage{
		C1: computeC1(r),
		C2: computeC2(sigma, r, epochID, eonPublicKey),
		C3: computeC3(messageBlocks, sigma),
	}
	return &result
}

func computeR(sigma Block) *big.Int {
	return HashBlockToInt(sigma)
}

func computeC1(r *big.Int) *bn256.G2 {
	return new(bn256.G2).ScalarBaseMult(r)
}

func computeC2(sigma Block, r *big.Int, epochID *EpochID, eonPublicKey *EonPublicKey) Block {
	pairing := bn256.Pair((*bn256.G1)(epochID), (*bn256.G2)(eonPublicKey))
	preimage := new(bn256.GT).ScalarMult(pairing, r)
	key := HashGTToBlock(preimage)
	return XORBlocks(sigma, key)
}

func computeC3(blocks []Block, sigma Block) []Block {
	encryptedBlocks := []Block{}
	numBlocks := len(blocks)
	keys := computeBlockKeys(sigma, numBlocks)
	for i := 0; i < numBlocks; i++ {
		encryptedBlock := XORBlocks(keys[i], blocks[i])
		encryptedBlocks = append(encryptedBlocks, encryptedBlock)
	}
	return encryptedBlocks
}

func computeBlockKeys(sigma Block, n int) []Block {
	keys := []Block{}
	buf := make([]byte, binary.MaxVarintLen64)
	for i := int64(0); i < int64(n); i++ {
		binary.PutVarint(buf, i)
		preimage := append(sigma[:], buf...)
		key := HashBytesToBlock(preimage)
		keys = append(keys, key)
	}
	return keys
}

// Decrypt decrypts the given message using the given epoch secret key.
func (m *EncryptedMessage) Decrypt(epochSecretKey *EpochSecretKey) ([]byte, error) {
	sigma := m.Sigma(epochSecretKey)
	decryptedBlocks := decryptBlocks(m.C3, sigma)
	return UnpadMessage(decryptedBlocks)
}

// Sigma computes the sigma value of the encrypted message given the epoch secret key.
func (m *EncryptedMessage) Sigma(epochSecretKey *EpochSecretKey) Block {
	pairing := bn256.Pair((*bn256.G1)(epochSecretKey), m.C1)
	key := HashGTToBlock(pairing)
	sigma := XORBlocks(m.C2, key)
	return sigma
}

func decryptBlocks(encryptedBlocks []Block, sigma Block) []Block {
	numBlocks := len(encryptedBlocks)
	keys := computeBlockKeys(sigma, numBlocks)
	decryptedBlocks := []Block{}
	for i := 0; i < numBlocks; i++ {
		decryptedBlock := XORBlocks(encryptedBlocks[i], keys[i])
		decryptedBlocks = append(decryptedBlocks, decryptedBlock)
	}
	return decryptedBlocks
}

// PadMessage pads a message and returns it as a sequence of blocks.
// Implements PKCS #7 according to https://www.ietf.org/rfc/rfc2315.txt
func PadMessage(m []byte) []Block {
	paddingLength := BlockSize - len(m)%BlockSize
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	padded := append(m, padding...)

	blocks := []Block{}
	numBlocks := len(padded) / BlockSize
	for i := 0; i < numBlocks; i++ {
		var block Block
		copy(block[:], padded[i*BlockSize:(i+1)*BlockSize])
		blocks = append(blocks, block)
	}
	return blocks
}

// UnpadMessage returns the message provided in padded form as a sequence of blocks.
func UnpadMessage(blocks []Block) ([]byte, error) {
	if len(blocks) == 0 {
		return []byte{}, nil
	}

	lastBlock := blocks[len(blocks)-1]
	paddingLength := int(lastBlock[BlockSize-1])
	if paddingLength == 0 {
		return nil, fmt.Errorf("invalid padding length 0")
	}
	if paddingLength > BlockSize {
		return nil, fmt.Errorf("invalid padding length %d (greater than block size %d)", paddingLength, BlockSize)
	}

	m := make([]byte, len(blocks)*BlockSize-paddingLength)
	// copy unpadded blocks
	for i := 0; i < len(blocks)-1; i++ {
		start := i * BlockSize
		end := (i + 1) * BlockSize
		copy(m[start:end], blocks[i][:])
	}
	// copy padded block
	start := (len(blocks) - 1) * BlockSize
	end := len(blocks)*BlockSize - paddingLength
	copy(m[start:end], lastBlock[:BlockSize-paddingLength])

	return m, nil
}
