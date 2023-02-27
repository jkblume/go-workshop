package main

import (
	"fmt"
	algorithms "jblume.dev/go-beginner-workshop/example-project/internal"
	"net/http"
	"os"
)

func main() {
	mux := algorithms.SetupMux(&algorithms.Jamboo{})

	addr := ":8091"
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
