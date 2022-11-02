package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type BasedEvent struct {
	ID          string
	AggregateID string
	CreatedAt   string
	UserID      string
	Revision    int
	Data        []byte
	Type        string
}

func (e *BasedEvent) ApplyOn(b *Book) {
	var d IDelta
	switch e.Type {
	case "книга.создана":
		d = &CreateBookDelta{}
	case "книга.удалена":
		d = &DeleteBookDelta{}
	case "книга.восстановлена":
		d = &RestoreBookDelta{}
	case "название.изменено":
		d = &ChangeBookTitleDelta{}
	case "авторство.изменено":
		d = &ChangeBookAuthorsDelta{}
	case "откат.версии":
		d = &RollbackBookDelta{}
	}
	json.Unmarshal(e.Data, &d)
	d.ApplyOn(b)
}

func NewEvent(a BasedAggregate, delta IDelta, userID string) BasedEvent {
	jsonDelta, _ := json.Marshal(delta)
	eventType := delta.Key()

	return BasedEvent{
		ID:          uuid.New().String(),
		AggregateID: a.ID,
		CreatedAt:   time.Now().UTC().String(),
		UserID:      userID,
		Revision:    a.Meta + 1,
		Data:        jsonDelta,
		Type:        eventType,
	}
}
