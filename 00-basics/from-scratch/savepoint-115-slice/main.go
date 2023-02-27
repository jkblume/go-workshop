package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := fizzbuzzSlize(20)
	fmt.Println(result)
}

func fizzbuzzSlize(number int) []string {
	var result []string
	for i := 1; i <= number; i++ {
		singleResult := fizzbuzz(i)
		result = append(result, singleResult)
	}
	return result
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

func r(mySlice []int) {

}
