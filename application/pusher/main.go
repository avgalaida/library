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
			hub.broadcast(&domain.CreateBookDelta{
				ID:        m.ID,
				Status:    m.Status,
				Title:     m.Title,
				Authors:   m.Authors,
				CreatedAt: m.CreatedAt,
			})
		})
		event_pubsub.OnBookDeleted(func(m domain.DeleteBookDelta) {
			log.Printf("Удалена книга: %v\n", m)
			hub.broadcast(&domain.DeleteBookDelta{
				ID: m.ID,
			})
		})
		event_pubsub.OnBookRestored(func(m domain.RestoreBookDelta) {
			log.Printf("Восстановлена книга: %v\n", m)
			hub.broadcast(&domain.RestoreBookDelta{
				ID:     m.ID,
				Status: m.Status,
			})
		})
		event_pubsub.OnBookTitleChanged(func(m domain.ChangeBookTitleDelta) {
			log.Printf("Изменено название книги: %v\n", m)
			hub.broadcast(&domain.ChangeBookTitleDelta{
				ID:    m.ID,
				Title: m.Title,
			})
		})
		event_pubsub.OnBookAuthorsChanged(func(m domain.ChangeBookAuthorsDelta) {
			log.Printf("Изменено авторство книги: %v\n", m)
			hub.broadcast(&domain.ChangeBookAuthorsDelta{
				ID:      m.ID,
				Authors: m.Authors,
			})
		})
		event_pubsub.OnBookRollbacked(func(m domain.RollbackBookDelta) {
			log.Printf("Откат версии книги: %v\n", m)
			hub.broadcast(&domain.RollbackBookDelta{
				ID:      m.ID,
				Status:  m.Status,
				Title:   m.Title,
				Authors: m.Authors,
			})
		})

		return nil
	})
	defer event_pubsub.Close()

	go hub.run()
	http.HandleFunc("/pusher", hub.handleWebSocket)
	http.ListenAndServe(":8080", nil)
}
