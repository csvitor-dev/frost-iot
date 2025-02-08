package device

import (
	"fmt"

	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/types"
)

// throwTemperatureEvent simulates an event related to temperature changes.
func (r *Refrigerator) ThrowTemperatureEvent() {
	message := owtp.NewSchema(owtp.Header{
		Type: "sensor_data",
		Id:   types.NewUUID(),
		Role: "sensor",
	}, req.TemperatureRequest{
		Temperature: r.temperature,
	})

	fmt.Println("[ThrowTemperatureEvent] Event trigged!")

	r.tempSensor.LoadMessage(message)
}

// throwPortEvent simulates an event related to port state changes.
func (r *Refrigerator) ThrowPortEvent() {
	message := owtp.NewSchema(owtp.Header{
		Type: "sensor_data",
		Id:   types.NewUUID(),
		Role: "sensor",
	}, req.PortStateRequest{
		PortState: r.portState,
	})

	fmt.Println("[ThrowPortEvent] Event trigged!")

	r.portSensor.LoadMessage(message)
}

// throwStockEvent simulates an event related to stock level changes.
func (r *Refrigerator) ThrowStockEvent() {
	message := owtp.NewSchema(owtp.Header{
		Type: "sensor_data",
		Id:   types.NewUUID(),
		Role: "sensor",
	}, req.StockLevelRequest{
		StockLevel: r.stock,
	})

	fmt.Println("[ThrowStockEvent] Event trigged!")

	r.stockSensor.LoadMessage(message)
}
