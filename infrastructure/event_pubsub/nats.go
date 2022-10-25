package event_pubsub

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_sourcing"
	"github.com/nats-io/nats.go"
)

type NatsEventPublisher struct {
	nc                 *nats.Conn
	eventsSubscription *nats.Subscription
}

func NewNats(url string) (*NatsEventPublisher, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventPublisher{nc: nc}, nil
}

func (ep *NatsEventPublisher) Close() {
	ep.nc.Close()
	ep.eventsSubscription.Unsubscribe()
}

func (ep *NatsEventPublisher) writeMessage(m domain.Message) []byte {
	b := bytes.Buffer{}
	gob.NewEncoder(&b).Encode(m)
	return b.Bytes()
}

func (ep *NatsEventPublisher) readMessage(data []byte, m interface{}) {
	b := bytes.Buffer{}
	b.Write(data)
	gob.NewDecoder(&b).Decode(m)
}

func (ep *NatsEventPublisher) Publish(event event_sourcing.BasedEvent) {
	var m domain.Message
	switch event.Type {
	case "CreateBookDelta":
		m = &domain.CreateBookDelta{}
	case "DeleteBookDelta":
		m = &domain.DeleteBookDelta{}
	case "RestoreBookDelta":
		m = &domain.RestoreBookDelta{}
	case "ChangeBookTitleDelta":
		m = &domain.ChangeBookTitleDelta{}
	case "ChangeBookAuthorsDelta":
		m = &domain.ChangeBookAuthorsDelta{}
	}
	json.Unmarshal(event.Data, &m)
	data := ep.writeMessage(m)
	ep.nc.Publish(m.Key(), data)
}

func (ep *NatsEventPublisher) OnBookCreated(f func(domain.CreateBookDelta)) {
	m := domain.CreateBookDelta{}
	ep.eventsSubscription, _ = ep.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ep.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ep *NatsEventPublisher) OnBookDeleted(f func(domain.DeleteBookDelta)) {
	m := domain.DeleteBookDelta{}
	ep.eventsSubscription, _ = ep.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ep.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ep *NatsEventPublisher) OnBookRestored(f func(domain.RestoreBookDelta)) {
	m := domain.RestoreBookDelta{}
	ep.eventsSubscription, _ = ep.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ep.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ep *NatsEventPublisher) OnBookTitleChanged(f func(delta domain.ChangeBookTitleDelta)) {
	m := domain.ChangeBookTitleDelta{}
	ep.eventsSubscription, _ = ep.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ep.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ep *NatsEventPublisher) OnBookAuthorsChanged(f func(delta domain.ChangeBookAuthorsDelta)) {
	m := domain.ChangeBookAuthorsDelta{}
	ep.eventsSubscription, _ = ep.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ep.readMessage(msg.Data, &m)
		f(m)
	})
}
