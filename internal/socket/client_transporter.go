package socket

import (
	"errors"

	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	res "github.com/csvitor-dev/frost-iot/internal/messages/responses"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/socket.go/src/tcp"
)

func ClientTransporter(request owtp.Schema[req.ClientRequest]) (owtp.Schema[res.ClientResponse], error) {
	var response owtp.Schema[res.ClientResponse]
	conn, err := tcp.NewTCPConnection(":8080")

	if err != nil {
		return response, errors.New("connection fail")
	}
	defer conn.Close()

	bytes, err := ParseSchemaToByte(request)
	
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
	response, err = ParseByteToSchema[res.ClientResponse](bytes)

	if err != nil {
		return response, err
	}

	return response, nil
}

func ConfigTransporter(request owtp.Schema[req.ClientConfigRequest]) error {
	conn, err := tcp.NewTCPConnection(":8080")

	if err != nil {
		return errors.New("connection fail")
	}
	defer conn.Close()

	bytes, err := ParseSchemaToByte(request)
	
	if err != nil {
		return err
	}
	err = conn.SendMessage(bytes)

	if err != nil {
		return err
	}
	bytes, err = conn.RetriveMessage()

	if err != nil {
		return err
	}
	response, err := ParseByteToSchema[res.ClientConfigResponse](bytes)

	if err != nil {
		return err
	}
	return response.Body.State
}
