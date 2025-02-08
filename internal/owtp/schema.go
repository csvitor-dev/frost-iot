package owtp

import (
	"time"

	"github.com/csvitor-dev/frost-iot/internal/types"
)

type Header struct {
	Type string
	Id   types.UUID
	Role string
}

type BodyMessage interface {
	Validate() error
}

// Schema defines the structure for the OWTP protocol.
type Schema[T BodyMessage] struct {
	Header
	Body      T
	Timestamp time.Time
}
