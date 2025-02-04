package socket

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
)

var (
	errSchemaToByte = errors.New(`cannot convert data from 'Schema' to '[]byte'`)
	errByteToSchema = errors.New(`cannot convert data from '[]byte' to 'Schema'`)
)

func init() {
	gob.Register(time.Time{})
	gob.Register(owtp.Header{})
	gob.Register(owtp.Schema{})
}

func ParseSchemaToByte(message *owtp.Schema) ([]byte, error) {
	var network bytes.Buffer
	encoder := gob.NewEncoder(&network)

	err := encoder.Encode(message)

	if err != nil {
		return []byte{}, errSchemaToByte
	}
	return network.Bytes(), nil
}

func ParseByteToSchema(data []byte, retrieve *owtp.Schema) error {
	network := bytes.NewReader(data)
	decoder := gob.NewDecoder(network)

	err := decoder.Decode(retrieve)

	if err != nil {
		return errByteToSchema
	}
	return nil
}
