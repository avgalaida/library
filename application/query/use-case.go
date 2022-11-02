package main

import (
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_store"
)

func GetBookList() []domain.BookView {
	var books []domain.BookView

	aemap := event_store.GetAll()
	for aggregate, events := range aemap {
		book := domain.Book{}
		book.Base = aggregate

		for _, event := range events {
			event.ApplyOn(&book)
		}

		bookView := domain.Materialize(book)

		books = append(books, bookView)
	}

	return books
}

func GetBookWithVersion(id, version string) domain.BookView {
	aggregate, events := event_store.GetAggregateVersion(id, version)

	book := domain.Book{}
	book.Base = aggregate
	book.Base.Meta = 0

	for _, event := range events {
		event.ApplyOn(&book)
		book.Base.Meta++
	}

	return domain.Materialize(book)
}
