package endpoint

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/cecepsprd/gokit-skeleton/api/proto/v1"
	"github.com/cecepsprd/gokit-skeleton/commons/cache"
	"github.com/cecepsprd/gokit-skeleton/internal/model"
	"github.com/cecepsprd/gokit-skeleton/internal/service"
	"github.com/go-kit/kit/endpoint"
)

type PersonEndpoint struct {
	GetHandler    endpoint.Endpoint
	GetAllHandler endpoint.Endpoint
}

func MakePersonEndpoint(s service.PersonService, pc cache.PersonCache) PersonEndpoint {
	return PersonEndpoint{
		GetHandler:    getPersonEndpoint(s),
		GetAllHandler: getPersonsEndpoint(s, pc),
	}
}

func getPersonEndpoint(s service.PersonService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetPersonRequest)
		person, err := s.GetPerson(ctx, req.ID)
		return person, err
	}
}

func getPersonsEndpoint(s service.PersonService, cache cache.PersonCache) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(v1.ReadAllPersonRequest)

		fmt.Println("===============================")
		fmt.Println(req)
		fmt.Println("===============================")

		// Get cached persons
		personsCache := cache.Get("persons")

		fmt.Println(personsCache)

		// if not exists
		if personsCache != nil {
			fmt.Println("CACHED PERSONS NIL .......")
			persons, err := s.GetPersons(context.TODO())

			// cache persons
			cache.SetPersons("persons", persons)
			return persons, err
		}

		log.Println("get cached persons")
		return personsCache, nil
	}
}
