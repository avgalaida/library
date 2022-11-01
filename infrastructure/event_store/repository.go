package event_store

import (
	"github.com/avgalaida/library/domain"
)

type Repository interface {
	Close()
	InsertAggregate(b domain.BasedAggregate)
	GetAggregate(id string) domain.BasedAggregate
	UpdateAggregateRevision(id string)
	GetAggregateVersion(id string, version string) (domain.BasedAggregate, []domain.BasedEvent)
	InsertEvent(e domain.BasedEvent)
	GetAll() map[domain.BasedAggregate][]domain.BasedEvent
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertAggregate(a domain.BasedAggregate) {
	impl.InsertAggregate(a)
}

func GetAggregate(id string) domain.BasedAggregate {
	return impl.GetAggregate(id)
}

func UpdateAggregateRevision(id string) {
	impl.UpdateAggregateRevision(id)
}

func GetAggregateVersion(id string, version string) (domain.BasedAggregate, []domain.BasedEvent) {
	return impl.GetAggregateVersion(id, version)
}

func InsertEvent(e domain.BasedEvent) {
	impl.InsertEvent(e)
}

func GetAll() map[domain.BasedAggregate][]domain.BasedEvent {
	return impl.GetAll()
}
