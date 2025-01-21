package device

type Refrigerator struct {
	temperature float64
	stock       float32
	portState   bool
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
