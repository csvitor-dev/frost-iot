package owtp

type Sensor interface {
	catchEvent() *Schema
	SendMessages()
}
