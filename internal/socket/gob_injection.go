package socket

import (
	"encoding/gob"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

func InjectBodyMessage(body owtp.BodyMessage) {
	gob.Register(body)
}