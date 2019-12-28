package bloomfilter

import (
	"fmt"
	"math/bits"
	//"strconv"
	"sync"

	"github.com/axamon/hashstring"
)

// BloomFilter ...
type BloomFilter struct {
	Slice uint
	sync.Mutex
}

// New ...
func New() *BloomFilter {

	var filter = new(BloomFilter)
	filter.Slice = uint(1000000)

	return filter
}

// HashToInt ...
func HashToInt(s string) int {

	f := New()

	
	h := hashstring.Md5Sum(s)

	ss := fmt.Sprintf("%x\n", h)

	var k int
	for _, v := range ss {
		k = k + int(v)
	}

	//n, _ := strconv.Atoi(k)

	f.Slice |= (1 << uint(k))

	fmt.Println(n, f.Slice)


	fmt.Println()
	// fmt.Println(h, n)
	//fmt.Println(l, int(v), k)

	// value := uint(1000000)
	// // Loop over all bits in the uint.
	for i := 0; i < bits.UintSize; i++ {
		// If there is a bit set at this position, write a 1.
		// ... Otherwise write a 0.
		if f.Slice&(1<<uint(i)) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		//value |= (1 << uint(30))
	}

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

	return

}
