package cuckoo

import (
	"math/bits"

	"github.com/dgryski/go-metro"
)

// represents the available size of slots at each index
const bucketSize = 4
const hashSeed uint64 = 1337
const maxTries = 500

// instead of storing the data we store the fingerprint
type fingerprint byte
type bucket [bucketSize]fingerprint

type Filter struct {
	buckets   []bucket // 2 Dimensional structure
	count     uint     // total number of elements
	bucketPow uint
}

// cache the hash, to avoid recomputation

// fingerprint is of 1 byte i.e. 8 bit so we calculate hash of every possible fingerprint 2^8 possibilities hence 256 sized array
var altHash = [256]uint{}

// mask is used to  fetch the last n bits of a hash value.e, code calculates a mask as 1 << i - 1, which generates a mask with i bits set to 1. This mask is then used to extract a specific number of bits from a hash value.
var masks = [65]uint{}

func init() {
	for i := 0; i < 256; i += 1 {
		altHash[i] = uint(metro.Hash64([]byte{byte(i)}, hashSeed))
	}
	for i := uint(0); i <= 64; i += 1 {
		masks[i] = (1 << i) - 1
	}
}

func New(capacity uint) *Filter {
	bucketCapacity := getNextPow2(uint64(capacity)) / bucketSize
	if bucketCapacity < 1 {
		bucketCapacity = 1
	}
	buckets := make([]bucket, capacity)
	return &Filter{
		buckets:   buckets,
		count:     0,
		bucketPow: uint(bits.TrailingZeros(capacity)),
	}
}

func (cf *Filter) Insert(data []byte) bool {
	pi, fp := getIndexAndFingerprint(data, cf.bucketPow)
	if cf.primaryInsert(fp, pi) {
		return true
	}

	si := getAltIndex(fp, pi, cf.bucketPow)
	if cf.primaryInsert(fp, si) {
		return true
	}

	index := pickRandomIndex(pi, si)
	return cf.secondaryInsert(fp, index)
}

func (cf *Filter) Lookup(data []byte) bool {
	pi, fp := getIndexAndFingerprint(data, cf.bucketPow)
	if cf.buckets[pi].getFingerprintIndex(fp) != -1 {
		return true
	}
	si := getAltIndex(fp, pi, cf.bucketPow)
	return cf.buckets[si].getFingerprintIndex(fp) != -1
}
