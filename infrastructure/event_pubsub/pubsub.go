package event_pubsub

import (
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_sourcing"
)

type PubSub interface {
	Close()
	Publish(event event_sourcing.BasedEvent)
	OnBookCreated(f func(domain.CreateBookDelta))
	OnBookDeleted(f func(domain.DeleteBookDelta))
	OnBookRestored(f func(domain.RestoreBookDelta))
}

var impl PubSub

func SetPubSub(ep PubSub) {
	impl = ep
}

func Close() {
	impl.Close()
}

func Publish(event event_sourcing.BasedEvent) {
	impl.Publish(event)
}

func OnBookCreated(f func(domain.CreateBookDelta)) {
	impl.OnBookCreated(f)
}

func OnBookDeleted(f func(domain.DeleteBookDelta)) {
	impl.OnBookDeleted(f)
}
func OnBookRestored(f func(domain.RestoreBookDelta)) {
	impl.OnBookRestored(f)
}
