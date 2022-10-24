package event_pubsub

import (
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_sourcing"
)

type Pubsub interface {
	Close()
	Publish(event event_sourcing.BasedEvent)
	OnBookCreated(f func(domain.CreateBookDelta))
	OnBookDeleted(f func(domain.DeleteBookDelta))
}

var impl Pubsub

func SetPubSub(ep Pubsub) {
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
