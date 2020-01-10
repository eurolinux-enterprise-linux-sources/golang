// run

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var c chan int

func main() {
	c = make(chan int);
	go func() { c <- 0 } ();
	<-c
}
