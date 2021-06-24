package handler

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

type Handle func(params *Parameters) (interface{}, int)

type MuxResponse struct {
	Handle Handle
}

type Parameters struct {
	Body []byte
}

func (m MuxResponse) Callback(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	result, status := m.Handle(&Parameters{
		Body: body,
	})
	if status != http.StatusOK {
		w.WriteHeader(status)
	}

	json.NewEncoder(w).Encode(result)
}
