package helpers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cecepsprd/gokit-skeleton/internal/model"
	"github.com/gorilla/mux"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodePersonsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func DecodePersonRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetPersonRequest
	vars := mux.Vars(r)

	req = model.GetPersonRequest{
		ID: vars["id"],
	}
	return req, nil
}
