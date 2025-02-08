package process

import (
	"log"

	"github.com/csvitor-dev/frost-iot/internal/types"
)

type OwtpProcess struct {
	types.Interceptor
}

func (p *OwtpProcess) Push(request []byte) []byte {
	var response []byte

	if p.Interceptor == nil {
		return response
	}
	response, err := p.Invoke(request)

	if err != nil {
		log.Println(err)
	}
	return response
}
