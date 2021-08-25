package common_di

import (
	"github.com/rozturac/rmqc"
	"go-ddd-example/api/configs"
	"go-ddd-example/domain/common"
	"go-ddd-example/infrastructure/common/event-handlers"
	"sync"
)

var (
	once = sync.Once{}
	rbt  *rmqc.RabbitMQ
)

func NewEventHandlerResolve(rbt *rmqc.RabbitMQ) common.IEventHandler {
	return event_handlers.NewRabbitMQEventHandler(rbt)
}

func NewRabbitMQResolve(config configs.Config) *rmqc.RabbitMQ {
	var err error

	once.Do(func() {
		rbt, err = rmqc.Connect(rmqc.RabbitMQConfig{
			Host:           config.RabbitMQ.Host,
			Username:       config.RabbitMQ.Username,
			Password:       config.RabbitMQ.Password,
			Port:           config.RabbitMQ.Port,
			VHost:          config.RabbitMQ.VHost,
			ConnectionName: config.RabbitMQ.ConnectionName,
			Reconnect: rmqc.Reconnect{
				MaxAttempt: config.RabbitMQ.Reconnect.MaxAttempt,
				Interval:   config.RabbitMQ.Reconnect.Interval,
			},
		})
	})

	if err != nil {
		panic(err)
	}

	return rbt
}
