package sensors

import (
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/types"
	pkg "github.com/csvitor-dev/frost-iot/pkg/types"
)

// returns the reference an sensor base with injectable struct
func NewSensor[T owtp.BodyMessage](kind string, inject pkg.SensorApplication[T]) pkg.Sensor[T] {
	return &SensorBase[T]{
		id:           types.NewUUID(),
		kind:         kind,
		lastMessages: []owtp.Schema[T]{},
		server:       nil,
		children:     inject,
	}
}
