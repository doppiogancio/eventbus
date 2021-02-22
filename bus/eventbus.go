package bus

import (
	"fmt"
	"github.com/doppiogancio/eventbus/event"
	"github.com/doppiogancio/eventbus/subscriber"
)

type (
	eventBus struct {
		debug       bool
		subscribers map[string]map[string]subscriber.Subscriber
	}
)

func EventBus(debug bool) *eventBus {
	return &eventBus{
		debug:       debug,
		subscribers: map[string]map[string]subscriber.Subscriber{},
	}
}

func (b *eventBus) Publish(e event.Event) {
	arg := e.Fire()
	if b.debug {
		fmt.Println("Published event", e.Name())
	}

	for _, s := range b.subscribers[e.Name()] {
		if b.debug {
			fmt.Println("Event received by", s.Name())
		}
		s.Receive(arg)
	}
}

func (b *eventBus) Subscribe(event string, s subscriber.Subscriber) {
	_, ok := b.subscribers[event]
	if !ok {
		b.subscribers[event] = map[string]subscriber.Subscriber{}
	}
	b.subscribers[event][s.Name()] = s
}

func (b *eventBus) Unsubscribe(event string, s subscriber.Subscriber) {
	b.subscribers[event][s.Name()] = nil
}
