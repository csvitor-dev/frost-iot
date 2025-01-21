package owtp

type Actuator interface {
	HandleEvent(Schema)
}
