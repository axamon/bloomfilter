package bloomfilter_test

import (
	"github.com/axamon/bloomfilter"
)

func ExampleHashToInt() {

	bloomfilter.HashToInt("pippo4$wrwew")
	bloomfilter.HashToInt("pippo$wrwew")

	// Output:
	//
}
