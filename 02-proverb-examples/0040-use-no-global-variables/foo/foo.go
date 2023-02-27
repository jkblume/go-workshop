package foo

import (
	"fmt"
)

var (
	data string = ""
)

func init() {
	data = "data"
}

func DoSomething() {
	fmt.Println(data)
}
