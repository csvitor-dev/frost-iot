package requests

import "github.com/csvitor-dev/frost-iot/internal/owtp"

func init() {
	owtp.Inject(TemperatureRequest{})
	owtp.Inject(StockLevelRequest{})
	owtp.Inject(PortStateRequest{})
}

type TemperatureRequest struct {
	Temperature float64
}

func (r TemperatureRequest) Validate() error {
	return nil
}

type StockLevelRequest struct {
	StockLevel float32
}

func (r StockLevelRequest) Validate() error {
	return nil
}

type PortStateRequest struct {
	PortState bool
}

func (r PortStateRequest) Validate() error {
	return nil
}
