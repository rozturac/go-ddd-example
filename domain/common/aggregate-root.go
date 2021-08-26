package common

type IAggregateRoot interface {
	AddEvent(event IBaseEvent)
	GetDomainEvents() []IBaseEvent
	ClearDomainEvents()
}
