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
	nc                             *nats.Conn
	bookCreatedSubscription        *nats.Subscription
	bookDeletedSubscription        *nats.Subscription
	bookRestoredSubscription       *nats.Subscription
	bookTitleChangedSubscription   *nats.Subscription
	bookAuthorsChangedSubscription *nats.Subscription
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
	ep.bookCreatedSubscription.Unsubscribe()
	ep.bookDeletedSubscription.Unsubscribe()
	ep.bookRestoredSubscription.Unsubscribe()
	ep.bookTitleChangedSubscription.Unsubscribe()
	ep.bookAuthorsChangedSubscription.Unsubscribe()
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
	switch event.Type {
	case "CreateBookDelta":
		m := domain.CreateBookDelta{}
		json.Unmarshal(event.Data, &m)
		data := ep.writeMessage(&m)
		ep.nc.Publish(m.Key(), data)

	case "DeleteBookDelta":
		m := domain.DeleteBookDelta{}
		json.Unmarshal(event.Data, &m)
		data := ep.writeMessage(&m)
		ep.nc.Publish(m.Key(), data)

	case "RestoreBookDelta":
		m := domain.RestoreBookDelta{}
		json.Unmarshal(event.Data, &m)
		data := ep.writeMessage(&m)
		ep.nc.Publish(m.Key(), data)

	case "ChangeBookTitleDelta":
		m := domain.ChangeBookTitleDelta{}
		json.Unmarshal(event.Data, &m)
		data := ep.writeMessage(&m)
		ep.nc.Publish(m.Key(), data)

	case "ChangeBookAuthorsDelta":
		m := domain.ChangeBookAuthorsDelta{}
		json.Unmarshal(event.Data, &m)
		data := ep.writeMessage(&m)
		ep.nc.Publish(m.Key(), data)
	}
}

func (ep *NatsEventPublisher) OnBookCreated(f func(domain.CreateBookDelta)) {
	m := domain.CreateBookDelta{}
	ep.bookCreatedSubscription, _ = ep.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ep.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ep *NatsEventPublisher) OnBookDeleted(f func(domain.DeleteBookDelta)) {
	m := domain.DeleteBookDelta{}
	ep.bookDeletedSubscription, _ = ep.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ep.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ep *NatsEventPublisher) OnBookRestored(f func(domain.RestoreBookDelta)) {
	m := domain.RestoreBookDelta{}
	ep.bookRestoredSubscription, _ = ep.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ep.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ep *NatsEventPublisher) OnBookTitleChanged(f func(delta domain.ChangeBookTitleDelta)) {
	m := domain.ChangeBookTitleDelta{}
	ep.bookTitleChangedSubscription, _ = ep.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ep.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ep *NatsEventPublisher) OnBookAuthorsChanged(f func(delta domain.ChangeBookAuthorsDelta)) {
	m := domain.ChangeBookAuthorsDelta{}
	ep.bookAuthorsChangedSubscription, _ = ep.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ep.readMessage(msg.Data, &m)
		f(m)
	})
}
