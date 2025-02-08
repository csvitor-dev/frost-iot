package actuators

import (
	"github.com/csvitor-dev/frost-iot/internal/types"
	pkg "github.com/csvitor-dev/frost-iot/pkg/types"
)

// returns the reference an actuator base with injectable struct
func NewActuator(inject pkg.ActuatorApplication) pkg.Actuator {
	return &ActuatorBase{
		id:       types.NewUUID(),
		state:    false,
		server:   nil,
		children: inject,
	}
}
