package main

import (
	"fmt"
	"os"
	"strconv"
)

const MAX = 20
const MIN = 5

func main() {

	number := 10
	fizzObj := fizzbuzz{
		min: MIN,
		max: MAX,
	}

	result, err := fizzObj.calculate(number)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}

type fizzbuzz struct {
	min int
	max int
}

func (f *fizzbuzz) calculate(number int) (string, error) {
	if number < f.min || number > f.max {
		return "", fmt.Errorf("number is out of range")
	}

	result := strconv.Itoa(number)
	if number%3 == 0 && number%5 == 0 {
		result = "FizzBuzz"
	} else if number%3 == 0 {
		result = "Fizz"
	} else if number%5 == 0 {
		result = "Buzz"
	}

	return result, nil
}
