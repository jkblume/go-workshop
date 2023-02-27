package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const MAX = 20
const MIN = 5

type algorithmServer struct {
	algorithm algorithm
}

type calculationRequest struct {
	Number int
}

func (server *algorithmServer) calculate(w http.ResponseWriter, req *http.Request) {
	var calcRequest calculationRequest
	err := json.NewDecoder(req.Body).Decode(&calcRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	calcResult, err := server.algorithm.calculate(calcRequest.Number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Result: %v", calcResult)
}

func main() {
	algoServer := algorithmServer{
		algorithm: &jamboo{},
	}

	http.HandleFunc("/calculate", algoServer.calculate)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
