package common

type IAggregateRoot interface {
	AddEvent(event IBaseEvent)
	RaiseEvents(handler IEventHandler)
}
