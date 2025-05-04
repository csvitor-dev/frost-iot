package types

import "github.com/google/uuid"

// represents a UUID version
type Uuid string

func NewUuid() Uuid {
	return Uuid(uuid.NewString())
}