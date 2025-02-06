package types

import "github.com/csvitor-dev/frost-iot/internal/owtp"

type Sensor interface {
	CatchEvent() owtp.Schema
	SendMessages()
	LoadMessage(owtp.Schema)
}
