package owtp

import (
	"time"
)

func NewSchema[T BodyMessage](header Header, body T) Schema[T] {
	return Schema[T]{
		Header:    header,
		Body:      body,
		Timestamp: time.Now(),
	}
}
