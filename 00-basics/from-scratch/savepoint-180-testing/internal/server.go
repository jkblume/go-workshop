package algorithms

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type algorithm interface {
	Calculate(number int) (string, error)
}

type AlgorithmServer struct {
	Algorithm algorithm
}

type calculationRequest struct {
	Number int
}

func (server *AlgorithmServer) Calculate(w http.ResponseWriter, req *http.Request) {
	var calcRequest calculationRequest
	err := json.NewDecoder(req.Body).Decode(&calcRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	calcResult, err := server.Algorithm.Calculate(calcRequest.Number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Result: %v", calcResult)
}

func SetupMux(algo algorithm) *http.ServeMux {
	server := AlgorithmServer{
		Algorithm: algo,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/calculate", server.Calculate)

	return mux
}
