package common_di

import (
	"go-ddd-example/domain/common"
	"go-ddd-example/infrastructure/common/event_handlers"
)

func NewEventHandlerResolve() common.IEventHandler {
	return event_handlers.NewRabbitMQEventHandler()
}
