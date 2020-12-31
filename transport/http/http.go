package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cecepsprd/gokit-skeleton/commons/cache"
	"github.com/cecepsprd/gokit-skeleton/commons/config"
	"github.com/cecepsprd/gokit-skeleton/internal/endpoint"
	"github.com/cecepsprd/gokit-skeleton/internal/handler"
	"github.com/cecepsprd/gokit-skeleton/internal/service"
)

func RunServer(personSvc service.PersonService, personCache cache.PersonCache, cfg config.Config) error {

	var (
		personEndpoint = endpoint.MakePersonEndpoint(personSvc, personCache)
	)

	h := handler.NewPersonHandler(context.TODO(), personEndpoint)

	// create a new server
	s := http.Server{
		Addr:    cfg.App.HttpPort, // configure the bind address
		Handler: h,                // set the default handler
		// ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		log.Println("Starting http server on port: ", cfg.App.HttpPort)

		err := s.ListenAndServe()
		if err != nil {
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	return s.Shutdown(ctx)
}
