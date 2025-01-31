package owtp

import (
	"time"

	"github.com/csvitor-dev/frost-iot/internal/types"
)

type Header struct {
	Type string
	Id types.UUID
	Role string
}

type BodyMessage interface {
}

// Schema defines the structure for the OWTP protocol.
type Schema struct {
	Header
	BodyMessage
	Timestamp time.Time
}
