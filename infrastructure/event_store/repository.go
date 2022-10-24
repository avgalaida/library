package event_store

import (
	"github.com/avgalaida/library/infrastructure/event_sourcing"
)

type Repository interface {
	Close()
	InsertAggregate(b event_sourcing.BasedAggregate)
	GetAggregate(id string) event_sourcing.BasedAggregate
	UpdateAggregateRevision(id string)
	InsertEvent(e event_sourcing.BasedEvent)
	GetAll() map[event_sourcing.BasedAggregate][]event_sourcing.BasedEvent
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertAggregate(a event_sourcing.BasedAggregate) {
	impl.InsertAggregate(a)
}

func GetAggregate(id string) event_sourcing.BasedAggregate {
	return impl.GetAggregate(id)
}

func UpdateAggregateRevision(id string) {
	impl.UpdateAggregateRevision(id)
}

func InsertEvent(e event_sourcing.BasedEvent) {
	impl.InsertEvent(e)
}

func GetAll() map[event_sourcing.BasedAggregate][]event_sourcing.BasedEvent {
	return impl.GetAll()
}
