package types

import "github.com/csvitor-dev/frost-iot/internal/owtp"

type Sensor[T owtp.BodyMessage] interface {
	CatchEvent() owtp.Schema[T]
	SendMessages()
	LoadMessage(owtp.Schema[T])
	Role() string
}
