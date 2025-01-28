package device

import "fmt"

// throwTemperatureEvent simulates an event related to temperature changes.
func (r *Refrigerator) ThrowTemperatureEvent() {
	fmt.Printf("Temperature event triggered! Current temperature: %.2f°C\n", r.temperature)
}

// throwPortEvent simulates an event related to port state changes.
func (r *Refrigerator) ThrowPortEvent() {
	state := "closed"
	if r.portState {
		state = "open"
	}
	fmt.Printf("Port event triggered! Current port state: %s\n", state)
}

// throwStockEvent simulates an event related to stock level changes.
func (r *Refrigerator) ThrowStockEvent() {
	fmt.Printf("Stock event triggered! Current stock level: %.2f units\n", r.stock)
}
