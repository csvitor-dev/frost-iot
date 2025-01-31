package types

import "github.com/csvitor-dev/frost-iot/internal/owtp"

type SensorApplication interface {
	CatchEvent() *owtp.Schema
}