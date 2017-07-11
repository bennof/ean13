// Copyright 2017 Benjamin 'Benno' Falkner. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ean13

import (
	"errors"
	"math"
	"strconv"
)

const (
	MAX     EAN13 = 9999999999999 // Maximal value for EAN13
	PRIVATE EAN13 = 2000000000000 // Private numbers
	BOOKS   EAN13 = 9700000000000 // Printed work
	ISBN1   EAN13 = 9780000000000 // Area 1 for ISBN
	ISBN2   EAN13 = 9790000000000 // Area 2 for ISBN
	ISSN    EAN13 = 9770000000000 // ISSN
	NULL    EAN13 = 0000000000000 // No Number
)

var (
	NaN           = EAN13(math.NaN()) // NOT A Number
	POWERS_OF_TEN = []int64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000, 100000000000, 1000000000000}

	ERROR_INVALID         = errors.New("EAN is invalid.")
	ERROR_PREFIX_MISMATCH = errors.New("EAN prefix mismatch.")
	ERROR_OUT_OF_RANGE    = errors.New("EAN is out of range.")
	ERROR_CREATE_FAILED   = errors.New("EAN not created.")
)

// returns the nth diget of a int64 base 10
func nthdigit(x int64, n int) int64 {
	return ((x / POWERS_OF_TEN[n]) % 10)
}

// calculates the Checksum defined in GS1 General Specifications
func checksum(x int64) int64 {
	var sum int64
	var multip int64 = 1
	for i := 12; i > 0; i-- {
		n := nthdigit(x, i)
		//log.Println(n,"*",multip)
		sum += n * multip
		if multip == 3 {
			multip = 1
		} else {
			multip = 3
		}
	}
	sum_ := sum + 9
	sum_ = sum_ - (sum_ % 10)
	return sum_ - sum
}

// type for EAN13 hiding an int64
type EAN13 int64

// validate an ean13 using its last digit
func (e EAN13) Validate() bool {
	if e > MAX || e < 0 {
		return false
	}

	s := checksum(int64(e))
	// log.Println("Checksum",s,nthdigit(int64(e),0))
	return s == nthdigit(int64(e), 0)
}

// Converts EAN13 to a string
func (e EAN13) String() string {
	return strconv.FormatInt(int64(e), 10)
}

// check for prefix
func (e EAN13) Is(pre EAN13, n int) bool {
	var r bool = true
	for i := 12; i > 13-n; i-- {
		r = r && nthdigit(int64(e), i) == nthdigit(int64(pre), i)
	}
	return r
}

func (e EAN13) Decode(prefix EAN13, n int) (int64, error) {
	if !e.Is(prefix,n) {return 0, ERROR_PREFIX_MISMATCH}
	i := int64(e-prefix)
	i /=10
	return i, nil
} 

// Create a new ean from string
func New(ean string) (EAN13, error) {
	in, err := strconv.Atoi(ean)
	if err != nil {
		return NaN, err
	}
	e := EAN13(in)
	if !e.Validate() {
		return e, ERROR_INVALID
	}
	return e, nil
}

// Create a new ean using a prefix (const PRIVATE i.g.) and a
// number. Return a valid EAN13 or NaN
func Encode(in int64, prefix EAN13) (EAN13, error) {
	e := EAN13(in*10) + prefix
	if e > MAX {
		return NaN, ERROR_OUT_OF_RANGE
	}
	s := checksum(int64(e))
	e += EAN13(s)
	if !e.Validate() {
		return NaN, ERROR_CREATE_FAILED
	}
	return e, nil
}

