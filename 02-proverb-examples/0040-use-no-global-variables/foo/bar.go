package foo

import (
	"fmt"
)

type ExampleFoo struct {
	data string
}

func NewFoo() ExampleFoo {
	return ExampleFoo{
		data: loadData(),
	}
}

func (f *ExampleFoo) DoThis() {
	fmt.Println(f.data)
}

func loadData() string {
	return "data"
}
