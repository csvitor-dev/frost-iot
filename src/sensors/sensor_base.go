package sensors

import (
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/pkg/types"
	"github.com/csvitor-dev/frost-iot/src/manager"
)

type SensorBase struct {
	Id           types.UUID
	Kind         string
	LastMessages []owtp.Schema
	Server       manager.ServerManager
}
