package bloomfilter

import (
	"bytes"
	"math/bits"
	"sync"

	"github.com/axamon/hashstring"
)

// BloomFilter ...
type BloomFilter struct {
	Buf *bytes.Buffer
	sync.Mutex
}

// New ...
func New() *BloomFilter {

	var filter = new(BloomFilter)

	filter.Buf = bytes.NewBuffer(make([]byte, 5000000))

	return filter
}

// HashToInt ...
func HashToInt(s string) int {

	hashstring.Md5Sum(s)

	return 0
}

// CountOneBits ...
func CountOneBits(i uint) int {
	return bits.OnesCount(i)
}

// SetBit ...
func (f *BloomFilter) SetBit(nbit int) {
	f.Lock()
	defer f.Unlock()
	1

}
