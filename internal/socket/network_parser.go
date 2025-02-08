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
	gob.Register(owtp.Schema[owtp.BodyMessage]{})
}

func ParseSchemaToByte[T owtp.BodyMessage](message owtp.Schema[T]) ([]byte, error) {
	var network bytes.Buffer
	encoder := gob.NewEncoder(&network)

	err := encoder.Encode(&message)

	if err != nil {
		return []byte{}, errSchemaToByte
	}
	return network.Bytes(), nil
}

func ParseByteToSchema[T owtp.BodyMessage](data []byte) (owtp.Schema[T], error) {
	network := bytes.NewReader(data)
	decoder := gob.NewDecoder(network)

	var retrieve owtp.Schema[T]
	err := decoder.Decode(&retrieve)

	if err != nil {
		return retrieve, errByteToSchema
	}
	return retrieve, nil
}
