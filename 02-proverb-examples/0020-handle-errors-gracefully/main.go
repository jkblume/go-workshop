package main

import (
	foo "0020-handle-errors-gracefully/internal"
	"errors"
	"fmt"
)

func main() {
	attempt3()
}

func attempt1() {
	example, err := foo.GetExample(true)
	if err != nil {
		fmt.Printf("Error occured: %s\n", err)

		return
	}
	fmt.Printf("Result: %s", example)
}

func attempt2() {
	example, err := foo.GetExample2(true)
	if err != nil {
		fmt.Printf("Error occured: %s\n", err)
		fmt.Printf("Error Type is: %T\n", err)

		return
	}
	fmt.Printf("Result: %s", example)
}

func attempt3() {
	example, err := foo.GetExample3(true)
	if err != nil {
		fmt.Printf("Error occured: %s\n", err)
		fmt.Printf("Error Type is: %T\n", err)

		err = errors.Unwrap(err)
		fmt.Printf("Error occured: %s\n", err)
		fmt.Printf("Error Type is: %T\n", err)

		err = errors.Unwrap(err)
		fmt.Printf("Error occured: %s\n", err)
		fmt.Printf("Error Type is: %T\n", err)

		return
	}
	fmt.Printf("Result: %s", example)
}
