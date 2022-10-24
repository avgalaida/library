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
				Type:        "create_book",
				ID:          m.ID,
				Meta:        m.Meta,
				Status:      m.Status,
				Title:       m.Title,
				Authors:     m.Authors,
				Description: m.Description,
				CreatedAt:   m.CreatedAt,
			}, nil)
		})
		event_pubsub.OnBookDeleted(func(m domain.DeleteBookDelta) {
			log.Printf("Удалена книга: %v\n", m)
			hub.broadcast(domain.DeleteBookDelta{
				Type:   "delete_book",
				ID:     m.ID,
				Meta:   m.Meta,
				Status: m.Status,
			}, nil)
		})

		return nil
	})
	defer event_pubsub.Close()

	go hub.run()
	http.HandleFunc("/pusher", hub.handleWebSocket)
	http.ListenAndServe(":8080", nil)
}
