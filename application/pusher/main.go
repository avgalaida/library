package main

import (
	"fmt"
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_pubsub"
	"github.com/avgalaida/library/infrastructure/utilits"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"time"
)

type Config struct {
	NatsAddress string `envconfig:"NATS_ADDRESS"`
}

func main() {
	var cfg Config
	envconfig.Process("", &cfg)

	hub := newHub()
	util.ForeverSleep(2*time.Second, func(_ int) error {
		ps, err := event_pubsub.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
		if err != nil {
			log.Println(err)
			return err
		}
		event_pubsub.SetPubSub(ps)

		event_pubsub.OnBookCreated(func(m domain.CreateBookDelta) {
			log.Printf("Создана книга: %v\n", m)
			hub.broadcast(domain.CreateBookDelta{
				Type:      "create_book",
				ID:        m.ID,
				Status:    m.Status,
				Title:     m.Title,
				Authors:   m.Authors,
				CreatedAt: m.CreatedAt,
			}, nil)
		})
		event_pubsub.OnBookDeleted(func(m domain.DeleteBookDelta) {
			log.Printf("Удалена книга: %v\n", m)
			hub.broadcast(domain.DeleteBookDelta{
				Type: "delete_book",
				ID:   m.ID,
			}, nil)
		})
		event_pubsub.OnBookRestored(func(m domain.RestoreBookDelta) {
			log.Printf("Восстановлена книга: %v\n", m)
			hub.broadcast(domain.RestoreBookDelta{
				Type:   "restore_book",
				ID:     m.ID,
				Status: m.Status,
			}, nil)
		})
		event_pubsub.OnBookTitleChanged(func(m domain.ChangeBookTitleDelta) {
			log.Printf("Изменено название книги: %v\n", m)
			hub.broadcast(domain.ChangeBookTitleDelta{
				Type:  "change_title",
				ID:    m.ID,
				Title: m.Title,
			}, nil)
		})
		event_pubsub.OnBookAuthorsChanged(func(m domain.ChangeBookAuthorsDelta) {
			log.Printf("Изменено авторство книги: %v\n", m)
			hub.broadcast(domain.ChangeBookAuthorsDelta{
				Type:    "change_authors",
				ID:      m.ID,
				Authors: m.Authors,
			}, nil)
		})

		return nil
	})
	defer event_pubsub.Close()

	go hub.run()
	http.HandleFunc("/pusher", hub.handleWebSocket)
	http.ListenAndServe(":8080", nil)
}
