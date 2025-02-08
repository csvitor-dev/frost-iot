package sensors

import (
	"fmt"

	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

type TemperatureSensor struct {
	temperatureReceived float64
}

/* getters */

func (s *TemperatureSensor) GetTemperature() float64 {
	return s.temperatureReceived
}

/* methods */

func (t *TemperatureSensor) CatchEvent() owtp.Schema[req.TemperatureRequest] {
	fmt.Println(" >> Event captured by Temperature Sensor")
	return owtp.Schema[req.TemperatureRequest]{}
}
