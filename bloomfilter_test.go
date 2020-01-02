// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bloomfilter_test

import (
	"fmt"
	"sync"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/axamon/bloomfilter"
)

func ExampleBloomFilter_Add() {

	f := bloomfilter.New()

	fmt.Println(f.Exists("pluto"))

	f.Add("pluto")
	fmt.Println(f.Exists("pluto"))
	// Output:
	// false
	// true
}

func ExampleBloomFilter_Exists() {

	f := bloomfilter.New()

	n := 10000
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			f.Add(randomdata.FirstName(randomdata.Male))
			f.Add(randomdata.FirstName(randomdata.Female))
		}()
	}
	wg.Wait()

	f.Add("pluto")
	f.Add("pippo")

	fmt.Println(f.Exists("pippo"))
	fmt.Println(f.Exists("pluto"))
	fmt.Println(f.Exists("minnie"))
	fmt.Println(f.Exists("paperino"))
	// Output:
	// true
	// true
	// false
	// false
}
