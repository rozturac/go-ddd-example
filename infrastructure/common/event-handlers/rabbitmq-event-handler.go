package event_handlers

import (
	"github.com/rozturac/rmqc"
	"go-ddd-example/api/consts"
	"go-ddd-example/domain/common"
	"reflect"
)

type RabbitMQEventHandler struct {
	rbt     *rmqc.RabbitMQ
	appName string
}

func NewRabbitMQEventHandler(rbt *rmqc.RabbitMQ) common.IEventHandler {
	return &RabbitMQEventHandler{rbt: rbt}
}

func (handler RabbitMQEventHandler) Handle(event common.IBaseEvent) {
	t := reflect.TypeOf(event)
	eventName := t.Elem().Name()
	handler.rbt.Publish(consts.AppName, eventName, event)
}
