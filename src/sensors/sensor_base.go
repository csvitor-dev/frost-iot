package sensors

import (
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/pkg/types"
	"github.com/csvitor-dev/frost-iot/src/manager"
)

import "fmt"

type SensorBase struct {
	Id           types.UUID
	Kind         string
	LastMessages []owtp.Schema
	Server       manager.ServerManager
}

func (t *SensorBase) catchEvent() owtp.Schema {
	// Implementação específica para TemperatureSensor
	fmt.Println("Event captured by Sensor")
	return owtp.Schema{}
}

func (t *SensorBase) SendMessages() {
	fmt.Println(t.LastMessages)
}

func (t *SensorBase) Connect(server manager.ServerManager) {

}

func (t *SensorBase) Ping() bool {
	return true
}
