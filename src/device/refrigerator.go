package device

import (
	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	pkg "github.com/csvitor-dev/frost-iot/pkg/types"
	"github.com/csvitor-dev/frost-iot/src/sensors"
)

type Refrigerator struct {
	temperature float64
	stock       float32
	portState   bool
	stockSensor pkg.Sensor[req.StockLevelRequest]
	tempSensor  pkg.Sensor[req.TemperatureRequest]
	portSensor  pkg.Sensor[req.PortStateRequest]
}

/* constructor */

func NewRefrigeratorDevice(temperature float64, stock float32) (*Refrigerator, error) {
	if err := validate(temperature, stock); err != nil {
		return &Refrigerator{}, err
	}

	return &Refrigerator{
		temperature: temperature,
		stock:       stock,
		portState:   false,
		tempSensor:  createSensor[req.TemperatureRequest](&sensors.TemperatureSensor{}),
		stockSensor: createSensor[req.StockLevelRequest](&sensors.StockLevelSensor{}),
		portSensor:  createSensor[req.PortStateRequest](&sensors.PortStateSensor{}),
	}, nil
}
func validate(temperature float64, stock float32) error {
	err := validateTemperature(temperature)

	if err != nil {
		return err
	}
	err = validateStock(stock)

	if err != nil {
		return err
	}
	return nil
}
func createSensor[T owtp.BodyMessage](inject pkg.SensorApplication[T]) pkg.Sensor[T] {
	return sensors.NewSensor[T]("sensor", inject)
}

/* Getters */

func (r *Refrigerator) GetTemperature() float64 {
	return r.temperature
}

func (r *Refrigerator) GetStock() float32 {
	return r.stock
}

func (r *Refrigerator) GetPortState() bool {
	return r.portState
}

/* Setters */

func (r *Refrigerator) SetTemperature(value float64) error {
	if err := validateTemperature(value); err != nil {
		return err
	}

	r.temperature = value
	return nil
}

func (r *Refrigerator) SetStock(value float32) error {
	if err := validateStock(value); err != nil {
		return err
	}

	r.stock = value
	return nil
}

func (r *Refrigerator) SetPortState() {
	r.portState = !r.portState
}
