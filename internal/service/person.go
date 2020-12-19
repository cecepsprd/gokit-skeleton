package service

import (
	"context"

	"github.com/cecepsprd/gokit-skeleton/internal/model"
	"github.com/cecepsprd/gokit-skeleton/internal/repository"
)

type PersonService interface {
	GetPerson(ctx context.Context, id string) (model.Person, error)
	GetPersons(ctx context.Context) ([]model.Person, error)
}

type personSvc struct {
	repo repository.PersonRepository
}

func NewPersonService(repo repository.PersonRepository) PersonService {
	return &personSvc{repo}
}

func (s *personSvc) GetPerson(ctx context.Context, id string) (model.Person, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *personSvc) GetPersons(ctx context.Context) ([]model.Person, error) {
	return s.repo.Get(ctx)
}
