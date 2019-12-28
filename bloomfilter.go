package bloomfilter

import (
	"fmt"
	"math/bits"
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

// Bittoflip finds the bit in a uint to flip as hashing.
func Bittoflip(s string) int {
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

	position := Bittoflip(s)

	// performes OR on the position bit.
	f.Lock()
	f.Slice |= (1 << uint(position))
	f.Unlock()

	return
}

// ShowBits shows the single bits in the uint.
func (f *BloomFilter) ShowBits() {

	for i := 0; i < bits.UintSize; i++ {
		// If there is a bit set at this position, write a 1.
		// ... Otherwise write a 0.
		if f.Slice&(1<<uint(i)) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()
}

// ShowPosBit shows the pos bit in the uint to verify.
func (f *BloomFilter) ShowPosBit(pos int) {

	// If there is a bit set at this position, write a 1.
	// ... Otherwise write a 0.
	if f.Slice&(1<<uint(pos)) != 0 {
		fmt.Print("1")
	} else {
		fmt.Print("0")
	}
	fmt.Println()
}

// NotExist ...
func (f *BloomFilter) NotExist(s string) bool {

	position := Bittoflip(s)

	f.ShowBits()
	if f.Slice&(1<<uint(position)) == 1 {
		return true
	}

	return false
}
