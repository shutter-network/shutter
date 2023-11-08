package shcrypto

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"

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
	var b Block
	l, err := r.Read(b[:])
	if l != BlockSize {
		return Block{}, errors.New("didn't read all random bytes")
	}
	if err != nil {
		return Block{}, err
	}
	return b, nil
}

// Encrypt encrypts a message for the epoch given by its id. It uses the eon public key and randomness
// provided in sigma.
func Encrypt(message []byte, eonPublicKey *EonPublicKey, epochID *EpochID, sigma Block) *EncryptedMessage {
	messageBlocks := PadMessage(message)
	r := computeR(sigma, message)
	result := EncryptedMessage{
		C1: computeC1(r),
		C2: computeC2(sigma, r, epochID, eonPublicKey),
		C3: computeC3(messageBlocks, sigma),
	}
	return &result
}

func computeR(sigma Block, message []byte) *big.Int {
	messageHash := HashBytesToBlock(message)
	return HashBlocksToInt(sigma, messageHash)
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
	for i := 0; i < numBlocks; i++ {
		key := computeBlockKey(sigma, uint64(i))
		encryptedBlock := XORBlocks(key, blocks[i])
		encryptedBlocks = append(encryptedBlocks, encryptedBlock)
	}
	return encryptedBlocks
}

func Uint64toBytes(i uint64) ([]byte, int) {
	if i == 0 {
		return []byte{0x00}, 1
	}
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	for i, bt := range b {
		if bt != 0x00 {
			return b[i:], 8 - i
		}
	}
	return b, 8
}

func computeBlockKey(sigma Block, blockNum uint64) Block {
	buf, _ := Uint64toBytes(blockNum)
	h := keccak256(sigma[:], buf)
	var b Block
	copy(b[:], h)
	return b
}

// Decrypt decrypts the given message using the given epoch secret key.
func (m *EncryptedMessage) Decrypt(epochSecretKey *EpochSecretKey) ([]byte, error) {
	sigma := m.Sigma(epochSecretKey)
	decryptedBlocks := DecryptBlocks(m.C3, sigma)
	message, err := UnpadMessage(decryptedBlocks)
	if err != nil {
		return []byte{}, err
	}

	r := computeR(sigma, message)
	expectedC1 := computeC1(r)
	if !EqualG2(m.C1, expectedC1) {
		return []byte{}, errors.New("invalid C1 in encrypted message")
	}

	return message, nil
}

// Sigma computes the sigma value of the encrypted message given the epoch secret key.
func (m *EncryptedMessage) Sigma(epochSecretKey *EpochSecretKey) Block {
	pairing := bn256.Pair((*bn256.G1)(epochSecretKey), m.C1)
	key := HashGTToBlock(pairing)
	sigma := XORBlocks(m.C2, key)
	return sigma
}

func DecryptBlocks(encryptedBlocks []Block, sigma Block) []Block {
	numBlocks := len(encryptedBlocks)
	decryptedBlocks := []Block{}
	for i := 0; i < numBlocks; i++ {
		key := computeBlockKey(sigma, uint64(i))
		decryptedBlock := XORBlocks(encryptedBlocks[i], key)
		decryptedBlocks = append(decryptedBlocks, decryptedBlock)
	}
	return decryptedBlocks
}

// PadMessage pads a message and returns it as a sequence of blocks.
// Implements PKCS #7 according to https://www.ietf.org/rfc/rfc2315.txt
func PadMessage(m []byte) []Block {
	paddingLength := BlockSize - len(m)%BlockSize
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	m = append(m, padding...)

	blocks := []Block{}
	numBlocks := len(m) / BlockSize
	for i := 0; i < numBlocks; i++ {
		var block Block
		copy(block[:], m[i*BlockSize:(i+1)*BlockSize])
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
		return nil, errors.New("invalid padding length 0")
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
