package responses

import (
	"errors"
	"time"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/types"
)

func init() {
	owtp.Inject(ClientResponse{})
}

type ClientResponse struct {
	SensorId    types.UUID
	Temperature float64
	StockLevel  float32
	PortState   bool
	LastUpdate  time.Time
	State       error
}

func (r ClientResponse) Validate() error {
	if len(r.SensorId) == 0 {
		return errors.New("sensor not provided")
	}
	if r.Temperature < 0.0 || r.Temperature > 4.0 {
		return errors.New("invalid temperature")
	}
	if r.StockLevel < 0.0 || r.StockLevel > 1.0 {
		return errors.New("invalid stock level")
	}
	return nil
}
