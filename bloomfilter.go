package bloomfilter

import (
	"fmt"
	//"strconv"
	"sync"

	"github.com/axamon/hashstring"
)

// BloomFilter ...
type BloomFilter struct {
	Slice uint
	sync.Mutex
}

// New creates a new *BloomFilter istance.
func New() *BloomFilter {

	return new(BloomFilter)
}

func bittoflip(s string) int {
	h := hashstring.Sha256Sum(s)

	ss := fmt.Sprintf("%x\n", h)

	var k int
	for _, v := range ss {
		k = k + int(v)
	}

	// make sure bit to flip is in the 64 bit range.
	pos := (k % 64)

	return pos
}

// Add adds an element to a *BloomFilter.
func (f *BloomFilter) Add(s string) {

	position := bittoflip(s)

	// performes OR on the position bit.
	f.Lock()
	f.Slice |= (1 << uint(position))
	f.Unlock()

	return
}

// NotExist ...
func (f *BloomFilter) NotExist(s string) bool {

	position := bittoflip(s)

	fmt.Println(f.Slice)
	if f.Slice&(1<<uint(position)) == 1 {
		return true
	}

	return false
}

// // CountOneBits ...
// func CountOneBits(i uint) int {
// 	return bits.OnesCount(i)
// }

// // SetBit ...
// func (f *BloomFilter) SetBit(nbit int) {
// 	f.Lock()
// 	defer f.Unlock()

// 	return

// }
