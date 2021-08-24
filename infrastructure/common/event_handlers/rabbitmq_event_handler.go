package event_handlers

import (
	"fmt"
	"go-ddd-example/domain/common"
)

type RabbitMQEventHandler struct {
}

func NewRabbitMQEventHandler() common.IEventHandler {
	return &RabbitMQEventHandler{}
}

func (handler RabbitMQEventHandler) Handle(event common.IBaseEvent) {
	fmt.Println(event)
}
