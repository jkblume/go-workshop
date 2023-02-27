package main

import (
	"fmt"
	"strconv"
)

func main() {

	numbers := [4]int{1, 2, 3, 4}

	for _, val := range numbers {
		result := fizzbuzz(val)

		fmt.Println(result)
	}
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
