package main

import (
	"fmt"
	"jblume.dev/go-beginner-workshop/example-project/internal/algorithms"
	"net/http"
	"os"
)

const (
	MIN = 1
	MAX = 20
)

func main() {
	algoServer := algorithms.AlgorithmServer{
		Algorithm: &algorithms.Fizzbuzz{
			Min: MIN,
			Max: MAX,
		},
	}

	http.HandleFunc("/calculate", algoServer.Calculate)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
