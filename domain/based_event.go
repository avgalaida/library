package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
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
	Delta       IDelta
}

func (e *BasedEvent) ApplyOn(b *Book) {
	e.Delta.ApplyOn(b)
}

func NewEvent(a BasedAggregate, delta IDelta, userID string) BasedEvent {
	jsonDelta, _ := json.Marshal(delta)
	eventType := reflect.TypeOf(delta).Name()

	return BasedEvent{
		ID:          uuid.New().String(),
		AggregateID: a.ID,
		CreatedAt:   time.Now().UTC().String(),
		UserID:      userID,
		Revision:    a.Meta + 1,
		Data:        jsonDelta,
		Type:        eventType,
		Delta:       delta,
	}
}
