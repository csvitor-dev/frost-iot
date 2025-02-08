package sensors

import (
	"fmt"

	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

type PortStateSensor struct {
	portStateReceived bool
}

/* getters */

func (s *PortStateSensor) GetPortState() bool {
	return s.portStateReceived
}

/* methods */

func (t *PortStateSensor) CatchEvent() owtp.Schema[req.SensorRequest] {
	fmt.Println(" >> Event captured by PortState Sensor")
	return owtp.Schema[req.SensorRequest]{}
}
