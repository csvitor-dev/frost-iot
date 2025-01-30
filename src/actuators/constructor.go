package actuators

import (
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/pkg/types"
)

type ActuatorApplication interface {
	HandleMessage(*owtp.Schema)
}

// returns the reference an actuator base with injectable struct
func NewActuator(inject ActuatorApplication) owtp.Actuator {
	return &ActuatorBase{
		id: types.NewUUID(),
		state: false,
		server: nil,
		children: inject,
	}
}