package search

import (
	"context"
	"github.com/avgalaida/library/domain"
)

type Repository interface {
	Close()
	InsertBook(book domain.BookView)
	SearchBooks(query string, skip uint64, take uint64) []domain.BookView
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertBook(ctx context.Context, book domain.BookView) {
	impl.InsertBook(book)
}

func SearchBooks(query string, skip uint64, take uint64) []domain.BookView {
	return impl.SearchBooks(query, skip, take)
}
