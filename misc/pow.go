package misc

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
)

const HashZeroBytesLength = 3

type message struct {
	payload     string
	payloadHash [32]byte
	powHash     [32]byte
	randNum     int64
}

func NewMessage(payload string) *message {
	msg := new(message)
	msg.payload = payload
	msg.randNum = rand.Int63()
	msg.payloadHash = sha256.Sum256([]byte(payload))
	return msg
}

func (b *message) isValidHash(hash [32]byte) bool {
	var zeroBytesCounter int
	for i := 0; i < HashZeroBytesLength; i++ {
		if hash[i] == 0 {
			zeroBytesCounter++
		}
	}
	if zeroBytesCounter == HashZeroBytesLength {
		return true
	} else {
		return false
	}
}

func (b *message) calcHash() [32]byte {
	// Temporarily buffers for powHash calculation
	hash2Buf := make([]byte, 64)
	randNumBytes := make([]byte, 8)
	// 0x1 => 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x1
	binary.LittleEndian.PutUint64(randNumBytes, uint64(b.randNum))
	randNumHash := sha256.Sum256(randNumBytes)
	// Concatenate random number powHash with the payload powHash to calculate the target one.
	copy(hash2Buf, randNumHash[:])
	copy(hash2Buf[32:], b.payloadHash[:])
	calcHash := sha256.Sum256(hash2Buf[:])

	return calcHash
}

// doWork finds a powHash value with N bytes equal to zero using the following computation:  H(H(randNum) + Hash(payload))
func (b *message) doWork() {
	var found bool

	for !found {
		calcHash := b.calcHash()
		if b.isValidHash(calcHash) {
			found = true
			b.powHash = calcHash
		} else {
			b.randNum++
		}
	}
}

func (b *message) verifyWork() bool {
	calcHash := b.calcHash()

	return bytes.Compare(calcHash[:], b.powHash[:]) == 0
}
