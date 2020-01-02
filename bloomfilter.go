// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bloomfilter is based on only one hashing funcion
// and with only 64 bits of memory.
// That is enough to speed up searches for up to 10 elements max.
package bloomfilter

import (
	"fmt"
	"sync"

	"github.com/axamon/hashstring"
)

const filterLength int = 10000

var numOfHashes = 14

// BloomFilter ...
type BloomFilter struct {
	M map[int]struct{}
	sync.Mutex
}

// New creates a new *BloomFilter istance.
func New() *BloomFilter {
	f := &BloomFilter{
		M: map[int]struct{}{},
	}
	return f
}

// Nhashings ...
func Nhashings(s string, n int) []int {

	var bits []int

	for i := 0; i < n; i++ {
		h := hashstring.Sha256Sum(s + string(i))
		ss := fmt.Sprintf("%x\n", h)

		var k int
		for index, v := range ss {
			k += int(v) * index
		}

		// make sure bit to flip is in the 64 bit range.
		pos := (k % filterLength)
		bits = append(bits, pos)
	}

	return bits
}

// Add adds an element to a *BloomFilter.
func (f *BloomFilter) Add(s string) {

	positions := Nhashings(s, numOfHashes)

	for _, bit := range positions {
		f.Lock()
		f.M[bit] = struct{}{}
		f.Unlock()
	}

	return
}

// Exists ...
func (f *BloomFilter) Exists(s string) bool {

	positions := Nhashings(s, numOfHashes)

	for _, bit := range positions {
		if _, ok := f.M[bit]; !ok {
			return false
		}
	}
	return true
}
