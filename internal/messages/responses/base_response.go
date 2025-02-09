package responses

import (
	"bytes"
	"encoding/gob"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/socket"
)

func init() {
	owtp.Inject(BaseResponse{})
}

type BaseResponse struct {
	Payload []byte
}

func (r BaseResponse) Validate() error {
	return nil
}

func ParseBaseResponseToByte(response BaseResponse) ([]byte, error) {
	var network bytes.Buffer
	encoder := gob.NewEncoder(&network)

	err := encoder.Encode(&response)

	if err != nil {
		return []byte{}, nil
	}
	return network.Bytes(), nil
}

func ParseByteToBaseResponse[T owtp.BodyMessage](data []byte) (BaseResponse, error) {
	network := bytes.NewReader(data)
	decoder := gob.NewDecoder(network)

	var retrieve BaseResponse
	err := decoder.Decode(&retrieve)

	if err != nil {
		return retrieve, nil
	}
	retrieve.Payload, err = buildResponseTo[T](retrieve.Payload)

	if err != nil {
		return retrieve, err
	}
	return retrieve, nil
}

func buildResponseTo[T owtp.BodyMessage](payload []byte) ([]byte, error) {
	schema, err := socket.ParseByteToSchema[T](payload)

	if err != nil {
		return nil, err
	}

	return socket.ParseSchemaToByte(schema)
}
