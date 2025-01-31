package sensors

import (
	"fmt"

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

func (s *StockLevelSensor) CatchEvent() *owtp.Schema {
	fmt.Println(" >> Event captured by StockLevel Sensor")
	return &owtp.Schema{}
}
