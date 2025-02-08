package types

import "github.com/csvitor-dev/frost-iot/internal/owtp"

type Actuator interface {
	HandleEvent(owtp.Schema[owtp.BodyMessage])
}
