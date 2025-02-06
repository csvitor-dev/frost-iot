package actuators

import (
	"fmt"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

type LightActuator struct {
}

/* construtor */

func NewLightActuator() *LightActuator {
	return &LightActuator{}
}

/* methods */

func (a *LightActuator) HandleMessage(message owtp.Schema) {
	fmt.Println("[LightActuator] retrieved message")
}
