package types

import "github.com/csvitor-dev/frost-iot/internal/owtp"

type ActuatorApplication interface {
	HandleMessage(*owtp.Schema)
}
