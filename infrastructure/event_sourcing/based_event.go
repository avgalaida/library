package event_sourcing

import (
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
	"strconv"
	"time"
)

type BasedEvent struct {
	ID          string
	AggregateID string
	CreatedAt   string
	UserID      string
	Revision    string
	Data        []byte
	Type        string
}

func NewEvent(a BasedAggregate, delta interface{}, userID string) BasedEvent {
	jsonDelta, _ := json.Marshal(delta)
	eventType := reflect.TypeOf(delta).Name()
	revision, _ := strconv.Atoi(a.Meta)

	return BasedEvent{
		ID:          uuid.New().String(),
		AggregateID: a.ID,
		CreatedAt:   time.Now().UTC().String(),
		UserID:      userID,
		Revision:    strconv.Itoa(revision + 1),
		Data:        jsonDelta,
		Type:        eventType,
	}
}
