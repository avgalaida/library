package main

import (
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_pubsub"
	"github.com/avgalaida/library/infrastructure/event_sourcing"
	"github.com/avgalaida/library/infrastructure/event_store"
	"github.com/avgalaida/library/infrastructure/utilits"
	"github.com/google/uuid"
	"html/template"
	"log"
	"net/http"
	"time"
)

type response struct {
	ID string `json:"id"`
}

func reposit(aggregateBase event_sourcing.BasedAggregate, event event_sourcing.BasedEvent) {
	event_store.UpdateAggregateRevision(aggregateBase.ID)
	event_store.InsertEvent(event)
	event_pubsub.Publish(event)
}

func createBookCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("createHandler")
	title := template.HTMLEscapeString(r.FormValue("title"))
	authors := template.HTMLEscapeString(r.FormValue("authors"))

	bookBase := event_sourcing.BasedAggregate{
		ID:        uuid.New().String(),
		Meta:      0,
		CreatedAt: time.Now().UTC().String(),
	}

	delta := domain.CreateBookDelta{
		ID:        bookBase.ID,
		Meta:      bookBase.Meta + 1,
		Status:    "Доступна",
		Title:     title,
		Authors:   authors,
		CreatedAt: bookBase.CreatedAt,
	}

	event := event_sourcing.NewEvent(bookBase, delta, r.RemoteAddr)
	event_store.InsertAggregate(bookBase)

	reposit(bookBase, event)

	util.ResponseOk(w, response{ID: bookBase.ID})
}

func deleteBookCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteHandler")
	id := template.HTMLEscapeString(r.FormValue("id"))

	bookBase := event_store.GetAggregate(id)

	delta := domain.DeleteBookDelta{
		ID:   bookBase.ID,
		Meta: bookBase.Meta + 1,
	}

	event := event_sourcing.NewEvent(bookBase, delta, r.RemoteAddr)

	reposit(bookBase, event)

	util.ResponseOk(w, response{ID: bookBase.ID})
}

func restoreBookCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("restoreHandler")
	id := template.HTMLEscapeString(r.FormValue("id"))
	status := template.HTMLEscapeString(r.FormValue("status"))

	bookBase := event_store.GetAggregate(id)

	delta := domain.RestoreBookDelta{
		ID:     bookBase.ID,
		Meta:   bookBase.Meta + 1,
		Status: status,
	}

	event := event_sourcing.NewEvent(bookBase, delta, r.RemoteAddr)

	reposit(bookBase, event)

	util.ResponseOk(w, response{ID: bookBase.ID})
}
