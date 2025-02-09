package transporters

import (
	"errors"

	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	res "github.com/csvitor-dev/frost-iot/internal/messages/responses"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/socket"
	"github.com/csvitor-dev/socket.go/src/tcp"
)

func ClientTransporter[T, U owtp.BodyMessage](request owtp.Schema[T]) (owtp.Schema[U], error) {
	var response owtp.Schema[U]
	conn, err := tcp.NewTCPConnection(":8080")

	if err != nil {
		return response, errors.New("connection fail")
	}
	defer conn.Close()

	payload, err := socket.ParseSchemaToByte(request)

	if err != nil {
		return response, err
	}
	bytes, err := req.ParseBaseRequestToByte(req.BaseRequest{
		Role:    "client",
		Payload: payload,
	})

	if err != nil {
		return response, err
	}
	err = conn.SendMessage(bytes)

	if err != nil {
		return response, err
	}
	bytes, err = conn.RetriveMessage()

	if err != nil {
		return response, err
	}
	base, err := res.ParseByteToBaseResponse[U](bytes)

	if err != nil {
		return response, err
	}
	response, err = socket.ParseByteToSchema[U](base.Payload)

	if err != nil {
		return response, err
	}

	return response, nil
}
