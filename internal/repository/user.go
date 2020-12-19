package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/cecepsprd/gokit-skeleton/commons/channels"
	"github.com/cecepsprd/gokit-skeleton/internal/model"
)

type PersonRepository interface {
	GetByID(ctx context.Context, id string) (model.Person, error)
	Get(ctx context.Context) ([]model.Person, error)
}

type personRepository struct {
	DB *sql.DB
}

func NewPersonRepository(db *sql.DB) PersonRepository {
	return &personRepository{db}
}

func (pr *personRepository) GetByID(ctx context.Context, id string) (model.Person, error) {
	var result model.Person
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		qry := "SELECT id, name, email FROM person WHERE id = ?"
		err = pr.DB.QueryRow(qry, id).Scan(
			&result.ID, &result.Name, &result.Email,
		)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return result, nil
	}

	if err == sql.ErrNoRows {
		return result, errors.New("person not found")
	}

	fmt.Println(result)

	return result, nil
}

func (pr *personRepository) Get(ctx context.Context) ([]model.Person, error) {
	// var result []model.Person

	qry := "SELECT id, name, email FROM person"
	rows, err := pr.DB.QueryContext(ctx, qry)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	var result = make([]model.Person, 0)
	for rows.Next() {
		var person model.Person
		err = rows.Scan(&person.ID, &person.Name, &person.Email)
		if err != nil {
			return nil, err
		}
		result = append(result, person)
	}

	return result, nil
}
