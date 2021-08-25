package consumers

import (
	"fmt"
	"github.com/rozturac/rmqc"
	"go-ddd-example/domain/users/events"
)

type UserCreatedConsumer struct {
}

func NewUserCreatedConsumer() *UserCreatedConsumer {
	return &UserCreatedConsumer{}
}

func (u UserCreatedConsumer) Configure(builder *rmqc.ConsumerBuilder) {
	builder.BindQueue("go-ddd-example-user-UserCreated")
	builder.SubscribeAsTopic("go-ddd-example-user", "UserCreated")
	builder.SetConsumerName("UserCreated")
	builder.SetPrefetchCount(3)
	builder.SetConsumerCount(3)
}

func (u UserCreatedConsumer) Consume(context *rmqc.ConsumerContext) {
	var event events.UserCreated
	context.Unmarshal(&event)
	fmt.Println(event)
}
