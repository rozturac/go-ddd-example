package event_dispatcher

import (
	"github.com/rozturac/rmqc"
	"go-ddd-example/application/consts"
	"go-ddd-example/domain/common"
	"reflect"
)

type RabbitMQEventDispatcher struct {
	rbt     *rmqc.RabbitMQ
	appName string
}

func NewRabbitMQEventDispatcher(rbt *rmqc.RabbitMQ) common.IEventDispatcher {
	return &RabbitMQEventDispatcher{rbt: rbt}
}

func (handler RabbitMQEventDispatcher) Dispatch(events []common.IBaseEvent) {
	for _, event := range events {
		t := reflect.TypeOf(event)
		eventName := t.Elem().Name()
		handler.rbt.Publish(consts.AppName, eventName, event)
	}
}
