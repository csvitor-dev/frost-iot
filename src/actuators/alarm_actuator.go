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

func (a *AlarmActuator) HandleMessage(message owtp.Schema[owtp.BodyMessage]) {
	fmt.Println("[AlarmActuator] retrieved message")
}
