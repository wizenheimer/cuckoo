package cuckoo

const emptyfp = 0

// This function is used to insert fingerprint into first available slot along the bucket
func (b *bucket) insert(fp fingerprint) bool {
	for i, tp := range b {
		// incase the slot is empty
		if tp == emptyfp {
			b[i] = fp
			return true
		}
	}
	return false
}

// This function is used to delete fingerprint along the bucket
func (b *bucket) delete(fp fingerprint) bool {
	for i, cp := range b {
		// incase the fingerprint is found
		if cp == fp {
			b[i] = emptyfp
			return true
		}
	}
	return false
}

// This function is used to reset fingerprint along the bucket
func (b *bucket) _() {
	for i := range b {
		b[i] = emptyfp
	}
}

// Get the index associated with fingerprint
func (b *bucket) getFingerprintIndex(fp fingerprint) int {
	for i, cf := range b {
		if cf == fp {
			return i
		}
	}
	return -1
}
