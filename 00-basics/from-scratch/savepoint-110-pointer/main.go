package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 20
	result := fizzbuzzPointer(&number)
	fmt.Println(result)
}

func fizzbuzzPointer(number *int) string {
	numberValue := *number
	return fizzbuzz(numberValue)
}

func fizzbuzz(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else if number%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(number)
}
