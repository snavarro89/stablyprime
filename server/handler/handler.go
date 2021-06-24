package handler

import (
	"encoding/json"

	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Handle func(params *Parameters) (interface{}, int)

type MuxResponse struct {
	Handle Handle
}

type TestResponse struct {
	Handle Handle
}

type Parameters struct {
	Body        []byte
	PathParams  map[string]string
	QueryParams map[string]string
}

func (m MuxResponse) Callback(w http.ResponseWriter, r *http.Request) {

	params := make(map[string]string)
	body, _ := ioutil.ReadAll(r.Body)

	for key, value := range r.URL.Query() {
		params[key] = value[0]
	}
	result, status := m.Handle(&Parameters{
		PathParams:  mux.Vars(r),
		QueryParams: params,
		Body:        body,
	})
	if status != http.StatusOK {
		w.WriteHeader(status)
	}

	json.NewEncoder(w).Encode(result)
}
