package algorithms

import (
	"fmt"
	"strconv"
)

type Fizzbuzz struct {
	Min int
	Max int
}

func (f *Fizzbuzz) Calculate(number int) (string, error) {
	if number < f.Min || number > f.Max {
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

type Jamboo struct {
}

func (f *Jamboo) Calculate(number int) (string, error) {
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
