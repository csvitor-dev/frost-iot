package sensors

import "github.com/csvitor-dev/frost-iot/internal/owtp"
import "fmt"

type StockLevelSensor struct {
	SensorBase         // Incorpora o comportamento de SensorBase
	StockLevelReceived float64
	owtp.Schema
}

// catchEvent implementa o método para TemperatureSensor
func (t *StockLevelSensor) catchEvent() owtp.Schema {
	// Implementação específica para TemperatureSensor
	fmt.Println("Event captured by StockLevel Sensor")
	return owtp.Schema{}
}
