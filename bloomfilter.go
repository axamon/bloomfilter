// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bloomfilter is a thread-safe probabilisic filter.
// With default configuration it can store up to one million elements.
// All negative results to queries are meant to be 100% accurate, while false
// positives should happen less than once every three million queries.
package bloomfilter

import (
	"sync"

	"github.com/axamon/hashstring"
)

/*
Bloom filter probability is:
	p = pow(1 - exp(-k / (m / n)), k)
*/

// Elements is the MAX number of elements to store in the filter.
var Elements = 1000000 // n

var filterLength = 100 * Elements // m

// NumOfHashes is the number of hash functions to use.
var NumOfHashes = 5 // k

// New creates a new *BloomFilter istance.
func New() *BloomFilter {
	f := &BloomFilter{
		M: make(map[int]struct{}, filterLength),
	}
	return f
}

// Nhashings creates in parallel the list of nums to insert
// in the filter based on n trivial different hashing functions.
func Nhashings(s string, n int) []int {

	var b Bits

	var wgN sync.WaitGroup

	for i := 1; i <= n; i++ {
		wgN.Add(1)
		go func(i int) {
			defer wgN.Done()
			h := hashstring.Sha512Sum(s + string(i))
			// ss := fmt.Sprintf("%x\n", h)

			var k int
			for index, v := range h {
				k += int(v) * index * i
			}
			// make sure bit to flip is in the 64 bit range.
			pos := (k % filterLength)
			b.Lock()
			b.list = append(b.list, pos)
			b.Unlock()
		}(i)
	}
	wgN.Wait()

	return b.list
}

// Add adds an element to a *BloomFilter.
func (f *BloomFilter) Add(s string) {

	positions := Nhashings(s, NumOfHashes)

	for _, bit := range positions {
		f.Lock()
		f.M[bit] = struct{}{}
		f.Unlock()
	}

	return
}

// Exists returns true false with 100% certainty if element is NOT in the filter,
// returns true with LESS than 100% certainty if element seems to be present.
func (f *BloomFilter) Exists(s string) bool {
	f.Lock()
	defer f.Unlock()
	positions := Nhashings(s, NumOfHashes)

	for _, bit := range positions {
		if _, ok := f.M[bit]; !ok {
			return false
		}
	}
	return true
}
