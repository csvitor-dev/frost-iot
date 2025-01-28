package sensors

import "github.com/csvitor-dev/frost-iot/internal/owtp"
import "fmt"

type PortStateSensor struct {
	SensorBase         // Incorpora o comportamento de SensorBase
	StockLevelReceived bool
	owtp.Schema
}

// catchEvent implementa o método para TemperatureSensor
func (t *PortStateSensor) catchEvent() owtp.Schema {
	// Implementação específica para TemperatureSensor
	fmt.Println("Event captured by PortState Sensor")
	return owtp.Schema{}
}
