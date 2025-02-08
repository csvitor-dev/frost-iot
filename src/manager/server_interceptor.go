package manager

import (
	"errors"

	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	"github.com/csvitor-dev/frost-iot/internal/socket"
)

type ServerInterceptor struct {
	manager *ServerManager
}

func NewServerInterceptor(manager *ServerManager) *ServerInterceptor {
	return &ServerInterceptor{
		manager: manager,
	}
}

func (s *ServerInterceptor) Invoke(request []byte) ([]byte, error) {
	var result []byte

	base, err := req.ParseByteToBaseRequest(request)

	if err != nil {
		return result, err
	}

	switch base.Role {
	case "actuator":
		// ...
	case "client":
		message, err := socket.ParseByteToSchema[req.ClientRequest](base.Payload)

		if err != nil {
			return result, err
		}
		result, err = socket.ParseSchemaToByte(message)

		if err != nil {
			return result, err
		}
	case "sensor":
		// ...
	default:
		return result, errors.New("server could not resolve request")
	}
	return result, nil
}
