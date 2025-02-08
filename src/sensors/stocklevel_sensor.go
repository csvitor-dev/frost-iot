package sensors

import (
	"fmt"

	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

type StockLevelSensor struct {
	stockLevelReceived float32
}

/* getters */

func (s *StockLevelSensor) GetStockLevel() float32 {
	return s.stockLevelReceived
}

/* methods */

func (s *StockLevelSensor) CatchEvent() owtp.Schema[req.SensorRequest] {
	fmt.Println(" >> Event captured by StockLevel Sensor")
	return owtp.Schema[req.SensorRequest]{}
}

func (s *StockLevelSensor) Role() string {
	return "stock_sensor"
}
