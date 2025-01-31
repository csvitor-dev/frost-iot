package actuators

import (
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/pkg/types"
)

// returns the reference an actuator base with injectable struct
func NewActuator(inject types.ActuatorApplication) owtp.Actuator {
	return &ActuatorBase{
		id: types.NewUUID(),
		state: false,
		server: nil,
		children: inject,
	}
}