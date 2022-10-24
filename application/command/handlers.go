package main

import (
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_pubsub"
	"github.com/avgalaida/library/infrastructure/event_sourcing"
	"github.com/avgalaida/library/infrastructure/event_store"
	"github.com/avgalaida/library/infrastructure/utilits"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type response struct {
	ID string `json:"id"`
}

func createBookCommandHandler(w http.ResponseWriter, r *http.Request) {
	title := template.HTMLEscapeString(r.FormValue("title"))
	authors := template.HTMLEscapeString(r.FormValue("authors"))
	description := template.HTMLEscapeString(r.FormValue("desc"))

	bookBase := event_sourcing.BasedAggregate{
		ID:        uuid.New().String(),
		Meta:      "0",
		CreatedAt: time.Now().UTC().String(),
	}
	revision, _ := strconv.Atoi(bookBase.Meta)

	delta := domain.CreateBookDelta{
		ID:          bookBase.ID,
		Meta:        strconv.Itoa(revision + 1),
		Status:      "Доступна",
		Title:       title,
		Authors:     authors,
		Description: description,
		CreatedAt:   bookBase.CreatedAt,
	}

	event := event_sourcing.NewEvent(bookBase, delta, r.RemoteAddr)
	event_store.InsertAggregate(bookBase)
	event_store.UpdateAggregateRevision(bookBase.ID)
	event_store.InsertEvent(event)
	event_pubsub.Publish(event)

	util.ResponseOk(w, response{ID: bookBase.ID})
}

func deleteBookCommandHandler(w http.ResponseWriter, r *http.Request) {
	id := template.HTMLEscapeString(r.FormValue("id"))

	bookBase := event_store.GetAggregate(id)
	revision, _ := strconv.Atoi(bookBase.Meta)

	delta := domain.DeleteBookDelta{
		ID:     bookBase.ID,
		Meta:   strconv.Itoa(revision + 1),
		Status: "Недоступна",
	}

	event := event_sourcing.NewEvent(bookBase, delta, r.RemoteAddr)
	event_store.UpdateAggregateRevision(bookBase.ID)
	event_store.InsertEvent(event)
	event_pubsub.Publish(event)

	util.ResponseOk(w, response{ID: bookBase.ID})
}
