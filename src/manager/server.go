package manager

import (
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/types"
	pkg "github.com/csvitor-dev/frost-iot/pkg/types"
)

type ServerManager struct {
	temperatureTarget float64
	timeTarget        float32
	recentRecords     map[types.UUID]owtp.Schema[owtp.BodyMessage]
	sensors           []pkg.Sensor[owtp.BodyMessage]
	actuators         []pkg.Actuator
}

/* constructor */

func NewServerManager(temperature float64, time float32) *ServerManager {
	return &ServerManager{
		temperatureTarget: temperature,
		timeTarget:        time,
		recentRecords:     map[types.UUID]owtp.Schema[owtp.BodyMessage]{},
		sensors:           nil,
		actuators:         nil,
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

func (s *ServerManager) ReceiveRecord(uuid types.UUID, message owtp.Schema[owtp.BodyMessage]) {

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

func (s *ServerManager) ProvideRefrigeratorState() owtp.Schema[owtp.BodyMessage] {
	return owtp.Schema[owtp.BodyMessage]{}
}

func (s *ServerManager) ProvideSystemState() owtp.Schema[owtp.BodyMessage] {
	return owtp.Schema[owtp.BodyMessage]{}
}
