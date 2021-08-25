package consumers

import (
	"fmt"
	"github.com/rozturac/rmqc"
	"go-ddd-example/api/consts"
	"go-ddd-example/domain/users/events"
)

type UserCreatedConsumer struct {
}

func NewUserCreatedConsumer() *UserCreatedConsumer {
	return &UserCreatedConsumer{}
}

func (u UserCreatedConsumer) Configure(builder *rmqc.ConsumerBuilder) {
	builder.BindQueue(fmt.Sprintf("%s-%s", consts.AppName, "UserCreated"))
	builder.SubscribeAsTopic(consts.AppName, "UserCreated")
	builder.SetConsumerName(fmt.Sprintf("%s-%s", consts.AppName, "UserCreated"))
	builder.SetPrefetchCount(3)
	builder.SetConsumerCount(3)
}

func (u UserCreatedConsumer) Consume(context *rmqc.ConsumerContext) {
	var event events.UserCreated
	context.Unmarshal(&event)
	fmt.Println(event)
}
