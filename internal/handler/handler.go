package handler

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/cecepsprd/gokit-skeleton/commons/helpers"
	"github.com/cecepsprd/gokit-skeleton/internal/endpoint"
	"github.com/gorilla/mux"
)

func NewPersonHandler(ctx context.Context, ep endpoint.PersonEndpoint) http.Handler {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/person/{id}").Handler(httptransport.NewServer(
		ep.GetHandler,
		helpers.DecodePersonRequest,
		helpers.EncodeResponse,
	))

	r.Methods(http.MethodGet).Path("/person").Handler(httptransport.NewServer(
		ep.GetAllHandler,
		helpers.DecodePersonsRequest,
		helpers.EncodeResponse,
	))

	return r
}
