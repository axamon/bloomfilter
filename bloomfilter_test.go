package bloomfilter_test

import (
	"fmt"

	"github.com/axamon/bloomfilter"
)

func ExampleBloomFilter_Add() {

	f := bloomfilter.New()

	f.Add("pippo4$wrwew")
	// Output:
	//
}

func ExampleBittoflip() {

	fmt.Println(bloomfilter.Bittoflip("pippo"))
	fmt.Println(bloomfilter.Bittoflip("pluto"))
	// Output:
	// 32
	// 14
}

func ExampleBloomFilter_NotExist() {

	s := "pippo"

	f := bloomfilter.New()

	f.Add(s)

	fmt.Println(f.NotExist(s))
	fmt.Println(f.NotExist("pluto676gg7"))

	// Output:
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

	n := bloomfilter.Bittoflip("pippo")

	fmt.Println(n)

	f.ShowPosBit(n)

}
