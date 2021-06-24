package handler

import (
	"encoding/json"

	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gorilla/mux"
)

type Handle func(params *Parameters) (interface{}, int)

type Headers struct {
	Origin  string
	Methods string
	Headers string
}
type MuxResponse struct {
	Handle Handle
}

/*
	In here we are showing the benefits of having a response for each architecture instead of just having different function names (Callback)
	With mux we could also pass the Headers object to define it by route or we could define it overall. For demo purposes I will leave
	mux without implementing headers on the route level, and on AWS (which will be production) will have to Headers struct per route.
*/
type AwsResponse struct {
	Handle  Handle
	Headers Headers
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

func (r AwsResponse) Callback(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{"Access-Control-Allow-Origin": r.Headers.Origin, "Access-Control-Allow-Methods": r.Headers.Methods, "Access-Control-Allow-Headers": r.Headers.Headers}

	result, status := r.Handle(&Parameters{
		PathParams:  req.PathParameters,
		QueryParams: req.QueryStringParameters,
		Body:        []byte(req.Body),
	})

	jsonString, _ := json.Marshal(result)
	return events.APIGatewayProxyResponse{Body: string(jsonString), StatusCode: status, Headers: headers}, nil
}
