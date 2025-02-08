package requests

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

func init() {
	owtp.Inject(BaseRequest{})
}

type BaseRequest struct {
	Role    string
	Payload []byte
}

func (r BaseRequest) Validate() error {
	for _, role := range []string{"actuator", "client", "sensor"} {
		if r.Role == role {
			return nil
		}
	}
	return errors.New("fail")
}

var (
	errRequestToByte = errors.New(`cannot convert data from 'BaseRequest' to '[]byte'`)
	errByteToRequest = errors.New(`cannot convert data from '[]byte' to 'BaseRequest'`)
)

func ParseBaseRequestToByte(request BaseRequest) ([]byte, error) {
	var network bytes.Buffer
	encoder := gob.NewEncoder(&network)

	err := encoder.Encode(&request)

	if err != nil {
		return []byte{}, errRequestToByte
	}
	return network.Bytes(), nil
}

func ParseByteToBaseRequest(data []byte) (BaseRequest, error) {
	network := bytes.NewReader(data)
	decoder := gob.NewDecoder(network)

	var retrieve BaseRequest
	err := decoder.Decode(&retrieve)

	if err != nil {
		return retrieve, errByteToRequest
	}
	return retrieve, nil
}
