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
