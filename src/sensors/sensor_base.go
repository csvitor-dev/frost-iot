package sensors

import (
	"fmt"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
	trprt "github.com/csvitor-dev/frost-iot/internal/socket/transporters"
	"github.com/csvitor-dev/frost-iot/internal/types"
	pkg "github.com/csvitor-dev/frost-iot/pkg/types"
)

type SensorBase[T owtp.BodyMessage] struct {
	uuid         types.UUID
	kind         string
	lastMessages []owtp.Schema[T]
	children     pkg.SensorApplication[T]
}

func (s *SensorBase[T]) CatchEvent() owtp.Schema[T] {
	fmt.Println("[SensorBase] Event captured by Sensor")

	return s.children.CatchEvent()
}

func (s *SensorBase[T]) LoadMessage(message owtp.Schema[T]) {
	s.lastMessages = append(s.lastMessages, message)
}

func (s *SensorBase[T]) SendMessages() {
	for _, message := range s.lastMessages {
		err := trprt.SensorTransporter(message)

		if err != nil {
			fmt.Println("fail to send message")
			return
		}
	}
	s.lastMessages = []owtp.Schema[T]{}
}

func (s *SensorBase[T]) Role() string {
	return s.children.Role()
}
