package algorithms

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_FizzbuzzHandler(t *testing.T) {
	algo := &Fizzbuzz{
		Min: 1,
		Max: 20,
	}
	handler := SetupMux(algo)

	type args struct {
		req *http.Request
	}

	tests := []struct {
		name     string
		args     func(t *testing.T) args
		wantCode int
		wantBody string
	}{
		{
			name: "Calculate:Fizz",
			args: func(*testing.T) args {
				r := calculationRequest{
					Number: 7,
				}

				buffer := &bytes.Buffer{}
				err := json.NewEncoder(buffer).Encode(r)
				if err != nil {
					t.Fatalf("failed to encode request body: %s", err.Error())
				}

				req, err := http.NewRequest("POST", "/calculate", buffer)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				return args{
					req: req,
				}
			},
			wantCode: http.StatusOK,
			wantBody: "Result: 7",
		},
		{
			name: "Calculate:Error",
			args: func(*testing.T) args {
				r := calculationRequest{
					Number: 100,
				}

				buffer := &bytes.Buffer{}
				err := json.NewEncoder(buffer).Encode(r)
				if err != nil {
					t.Fatalf("failed to encode request body: %s", err.Error())
				}

				req, err := http.NewRequest("POST", "/calculate", buffer)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				return args{
					req: req,
				}
			},
			wantCode: http.StatusBadRequest,
			wantBody: "number is out of range\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, tArgs.req)

			if resp.Result().StatusCode != tt.wantCode {
				t.Fatalf("the status code should be [%d] but received [%d]", resp.Result().StatusCode, tt.wantCode)
			}

			if resp.Body.String() != tt.wantBody {
				t.Fatalf("the response body should be [%s] but received [%s]", resp.Body.String(), tt.wantBody)
			}

		})
	}
}
