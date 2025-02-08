package requests

import (
	"errors"

	"github.com/csvitor-dev/frost-iot/internal/messages/constants"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

func init() {
	owtp.Inject(ClientRequest{})
}

type ClientRequest struct {
	Action string
	Config float64
}

func (r ClientRequest) Validate() error {
	if constants.Actions[r.Action] == 0 {
		return errors.New("action there no exist")
	}
	config := constants.Actions[r.Action]

	switch config {
	case 4:
		if r.Config < 0.0 || r.Config > 4.0 {
			return errors.New("invalid temperature limit")
		}
	case 5:
		if r.Config < 0.0 || r.Config > 30.0 {
			return errors.New("invalid port time limit")
		}
	}
	return nil
}
