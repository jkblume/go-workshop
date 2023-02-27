package main

import (
	"fmt"
	algorithms "jblume.dev/go-beginner-workshop/example-project/internal"
	"net/http"
	"os"
)

const (
	MIN = 1
	MAX = 20
)

func main() {
	mux := algorithms.SetupMux(&algorithms.Fizzbuzz{
		Min: MIN,
		Max: MAX,
	})

	addr := ":8090"
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
