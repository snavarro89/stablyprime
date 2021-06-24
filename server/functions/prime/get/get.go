package prime

import (
	"net/http"

	A "github.com/snavarro89/stablyprime/api/numbers"
	H "github.com/snavarro89/stablyprime/handler"
)

func Mux(w http.ResponseWriter, r *http.Request) {
	m := H.MuxResponse{
		Handle: func(params *H.Parameters) (interface{}, int) {
			return handler(*params)
		},
	}
	m.Callback(w, r)
}

type Response struct {
	PrimeNumber int `json:"primeNumber"`
}

func handler(params H.Parameters) (interface{}, int) {

	prime := A.GetPrime(0)

	response := Response{
		PrimeNumber: prime,
	}
	return response, http.StatusOK
}
