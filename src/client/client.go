package client

import (
	"fmt"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

type ClientManager struct {
	lastReceive owtp.Schema
}

// Function to receive the last temperature record
func (cm *ClientManager) ReceiveLastTemperatureRecord() float64 {
	// Placeholder: Replace with actual logic to fetch the last temperature record
	return cm.lastReceive.GetTemperature()
}

// Function to receive the last stock level record
func (cm *ClientManager) ReceiveLastStockLevelRecord() float32 {
	// Placeholder: Replace with actual logic to fetch the last stock level record
	return cm.lastReceive.GetStockLevel()
}

// Function to receive the last open port record
func (cm *ClientManager) ReceiveLastOpenPortRecord() bool {
	// Placeholder: Replace with actual logic to fetch the last open port status
	return cm.lastReceive.GetOpenPort()
}

// Method to configure the temperature limit
func (cm *ClientManager) ConfigureTemperatureLimit(limit float64) {
	cm.lastReceive.SetTemperatureLimit(limit)
	fmt.Printf("Temperature limit configured to: %.2f\n", limit)
}

// Method to configure the open port time
func (cm *ClientManager) ConfigureOpenPortTime(time float64) {
	cm.lastReceive.SetOpenPortTime(time)
	fmt.Printf("Open port time configured to: %.2f\n", time)
}
