package manager

import (
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/pkg/types"
)

type ServerManager struct {
	temperatureTarget float64
	timeTarget        float64
	recentRecords     map[types.UUID]owtp.Schema
}

