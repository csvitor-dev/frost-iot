package sensors

import (
	"fmt"

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

func (t *TemperatureSensor) CatchEvent() owtp.Schema {
	fmt.Println(" >> Event captured by Temperature Sensor")
	return owtp.Schema{}
}
