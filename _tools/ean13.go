// Copyright 2017 Benjamin 'Benno' Falkner. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"ean13"
	"flag"
	"fmt"
	"os"
)

var (
	start = flag.Int64("s", 20000000000, "Start with number (don't set check digit).")
	count = flag.Int64("c", 1, "Number of EAN13 to generate.")
	name = flag.String("o", "", "Filename for output (default is STDOUT).")
)

func main() {
	fmt.Println("Certificate")
	fmt.Println("Copyright 2017 Benjamin 'Benno' Falkner. All rights reserved.")
	flag.Parse()
	
	var out *bufio.Writer
	if *name == "" {
		out = bufio.NewWriter(os.Stdout) 
	} else {
		f, err := os.Create(*name)
		if err != nil {fmt.Println(err); os.Exit(1)}
		out = bufio.NewWriter(f)
	}

	for i := int64(0); i < (*count) ; i++ {
		e,err := ean13.Encode(*start+i*10,ean13.NULL)
		if err != nil {fmt.Println(err); os.Exit(1)}
		out.WriteString(e.String())
		out.Write([]byte("\r\n"))
	}
	out.Flush()
}