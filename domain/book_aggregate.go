package domain

import (
	"encoding/json"
)

type Book struct {
	Base    BasedAggregate
	Status  string
	Title   string
	Authors string
}

func (b *Book) ApplyEvent(event BasedEvent) {
	switch event.Type {
	case "книга.создана":
		delta := CreateBookDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Status = delta.Status
		b.Title = delta.Title
		b.Authors = delta.Authors

	case "книга.удалена":
		delta := DeleteBookDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Status = "Недоступна"

	case "книга.восстановлена":
		delta := RestoreBookDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Status = delta.Status

	case "название.изменено":
		delta := ChangeBookTitleDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Title = delta.Title

	case "авторство.изменено":
		delta := ChangeBookAuthorsDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Authors = delta.Authors

	case "откат.версии":
		delta := RollbackBookDelta{}
		json.Unmarshal(event.Data, &delta)
		b.Title = delta.Title
		b.Authors = delta.Authors
		b.Status = delta.Status
	}

	b.Base.Meta += 1
}
