package manager

import (
	"errors"
	"time"

	consts "github.com/csvitor-dev/frost-iot/internal/messages/constants"
	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	res "github.com/csvitor-dev/frost-iot/internal/messages/responses"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/socket"
)

func (s *ServerInterceptor) clientRequestHandler(request req.ClientRequest) ([]byte, error) {
	switch consts.Actions[request.Action] {
	case 1:
		return s.findRecord("temp_sensor")
	case 2:
		return s.findRecord("stock_sensor")
	case 3:
		return s.findRecord("port_sensor")
	case 4:
		return s.configureServer(request, s.manager.ConfigureTemperatureLimit)
	case 5:
		return s.configureServer(request, s.manager.ConfigureOpenPortTime)
	default:
		return nil, nil
	}
}

func (s *ServerInterceptor) findRecord(sensorRole string) ([]byte, error) {
	for i := len(s.manager.recentRecords) - 1; i >= 0; i-- {
		record := s.manager.recentRecords[i]

		if record.Role != sensorRole {
			continue
		}
		sensorBody, ok := record.Body.(req.SensorRequest)

		if !ok {
			return nil, errors.New("parse fail")
		}
		schema := owtp.NewSchema(owtp.Header{
			Type: "response",
			Role: "server",
		}, res.ClientResponse{
			Temperature: sensorBody.Temperature,
			StockLevel:  sensorBody.StockLevel,
			PortState:   sensorBody.PortState,
		})
		payload, err := socket.ParseSchemaToByte(schema)

		if err != nil {
			return nil, err
		}
		response, err := res.ParseBaseResponseToByte(res.BaseResponse{
			Payload: payload,
		})

		if err != nil {
			return nil, err
		}
		return response, nil
	}
	return nil, nil
}

func (s *ServerInterceptor) configureServer(request req.ClientRequest, config func(float64)) ([]byte, error) {
	err := request.Validate()

	if err != nil {
		return nil, err
	}
	config(request.Config)
	update := time.Now()

	payload, err := socket.ParseSchemaToByte(owtp.NewSchema(owtp.Header{
		Type: "response",
		Role: "server",
	}, res.ClientResponse{
		State:      nil,
		LastUpdate: update,
	}))

	if err != nil {
		return nil, err
	}
	response, err := res.ParseBaseResponseToByte(res.BaseResponse{
		Payload: payload,
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}
