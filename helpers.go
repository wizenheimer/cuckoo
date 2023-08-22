package cuckoo

import (
	"math/rand"

	"github.com/dgryski/go-metro"
)

func getNextPow2(n uint64) uint {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++
	return uint(n)
}

func getAltIndex(fp fingerprint, i uint, bucketPow uint) uint {
	mask := masks[bucketPow]
	hash := altHash[fp] & mask
	return (i & mask) ^ hash
}

// least significant bits for fingerprint.
func getFingerprint(hash uint64) byte {
	fp := byte(hash%255 + 1)
	return fp
}

// returns the 2 bucket indices and fingerprint to be used
func getIndexAndFingerprint(data []byte, bucketPow uint) (uint, fingerprint) {
	hash := metro.Hash64(data, 1337)
	fp := getFingerprint(hash)
	// most significant bits for deriving index.
	i1 := uint(hash>>32) & masks[bucketPow]
	return i1, fingerprint(fp)
}

// selects a random index out of primary index and secondary index
func pickRandomIndex(pi uint, si uint) uint {
	if rand.Intn(2) == 0 {
		return pi
	}
	return si
}
