// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashstring

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// Md5Sum returns the md5sum in exadecimal format of the string s.
func Md5Sum(s string) string {
	h := md5.New()

	// new line is added at the end to obtain same output of
	// common md5sum cli tools.
	h.Write([]byte(s + "\n"))
	hashedpass := h.Sum(nil)

	// output is in exadecimal format.
	return fmt.Sprintf("%x", hashedpass)
}

// Sha256Sum returns the md5sum in exadecimal format of the string s.
func Sha256Sum(s string) string {
	h := sha256.New()

	// new line is added at the end to obtain same output of
	// common md5sum cli tools.
	h.Write([]byte(s + "\n"))
	hashedpass := h.Sum(nil)

	// output is in exadecimal format.
	return fmt.Sprintf("%x", hashedpass)
}

// Sha512Sum returns the md5sum in exadecimal format of the string s.
func Sha512Sum(s string) string {
	h := sha512.New()

	// new line is added at the end to obtain same output of
	// common md5sum cli tools.
	h.Write([]byte(s + "\n"))
	hashedpass := h.Sum(nil)

	// output is in exadecimal format.
	return fmt.Sprintf("%x", hashedpass)
}
