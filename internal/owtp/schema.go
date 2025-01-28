package owtp

// Schema defines the structure for the OWTP protocol.
type Schema struct {
	temperature      float64 // Last recorded temperature
	stockLevel       float32 // Last recorded stock level
	openPort         bool    // Last recorded open port status
	temperatureLimit float64 // Configured temperature limit
	openPortTime     float64 // Configured open port time
}

// Getter for Temperature
func (s *Schema) GetTemperature() float64 {
	return s.temperature
}

// Getter for StockLevel
func (s *Schema) GetStockLevel() float32 {
	return s.stockLevel
}

// Getter for OpenPort
func (s *Schema) GetOpenPort() bool {
	return s.openPort
}

// Setter for TemperatureLimit
func (s *Schema) SetTemperatureLimit(limit float64) {
	s.temperatureLimit = limit
}

// Setter for OpenPortTime
func (s *Schema) SetOpenPortTime(time float64) {
	s.openPortTime = time
}