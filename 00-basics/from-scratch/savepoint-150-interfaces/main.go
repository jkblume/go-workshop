package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const MAX = 20
const MIN = 5

func main() {
	http.ListenAndServe(":8090", nil)
}

type algorithm interface {
	calculate(number int) (string, error)
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

type jamboo struct {
}

func (f *jamboo) calculate(number int) (string, error) {
	result := strconv.Itoa(number)
	if number%3 == 0 && number%5 == 0 {
		result = "JamBoo"
	} else if number%3 == 0 {
		result = "Jam"
	} else if number%5 == 0 {
		result = "Boo"
	}

	return result, nil
}
