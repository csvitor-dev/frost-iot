package transporters

import (
	"errors"
	"fmt"

	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/socket"
	"github.com/csvitor-dev/socket.go/src/tcp"
)

func SensorTransporter[T owtp.BodyMessage](request owtp.Schema[T]) error {
	conn, err := tcp.NewTCPConnection(":8080")

	if err != nil {
		return errors.New("connection fail")
	}
	defer conn.Close()

	payload, err := socket.ParseSchemaToByte(request)

	if err != nil {
		return err
	}
	bytes, err := req.ParseBaseRequestToByte(req.BaseRequest{
		Role:    "sensor",
		Payload: payload,
	})

	if err != nil {
		return err
	}
	err = conn.SendMessage(bytes)

	if err != nil {
		return err
	}
	_, err = conn.RetriveMessage()

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
