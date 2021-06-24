package prime

import (
	"encoding/json"
	"errors"
	"testing"

	A "github.com/snavarro89/stablyprime/app"
	H "github.com/snavarro89/stablyprime/handler"
)

type NumbersModel struct {
	PrimeModel PrimeModel
}

type PrimeModel struct {
	ExpectedNumber int
	ExpectedError  error
}

func (nm NumbersModel) GetPrime(number int) (int, error) {
	return nm.PrimeModel.ExpectedNumber, nm.PrimeModel.ExpectedError
}

func TestZero(t *testing.T) {
	application := A.App{
		Data: A.Data{
			Numbers: NumbersModel{
				PrimeModel: PrimeModel{
					ExpectedNumber: 0,
					ExpectedError:  nil,
				},
			},
		},
	}
	Data(application)
	params := map[string]string{}

	vars := map[string]string{
		"number": "0",
	}
	emp := make(map[string]interface{})
	body, _ := json.Marshal(emp)

	var parameters = &H.Parameters{
		PathParams:  vars,
		QueryParams: params,
		Body:        body,
	}

	result, status := handler(*parameters)
	byteData, _ := json.Marshal(result)
	var res response
	json.Unmarshal(byteData, &res)

	expectStatus := 200
	if status != expectStatus {
		t.Fatalf(`handler(object) = %d, want match for %d`, status, expectStatus)
	}
	expectPrime := 0
	if res.PrimeNumber != expectPrime {
		t.Fatalf(`handler(object) = %d, want match for %d`, res.PrimeNumber, expectPrime)
	}
}

func TestNegative(t *testing.T) {
	application := A.App{
		Data: A.Data{
			Numbers: NumbersModel{
				PrimeModel: PrimeModel{
					ExpectedNumber: 0,
					ExpectedError:  errors.New("Only positive numbers allowed"),
				},
			},
		},
	}
	Data(application)
	params := map[string]string{}

	vars := map[string]string{
		"number": "-2",
	}
	emp := make(map[string]interface{})
	body, _ := json.Marshal(emp)

	var parameters = &H.Parameters{
		PathParams:  vars,
		QueryParams: params,
		Body:        body,
	}

	result, status := handler(*parameters)
	byteData, _ := json.Marshal(result)
	var res map[string]interface{}
	json.Unmarshal(byteData, &res)

	expectStatus := 400
	if status != expectStatus {
		t.Fatalf(`handler(object) = %d, want match for %d`, status, expectStatus)
	}
	//We could also test for the error code and the boolean error
	expectError := "Only positive numbers allowed"
	if res["desc"].(string) != expectError {
		t.Fatalf(`handler(object) = %s, want match for %s`, res["desc"].(string), expectError)
	}
}

func TestValidNumber(t *testing.T) {
	application := A.App{
		Data: A.Data{
			Numbers: NumbersModel{
				PrimeModel: PrimeModel{
					ExpectedNumber: 53,
					ExpectedError:  nil,
				},
			},
		},
	}
	Data(application)
	params := map[string]string{}

	vars := map[string]string{
		"number": "55",
	}
	emp := make(map[string]interface{})
	body, _ := json.Marshal(emp)

	var parameters = &H.Parameters{
		PathParams:  vars,
		QueryParams: params,
		Body:        body,
	}

	result, status := handler(*parameters)
	byteData, _ := json.Marshal(result)
	var res response
	json.Unmarshal(byteData, &res)

	expectStatus := 200
	if status != expectStatus {
		t.Fatalf(`handler(object) = %d, want match for %d`, status, expectStatus)
	}
	expectPrime := 53
	if res.PrimeNumber != expectPrime {
		t.Fatalf(`handler(object) = %d, want match for %d`, res.PrimeNumber, expectPrime)
	}
}

func TestNonNumericInput(t *testing.T) {
	application := A.App{
		Data: A.Data{
			Numbers: NumbersModel{
				PrimeModel: PrimeModel{
					ExpectedNumber: 0,
					ExpectedError:  nil,
				},
			},
		},
	}
	Data(application)
	params := map[string]string{}

	vars := map[string]string{
		"number": "aaa",
	}
	emp := make(map[string]interface{})
	body, _ := json.Marshal(emp)

	var parameters = &H.Parameters{
		PathParams:  vars,
		QueryParams: params,
		Body:        body,
	}

	result, status := handler(*parameters)
	byteData, _ := json.Marshal(result)
	var res map[string]interface{}
	json.Unmarshal(byteData, &res)

	expectStatus := 400
	if status != expectStatus {
		t.Fatalf(`handler(object) = %d, want match for %d`, status, expectStatus)
	}
	//We could also test for the error code and the boolean error
	expectError := "Please provide a valid number"
	if res["desc"].(string) != expectError {
		t.Fatalf(`handler(object) = %s, want match for %s`, res["desc"].(string), expectError)
	}
}
