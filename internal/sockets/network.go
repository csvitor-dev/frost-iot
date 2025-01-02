package sockets

import (
	"bytes"
	"encoding/gob"
)

func createNetwork() func() *bytes.Buffer {
	var network *bytes.Buffer = nil

	scope := func() *bytes.Buffer {
		if network != nil {
			return network
		}
		network = new(bytes.Buffer)
		return network
	}
	return scope
}

func CreateContext() func() *bytes.Buffer {
	return createNetwork()
}

func CreateEncoderChannel(context func() *bytes.Buffer) *gob.Encoder {
	network := context()
	return gob.NewEncoder(network)
}

func CreateDecoderChannel(context func() *bytes.Buffer) *gob.Decoder {
	network := context()
	return gob.NewDecoder(network)
}
