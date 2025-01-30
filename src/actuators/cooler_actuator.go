package actuators

import (
	"fmt"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

type CoolerActuator struct {
	temperatureTarget float64
}

/* construtor */

func NewCoolerActuator(target float64) *CoolerActuator {
	return &CoolerActuator{
		temperatureTarget: target,	
	}
}

/* methods */

func (a *CoolerActuator) HandleMessage(message *owtp.Schema) {
	fmt.Println("[CoolerActuator] retrieved message")
}
