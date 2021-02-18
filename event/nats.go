package event

import (
	"bytes"
	"encoding/gob"
	"github.com/bramalho/go-cqrs/model"
	"log"

	"github.com/nats-io/nats.go"
)

type NatsEventStore struct {
	nc                      *nats.Conn
	todoCreatedSubscription *nats.Subscription
	todoCreatedChan         chan TodoCreatedMessage
}

func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

func (es *NatsEventStore) SubscribeTodoCreated() (<-chan TodoCreatedMessage, error) {
	m := TodoCreatedMessage{}
	es.todoCreatedChan = make(chan TodoCreatedMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	es.todoCreatedSubscription, err = es.nc.ChanSubscribe(m.Key(), ch)
	if err != nil {
		return nil, err
	}
	// Decode message
	go func() {
		for {
			select {
			case msg := <-ch:
				if err := es.readMessage(msg.Data, &m); err != nil {
					log.Fatal(err)
				}
				es.todoCreatedChan <- m
			}
		}
	}()
	return (<-chan TodoCreatedMessage)(es.todoCreatedChan), nil
}

func (es *NatsEventStore) OnTodoCreated(f func(TodoCreatedMessage)) (err error) {
	m := TodoCreatedMessage{}
	es.todoCreatedSubscription, err = es.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		if err := es.readMessage(msg.Data, &m); err != nil {
			log.Fatal(err)
		}
		f(m)
	})
	return
}

func (es *NatsEventStore) Close() {
	if es.nc != nil {
		es.nc.Close()
	}
	if es.todoCreatedSubscription != nil {
		if err := es.todoCreatedSubscription.Unsubscribe(); err != nil {
			log.Fatal(err)
		}
	}
	close(es.todoCreatedChan)
}

func (es *NatsEventStore) PublishTodoCreated(todo model.Todo) error {
	m := TodoCreatedMessage{todo.ID, todo.Body, todo.CreatedAt}
	data, err := es.writeMessage(&m)
	if err != nil {
		return err
	}
	return es.nc.Publish(m.Key(), data)
}

func (es *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (es *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
