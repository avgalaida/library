package main

import (
	"fmt"
	"github.com/avgalaida/library/infrastructure/event_pubsub"
	"github.com/avgalaida/library/infrastructure/event_store"
	"github.com/avgalaida/library/infrastructure/utilits"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"time"
)

type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress      string `envconfig:"NATS_ADDRESS"`
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/books", createBookCommandHandler).
		Methods(http.MethodPost).
		Queries("title", "{title}", "authors", "{authors}")
	router.HandleFunc("/books", restoreBookCommandHandler).
		Methods(http.MethodPost).
		Queries("id", "{id}", "status", "{status}")
	router.HandleFunc("/books", changeBookTitleCommandHandler).
		Methods(http.MethodPost).
		Queries("id", "{id}", "title", "{title}")
	router.HandleFunc("/books", changeBookAuthorsCommandHandler).
		Methods(http.MethodPost).
		Queries("id", "{id}", "authors", "{authors}")
	router.HandleFunc("/books", deleteBookCommandHandler).
		Methods(http.MethodPost).
		Queries("id", "{id}")
	router.Use(mux.CORSMethodMiddleware(router))
	return
}

func main() {
	var cfg Config
	envconfig.Process("", &cfg)

	util.ForeverSleep(2*time.Second, func(attempt int) error {
		addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
		repo, err := event_store.NewPostgres(addr)
		if err != nil {
			log.Println(err)
			return err
		}
		event_store.SetRepository(repo)
		return nil
	})
	defer event_store.Close()

	util.ForeverSleep(2*time.Second, func(_ int) error {
		ps, err := event_pubsub.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
		if err != nil {
			log.Println(err)
			return err
		}
		event_pubsub.SetPubSub(ps)
		return nil
	})
	defer event_pubsub.Close()

	router := newRouter()
	http.ListenAndServe(":8080", router)
}
