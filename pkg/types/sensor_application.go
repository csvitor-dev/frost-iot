package types

import "github.com/csvitor-dev/frost-iot/internal/owtp"

type SensorApplication[T owtp.BodyMessage] interface {
	CatchEvent() owtp.Schema[T]
}
