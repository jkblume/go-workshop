package main

import (
	"fmt"
	"io"
)

func main() {
	t := T{data: "hello"}
	result, err := io.ReadAll(t)
	fmt.Println(err)
	fmt.Println(string(result))
}

type T struct {
	data    string
	counter int
}

var (
	calledOnce bool
)

func (t T) Read(buf []byte) (n int, err error) {
	fmt.Println(t.counter)
	if calledOnce {
		return 0, io.EOF
	}
	n = copy(buf, t.data)
	calledOnce = true
	t.counter++
	return
}
