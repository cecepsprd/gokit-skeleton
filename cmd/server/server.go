package server

import (
	"log"
	"time"

	rds "github.com/cecepsprd/gokit-skeleton/commons/cache"
	"github.com/cecepsprd/gokit-skeleton/commons/config"
	"github.com/cecepsprd/gokit-skeleton/commons/db"
	"github.com/cecepsprd/gokit-skeleton/internal/repository"
	"github.com/cecepsprd/gokit-skeleton/internal/service"
	grpc "github.com/cecepsprd/gokit-skeleton/transport/grpc"
	rest "github.com/cecepsprd/gokit-skeleton/transport/http"
)

func RunServer() error {
	cfg := config.LoadConfiguration()

	// connect to database
	db, err := db.MysqlConnect(cfg)
	if err != nil {
		log.Fatal("error connecting to database: ", err.Error())
	}

	// repository
	personRepo := repository.NewPersonRepository(db)

	// service
	personService := service.NewPersonService(personRepo)

	personCache := rds.NewPersonCache(cfg, time.Hour*1)

	go func() {
		_ = rest.RunServer(personService, personCache, cfg)
	}()

	return grpc.RunServer(personService, personCache, cfg)
}
