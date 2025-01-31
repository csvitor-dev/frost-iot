package sensors

import (
	"fmt"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/pkg/types"
	"github.com/csvitor-dev/frost-iot/src/manager"
)

type SensorBase struct {
	id           types.UUID
	kind         string
	lastMessages []*owtp.Schema
	server       *manager.ServerManager
	children     types.SensorApplication
}

func (s *SensorBase) CatchEvent() *owtp.Schema {
	fmt.Println("[SensorBase] Event captured by Sensor")

	return s.children.CatchEvent()
}

func (s *SensorBase) SendMessages() {
	fmt.Println("All messages:", s.lastMessages)
}

func (s *SensorBase) Connect(server manager.ServerManager) {

}

func (s *SensorBase) Ping() bool {
	return true
}
