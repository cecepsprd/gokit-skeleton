package endpoint

import (
	"context"

	"github.com/cecepsprd/gokit-skeleton/internal/model"
	"github.com/cecepsprd/gokit-skeleton/internal/service"
	"github.com/go-kit/kit/endpoint"
)

type PersonEndpoint struct {
	GetHandler    endpoint.Endpoint
	GetAllHandler endpoint.Endpoint
}

func MakePersonEndpoint(s service.PersonService) PersonEndpoint {
	return PersonEndpoint{
		GetHandler:    getPersonEndpoint(s),
		GetAllHandler: getPersonsEndpoint(s),
	}
}

func getPersonEndpoint(s service.PersonService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetPersonRequest)
		person, err := s.GetPerson(ctx, req.ID)
		return person, err
	}
}

func getPersonsEndpoint(s service.PersonService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetPersons(context.TODO())
	}
}
