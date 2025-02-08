package owtp

import (
	"encoding/gob"
)

func Inject(body BodyMessage) {
	gob.Register(body)
}
