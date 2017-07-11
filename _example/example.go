// Copyright 2017 Benjamin 'Benno' Falkner. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"ean13"
	"fmt" 
	"strconv"
	"os"
)

func main () {
	fmt.Println("EAN13:")
	if len(os.Args) < 2 { fmt.Println("Not enough Arguments"); os.Exit(1)}
	fmt.Println("Input:", os.Args[1])
	// read Argument as int64
	in, err := strconv.Atoi(os.Args[1])
	if err != nil {fmt.Println(err); os.Exit(1)}

	//Create ean in private section
	r := ean13.Encode(int64(in), ean13.PRIVATE)
	fmt.Println("Output:", r)

	// Validate result
	if r.Validate() {fmt.Println("VALID")} else {fmt.Println("INALID")}
}