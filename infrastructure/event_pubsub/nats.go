package event_pubsub

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/avgalaida/library/domain"
	"github.com/avgalaida/library/infrastructure/event_sourcing"
	"github.com/nats-io/nats.go"
)

type NatsPubSub struct {
	nc                 *nats.Conn
	eventsSubscription *nats.Subscription
}

func NewNats(url string) (*NatsPubSub, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsPubSub{nc: nc}, nil
}

func (ps *NatsPubSub) Close() {
	ps.nc.Close()
	ps.eventsSubscription.Unsubscribe()
}

func (ps *NatsPubSub) writeMessage(m domain.Message) []byte {
	b := bytes.Buffer{}
	gob.NewEncoder(&b).Encode(m)
	return b.Bytes()
}

func (ps *NatsPubSub) readMessage(data []byte, m interface{}) {
	b := bytes.Buffer{}
	b.Write(data)
	gob.NewDecoder(&b).Decode(m)
}

func (ps *NatsPubSub) Publish(event event_sourcing.BasedEvent) {
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
	case "RollbackBookDelta":
		m = &domain.RollbackBookDelta{}
	}
	json.Unmarshal(event.Data, &m)
	data := ps.writeMessage(m)
	ps.nc.Publish(m.Key(), data)
}

func (ps *NatsPubSub) OnBookCreated(f func(domain.CreateBookDelta)) {
	m := domain.CreateBookDelta{}
	ps.eventsSubscription, _ = ps.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ps.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ps *NatsPubSub) OnBookDeleted(f func(domain.DeleteBookDelta)) {
	m := domain.DeleteBookDelta{}
	ps.eventsSubscription, _ = ps.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ps.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ps *NatsPubSub) OnBookRestored(f func(domain.RestoreBookDelta)) {
	m := domain.RestoreBookDelta{}
	ps.eventsSubscription, _ = ps.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ps.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ps *NatsPubSub) OnBookTitleChanged(f func(delta domain.ChangeBookTitleDelta)) {
	m := domain.ChangeBookTitleDelta{}
	ps.eventsSubscription, _ = ps.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ps.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ps *NatsPubSub) OnBookAuthorsChanged(f func(delta domain.ChangeBookAuthorsDelta)) {
	m := domain.ChangeBookAuthorsDelta{}
	ps.eventsSubscription, _ = ps.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ps.readMessage(msg.Data, &m)
		f(m)
	})
}

func (ps *NatsPubSub) OnBookRollbacked(f func(domain.RollbackBookDelta)) {
	m := domain.RollbackBookDelta{}
	ps.eventsSubscription, _ = ps.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		ps.readMessage(msg.Data, &m)
		f(m)
	})
}
