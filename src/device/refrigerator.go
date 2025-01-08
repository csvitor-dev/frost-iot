package device

type refrigerator struct {
	temperature float64
	stock       float32
	portState   bool
}

/* constructor */

func NewRefrigeratorDevice(temperature float64, stock float32) (*refrigerator, error) {
	if err := validate(temperature, stock); err != nil {
		return &refrigerator{}, err
	}

	return &refrigerator{
		temperature: temperature,
		stock:       stock,
		portState:   false,
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

/* Getters */

func (r *refrigerator) GetTemperature() float64 {
	return r.temperature
}

func (r *refrigerator) GetStock() float32 {
	return r.stock
}

func (r *refrigerator) GetPortState() bool {
	return r.portState
}

/* Setters */

func (r *refrigerator) SetTemperature(value float64) error {
	if err := validateTemperature(value); err != nil {
		return err
	}

	r.temperature = value
	return nil
}

func (r *refrigerator) SetStock(value float32) error {
	if err := validateStock(value); err != nil {
		return err
	}

	r.stock = value
	return nil
}

func (r *refrigerator) SetPortState() {
	r.portState = !r.portState
}
