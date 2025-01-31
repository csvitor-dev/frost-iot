package manager

import (
	"github.com/csvitor-dev/frost-iot/internal/types"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	pkg "github.com/csvitor-dev/frost-iot/pkg/types"
)

type ServerManager struct {
	temperatureTarget float64
	timeTarget        float32
	recentRecords     map[types.UUID]owtp.Schema
	sensors []pkg.Sensor
	actuators []pkg.Actuator
}

/* constructor */

func NewServerManger(temperature float64, time float32) *ServerManager {
	return &ServerManager{
		temperatureTarget: temperature,
		timeTarget: time,
		recentRecords: map[types.UUID]owtp.Schema {},
		sensors: nil,
		actuators: nil,
	}
}

/* setters */

func (s *ServerManager) ConfigureTemperatureLimit(target float64) {

}

func (s *ServerManager) ConfigureOpenPortTime(time float32) {

}

/* methods */

func (s *ServerManager) TryConnectSensor(uuid types.UUID) bool {
	return false
}

func (s *ServerManager) TryConnectActuator(uuid types.UUID) bool {
	return false
}

func (s *ServerManager) ReceiveRecord(uuid types.UUID, message owtp.Schema) {

}

func (s *ServerManager) IsRefrigeratorOn() bool {
	return false
}

func (s *ServerManager) IsLightOn() bool {
	return false
}

func (s *ServerManager) IsAlarmOn() bool {
	return false
}

func (s *ServerManager) ProvideRefrigeratorState() *owtp.Schema {
	return nil
}

func (s *ServerManager) ProvideSystemState() *owtp.Schema {
	return nil
}
