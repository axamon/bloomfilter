// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bloomfilter_test

import (
	"fmt"

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

func ExampleBittoflip() {

	fmt.Println(bloomfilter.Bittoflip("pippo"))
	fmt.Println(bloomfilter.Bittoflip("pluto"))
	// Output:
	// 32
	// 14
}

func ExampleBloomFilter_Exist() {

	s := "pippo"

	f := bloomfilter.New()

	f.Add(s)

	fmt.Println(f.Exist(s))
	fmt.Println(f.Exist("pluto"))

	f.Add("pluto")
	fmt.Println(f.Exist("pluto"))

	// Output:
	// true
	// false
	// true
}

func ExampleBloomFilter_ShowBits() {

	f := bloomfilter.New()
	f.ShowBits()

	f.Add("pippo")
	f.ShowBits()

	f.Add("pippo")
	f.ShowBits()

	f.Add("pluto")
	f.ShowBits()
	// Output:
	// 0000000000000000000000000000000000000000000000000000000000000000
	// 0000000000000000000000000000000010000000000000000000000000000000
	// 0000000000000000000000000000000010000000000000000000000000000000
	// 0000000000000010000000000000000010000000000000000000000000000000
}

func ExampleBloomFilter_ShowPosBit() {
	f := bloomfilter.New()
	f.ShowBits()

	f.Add("pippo")
	f.ShowBits()

	n := bloomfilter.Bittoflip("pippo")

	fmt.Println(n)

	f.ShowPosBit(n)
	// Output:
	// 0000000000000000000000000000000000000000000000000000000000000000
	// 0000000000000000000000000000000010000000000000000000000000000000
	// 32
	// 1
}

func ExampleBloomFilter_Exists() {

	f := bloomfilter.New()

	go f.Add("paperina")
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
