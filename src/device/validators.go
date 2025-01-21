package device

import "errors"

var (
	errTemperatureOutRange = errors.New("temperature value outside of the acceptable range")
	errStockOutRange = errors.New("stock percentage outside of the acceptable range")
)

func validateTemperature(value float64) error {
	if value >= 15.0 || value < -20.0 {
		return errTemperatureOutRange
	}
	return nil
}

func validateStock(value float32) error {
	if value > 1 || value < 0 {
		return errStockOutRange
	}
	return nil
}