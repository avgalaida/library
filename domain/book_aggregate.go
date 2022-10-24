package domain

import (
	"encoding/json"
	"github.com/avgalaida/library/infrastructure/event_sourcing"
	"strconv"
)

type Book struct {
	Base        event_sourcing.BasedAggregate
	Status      string
	Title       string
	Authors     string
	Description string
}

func (b *Book) ApplyEvent(event event_sourcing.BasedEvent) {
	switch event.Type {
	case "CreateBookDelta":
		delta := CreateBookDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Status = delta.Status
		b.Title = delta.Title
		b.Authors = delta.Authors
		b.Description = delta.Description

	case "DeleteBookDelta":
		delta := DeleteBookDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Status = delta.Status
	}
	m, _ := strconv.Atoi(b.Base.Meta)
	b.Base.Meta = strconv.Itoa(m + 1)
}
