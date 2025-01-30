package owtp

type Sensor interface {
	CatchEvent() *Schema
	SendMessages()
}
