package actuators

import (
	"errors"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/pkg/types"
	"github.com/csvitor-dev/frost-iot/src/manager"
)

type ActuatorBase struct {
	id     types.UUID
	state  bool
	server *manager.ServerManager
	children types.ActuatorApplication
}

/* getters */

func (a *ActuatorBase) GetId() types.UUID {
	return a.id
}

func (a *ActuatorBase) GetState() bool {
	return a.state
}

/* setters */

func (a *ActuatorBase) SetState() {
	a.state = !a.state
}

/* methods */

func (a *ActuatorBase) HandleEvent(message *owtp.Schema) {
	a.children.HandleMessage(message)
}

func (a *ActuatorBase) Connect(server *manager.ServerManager) error {
	if (server == nil) {
		return errors.New("Fail")
	}
	a.server = server

	return nil
}

func (a *ActuatorBase) Ping() bool {
	return false
}
