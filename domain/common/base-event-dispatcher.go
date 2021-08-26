package common

type IEventDispatcher interface {
	Dispatch(events []IBaseEvent)
}
