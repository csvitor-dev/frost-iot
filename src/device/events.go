package device

import (
	"fmt"
	"time"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/types"
)

// throwTemperatureEvent simulates an event related to temperature changes.
func (r *Refrigerator) ThrowTemperatureEvent() {
	message := owtp.Schema{
		Header: owtp.Header{
			Type: "sensor_data",
			Id: types.NewUUID(),
			Role: "sensor",
		},
		BodyMessage: owtp.TemperatureSensorBody{
			Temperature: r.temperature,
		},
		Timestamp: time.Now(),
	}

	fmt.Println("[ThrowTemperatureEvent] Event trigged!")

	r.tempSensor.LoadMessage(message)
}

// throwPortEvent simulates an event related to port state changes.
func (r *Refrigerator) ThrowPortEvent() {
	message := owtp.Schema{
		Header: owtp.Header{
			Type: "sensor_data",
			Id: types.NewUUID(),
			Role: "sensor",
		},
		BodyMessage: owtp.PortStateSensorBody{
			PortState: r.portState,
		},
		Timestamp: time.Now(),
	}

	fmt.Println("[ThrowPortEvent] Event trigged!")

	r.portSensor.LoadMessage(message)
}

// throwStockEvent simulates an event related to stock level changes.
func (r *Refrigerator) ThrowStockEvent() {
	message := owtp.Schema{
		Header: owtp.Header{
			Type: "sensor_data",
			Id: types.NewUUID(),
			Role: "sensor",
		},
		BodyMessage: owtp.StockLevelSensorBody{
			StockLevel: r.stock,
		},
		Timestamp: time.Now(),
	}

	fmt.Println("[ThrowStockEvent] Event trigged!")

	r.stockSensor.LoadMessage(message)
}
