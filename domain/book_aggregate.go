package domain

import (
	"encoding/json"
	"github.com/avgalaida/library/infrastructure/event_sourcing"
)

type Book struct {
	Base    event_sourcing.BasedAggregate
	Status  string
	Title   string
	Authors string
}

func (b *Book) ApplyEvent(event event_sourcing.BasedEvent) {
	switch event.Type {
	case "CreateBookDelta":
		delta := CreateBookDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Status = delta.Status
		b.Title = delta.Title
		b.Authors = delta.Authors

	case "DeleteBookDelta":
		delta := DeleteBookDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Status = delta.Status

	case "RestoreBookDelta":
		delta := RestoreBookDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Status = delta.Status

	}

	b.Base.Meta += 1
}
