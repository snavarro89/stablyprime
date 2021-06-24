package prime

import (
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
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

func Aws(w http.ResponseWriter, r *http.Request) {

	a := H.AwsResponse{
		Handle: func(params *H.Parameters) (interface{}, int) {
			return handler(*params)
		},
		Headers: H.Headers{
			Origin:  "*",
			Methods: "GET,DELETE",
			Headers: "Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Allow-Origin,Content-Type,Authorization,X-Amz-Date,X-Api-Key,X-Amz-Security-Token",
		},
	}
	lambda.Start(a.Callback)
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
	prime, err := app.Data.GetPrime(number)

	if err != nil {
		errResponse := map[string]interface{}{
			"error": true,
			"desc":  err.Error(),
			"code":  "1001",
		}
		return errResponse, http.StatusBadRequest
	}

	response := response{
		PrimeNumber: prime,
	}
	return response, http.StatusOK
}
