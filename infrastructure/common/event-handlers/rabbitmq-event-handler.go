package event_handlers

import (
	"fmt"
	"github.com/rozturac/rmqc"
	"go-ddd-example/domain/common"
	"reflect"
)

type RabbitMQEventHandler struct {
	rbt *rmqc.RabbitMQ
}

func NewRabbitMQEventHandler(rbt *rmqc.RabbitMQ) common.IEventHandler {
	return &RabbitMQEventHandler{rbt: rbt}
}

func (handler RabbitMQEventHandler) Handle(event common.IBaseEvent) {
	t := reflect.TypeOf(event)
	eventName := t.Elem().Name()
	handler.rbt.Publish(fmt.Sprintf("%s-%s", "go-ddd-example", event.GetAggregateName()), eventName, event)
}
