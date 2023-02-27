package main

import "os"

func foo() func() error {
	f, _ := os.Open("/path")
	return f.Close
}

func bar() *os.File {
	f, _ := os.Open("/path")
	return f
}
