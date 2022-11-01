package main

import (
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_pubsub"
	"github.com/avgalaida/library/infrastructure/event_sourcing"
	"github.com/avgalaida/library/infrastructure/event_store"
	"github.com/google/uuid"
	"time"
)

func reposit(aggregateBase event_sourcing.BasedAggregate, event event_sourcing.BasedEvent) {
	event_store.UpdateAggregateRevision(aggregateBase.ID)
	event_store.InsertEvent(event)
	event_pubsub.Publish(event)
}

func CreateBook(title, authors, userId string) string {
	bookBase := event_sourcing.BasedAggregate{
		ID:        uuid.New().String(),
		Meta:      0,
		CreatedAt: time.Now().UTC().String(),
	}

	delta := domain.CreateBookDelta{
		ID:        bookBase.ID,
		Status:    "Доступна",
		Title:     title,
		Authors:   authors,
		CreatedAt: bookBase.CreatedAt,
	}

	event := event_sourcing.NewEvent(bookBase, &delta, userId)
	event_store.InsertAggregate(bookBase)
	reposit(bookBase, event)

	return bookBase.ID
}

func DeleteBook(id, userId string) string {
	bookBase := event_store.GetAggregate(id)

	delta := domain.DeleteBookDelta{ID: bookBase.ID}

	event := event_sourcing.NewEvent(bookBase, &delta, userId)
	reposit(bookBase, event)

	return bookBase.ID
}

func RestoreBook(id, status, userId string) string {
	bookBase := event_store.GetAggregate(id)

	delta := domain.RestoreBookDelta{
		ID:     bookBase.ID,
		Status: status,
	}

	event := event_sourcing.NewEvent(bookBase, &delta, userId)
	reposit(bookBase, event)

	return bookBase.ID
}

func ChangeBookTitle(id, title, userId string) string {
	bookBase := event_store.GetAggregate(id)

	delta := domain.ChangeBookTitleDelta{
		ID:    bookBase.ID,
		Title: title,
	}

	event := event_sourcing.NewEvent(bookBase, &delta, userId)
	reposit(bookBase, event)

	return bookBase.ID
}

func ChangeBookAuthors(id, authors, userId string) string {
	bookBase := event_store.GetAggregate(id)

	delta := domain.ChangeBookAuthorsDelta{
		ID:      bookBase.ID,
		Authors: authors,
	}

	event := event_sourcing.NewEvent(bookBase, &delta, userId)
	reposit(bookBase, event)

	return bookBase.ID
}

func RollBackBook(id, status, title, authors, userId string) string {
	bookBase := event_store.GetAggregate(id)

	delta := domain.RollbackBookDelta{
		ID:      bookBase.ID,
		Status:  status,
		Title:   title,
		Authors: authors,
	}

	event := event_sourcing.NewEvent(bookBase, &delta, userId)
	reposit(bookBase, event)

	return bookBase.ID
}
