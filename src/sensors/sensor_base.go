package sensors

import (
	"fmt"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/types"
	pkg "github.com/csvitor-dev/frost-iot/pkg/types"
	"github.com/csvitor-dev/frost-iot/src/manager"
)

type SensorBase[T owtp.BodyMessage] struct {
	id           types.UUID
	kind         string
	lastMessages []owtp.Schema[T]
	server       *manager.ServerManager
	children     pkg.SensorApplication[T]
}

func (s *SensorBase[T]) CatchEvent() owtp.Schema[T] {
	fmt.Println("[SensorBase] Event captured by Sensor")

	return s.children.CatchEvent()
}

func (s *SensorBase[T]) LoadMessage(message owtp.Schema[T]) {
	s.lastMessages = append(s.lastMessages, message)

	fmt.Println(s.lastMessages)
}

func (s *SensorBase[T]) SendMessages() {
	fmt.Println("All messages:", s.lastMessages)
}

func (s *SensorBase[T]) Connect(server manager.ServerManager) {

}

func (s *SensorBase[T]) Ping() bool {
	return true
}
