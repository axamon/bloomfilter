package bloomfilter

import "sync"

// BloomFilter is the thread-safe golang version of a bloom filter.
type BloomFilter struct {
	M map[int]struct{}
	sync.Mutex
}

// Bits is the list of numbers to insert in the bloom filter.
type Bits struct {
	list []int
	sync.Mutex
}
