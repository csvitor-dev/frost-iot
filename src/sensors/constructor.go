package sensors

import (
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/pkg/types"
)

// returns the reference an sensor base with injectable struct
func NewSensor(kind string, inject types.SensorApplication) owtp.Sensor {
	return &SensorBase{
		id:  types.NewUUID(),
		kind: kind,
		lastMessages: []*owtp.Schema{}, 
		server: nil,
		children: inject,
	}
}