package main

import (
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_store"
	"github.com/avgalaida/library/infrastructure/utilits"
	"html/template"
	"net/http"
)

func bookListQueryHandler(w http.ResponseWriter, _ *http.Request) {
	var books []domain.BookView

	aemap := event_store.GetAll()
	for aggregate, events := range aemap {
		book := domain.Book{}
		book.Base = aggregate
		book.Base.Meta = 0

		for _, event := range events {
			book.ApplyEvent(event)
		}

		bookView := domain.NewBookView(book)

		books = append(books, bookView)
	}
	util.ResponseOk(w, books)
}

func bookVersionQueryHandler(w http.ResponseWriter, r *http.Request) {
	id := template.HTMLEscapeString(r.FormValue("id"))
	version := template.HTMLEscapeString(r.FormValue("version"))
	aggregate, events := event_store.GetAggregateVersion(id, version)

	book := domain.Book{}
	book.Base = aggregate
	book.Base.Meta = 0

	for _, event := range events {
		book.ApplyEvent(event)
	}

	bookView := domain.NewBookView(book)

	util.ResponseOk(w, bookView)
}
