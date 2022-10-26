package main

import (
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_store"
	"github.com/avgalaida/library/infrastructure/search"
	"github.com/avgalaida/library/infrastructure/utilits"
	"html/template"
	"net/http"
	"strconv"
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

func searchBooksHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	query := r.FormValue("query")
	if len(query) == 0 {
		util.ResponseError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}
	skip := uint64(0)
	skipStr := r.FormValue("skip")
	take := uint64(100)
	takeStr := r.FormValue("take")
	if len(skipStr) != 0 {
		skip, err = strconv.ParseUint(skipStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid skip parameter")
			return
		}
	}
	if len(takeStr) != 0 {
		take, err = strconv.ParseUint(takeStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid take parameter")
			return
		}
	}

	books := search.SearchBooks(query, skip, take)

	util.ResponseOk(w, books)
}
