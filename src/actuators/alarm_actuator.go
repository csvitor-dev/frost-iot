package actuators

import (
	"fmt"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

type AlarmActuator struct {
}

/* construtor */

func NewAlarmActuator() *AlarmActuator {
	return &AlarmActuator{}
}

/* methods */

func (a *AlarmActuator) HandleMessage(message *owtp.Schema) {
	fmt.Println("[AlarmActuator] retrieved message")
}
