package common

type IEventHandler interface {
	Handle(event IBaseEvent)
}
