package device

import "errors"

func validateTemperature(value float64) error {
	if value >= 15.0 || value < -20.0 {
		return errors.New("temperature value outside of the acceptable range")
	}
	return nil
}

func validateStock(value float32) error {
	if value > 1 || value < 0 {
		return errors.New("stock percentage outside of the acceptable range")
	}
	return nil
}