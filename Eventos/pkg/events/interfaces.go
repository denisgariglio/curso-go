package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventHandlerInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventDispatcherInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventDispatcherInterface) bool
	Clear() error
}