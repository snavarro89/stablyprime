package prime

import (
	"net/http"
	"strconv"

	A "github.com/snavarro89/stablyprime/app"
	H "github.com/snavarro89/stablyprime/handler"
)

var app A.App

func Data(a A.App) {
	app = a
}

func Mux(w http.ResponseWriter, r *http.Request) {
	m := H.MuxResponse{
		Handle: func(params *H.Parameters) (interface{}, int) {
			return handler(*params)
		},
	}
	m.Callback(w, r)
}

type response struct {
	PrimeNumber int `json:"primeNumber"`
}

func handler(params H.Parameters) (interface{}, int) {

	number, err := strconv.Atoi(params.PathParams["number"])
	if err != nil {
		errResponse := map[string]interface{}{
			"error": true,
			"desc":  "Please provide a valid number",
			"code":  "1000",
		}
		return errResponse, http.StatusBadRequest
	}
	prime, _ := app.Data.GetPrime(number)

	response := response{
		PrimeNumber: prime,
	}
	return response, http.StatusOK
}
