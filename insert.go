package cuckoo

import "math/rand"

// Checks if there is an available slot in the bucket at index i. If a slot is available, it inserts the fingerprint and increments the count of elements in the filter.
func (cf *Filter) primaryInsert(fp fingerprint, index uint) bool {
	currentBucket := cf.buckets[index]
	if currentBucket.insert(fp) {
		cf.count += 1
		return true
	}
	return false
}

// This function is used when both primary and alternate insertions fail, and the filter enters a reinsertion process. The function iteratively moves fingerprints around in an attempt to make space for the new insertion. The process hard stops when maxTries is reached.
func (cf *Filter) secondaryInsert(fp fingerprint, index uint) bool {
	for k := 0; k < maxTries; k += 1 {
		// pick a random position in the bucket
		j := rand.Intn(bucketSize)
		// rearrange the fingerprints
		oldfp := fp
		fp = cf.buckets[index][j]
		cf.buckets[index][j] = oldfp
		// get new index of the removed fingerprint
		index = getAltIndex(fp, index, cf.bucketPow)
		if cf.primaryInsert(fp, index) {
			return true
		}
	}
	return false
}
