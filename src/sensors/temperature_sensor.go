package sensors

import "github.com/csvitor-dev/frost-iot/internal/owtp"
import "fmt"

type TemperatureSensor struct {
	SensorBase          // Incorpora o comportamento de SensorBase
	TemperatureReceived float64
	owtp.Schema
}

// catchEvent implementa o método para TemperatureSensor
func (t *TemperatureSensor) catchEvent() owtp.Schema {
	// Implementação específica para TemperatureSensor
	fmt.Println("Event captured by Temperature Sensor")
	return owtp.Schema{}
}

func (t *TemperatureSensor) SendMessages() {
	fmt.Println(t.LastMessages)
}