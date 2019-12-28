package bloomfilter_test

import (
	"github.com/axamon/bloomfilter"
)

func ExampleHashToInt() {

	bloomfilter.HashToInt("pippo4$wrwew")
	// Output:
	// 0000000000000000000010000000000000000000000000000000000000000000
}
