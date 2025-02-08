package process

import (
	"log"

	"github.com/csvitor-dev/frost-iot/internal/types"
)

type OwtpProcess struct {
	types.Intercepter
}

func (p *OwtpProcess) Push(request []byte) []byte {
	var response []byte

	if p.Intercepter == nil {
		return response
	}
	response, err := p.Invoke(request)

	if err != nil {
		log.Println(err)
	}
	return response
}
