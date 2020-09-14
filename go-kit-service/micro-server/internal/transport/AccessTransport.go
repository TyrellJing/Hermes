package transport

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"micro-server/internal/endpoint"
	"net/http"

	"github.com/tidwall/gjson"
)

func DecodeAccessRequest (ctx context.Context, r *http.Request) (interface{}, error) {
	body, _ := ioutil.ReadAll(r.Body)
	result := gjson.Parse(string(body))
	if result.IsObject() {
		username := result.Get("username")
		password := result.Get("password")
		return endpoint.AccessRequest{
			Username: username.String(),
			Password: password.String(),
			Method: r.Method,
		}, nil
	}
	return nil, errors.New("参数错误！")
}

func EncodeAccessResponse (ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Context-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}