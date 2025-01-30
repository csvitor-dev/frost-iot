package sensors

import (
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/pkg/types"
)

type SensorApplication interface {
	CatchEvent() *owtp.Schema
}

// returns the reference an sensor base with injectable struct
func NewSensor(kind string, inject SensorApplication) owtp.Sensor {
	return &SensorBase{
		id:  types.NewUUID(),
		kind: kind,
		lastMessages: []*owtp.Schema{}, 
		server: nil,
		children: inject,
	}
}