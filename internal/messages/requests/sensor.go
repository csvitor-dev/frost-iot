package requests

import "github.com/csvitor-dev/frost-iot/internal/owtp"

func init() {
	owtp.Inject(SensorRequest{})
}

type SensorRequest struct {
	Temperature float64
	StockLevel  float32
	PortState   bool
}

func (r SensorRequest) Validate() error {
	return nil
}
