package client

import (
	"fmt"
	"time"

	req "github.com/csvitor-dev/frost-iot/internal/messages/requests"
	res "github.com/csvitor-dev/frost-iot/internal/messages/responses"
	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/socket"
	"github.com/csvitor-dev/frost-iot/internal/types"
	"github.com/csvitor-dev/frost-iot/src/view"
)

type ClientManager struct {
	uuid        types.UUID
	lastReceive owtp.Schema[res.ClientResponse]
}

/* getter */

func NewClientManager() *ClientManager {
	return &ClientManager{}
}

func (cm *ClientManager) InitializeView() {
	for {
		view.ShowMenu()
		choice := view.GetUserChoice()

		switch choice {
		case 1:
			view.ShowTemperatureRecord(cm.ReceiveLastTemperatureRecord())
		case 2:
			view.ShowStockLevelRecord(cm.ReceiveLastStockLevelRecord())
		case 3:
			view.ShowOpenPortRecord(cm.ReceiveLastOpenPortRecord())
		case 4:
			cm.ConfigureTemperatureLimit(view.GetTemperatureLimit())
		case 5:
			cm.ConfigureOpenPortTime(view.GetOpenPortTime())
		case 6:
			view.ShowExitMessage()
			return
		default:
			view.ShowInvalidChoice()
		}
	}

}

func (cm *ClientManager) GetReceive() owtp.Schema[res.ClientResponse] {
	return cm.lastReceive
}

// receive the last temperature record
func (cm *ClientManager) ReceiveLastTemperatureRecord() float64 {
	request := owtp.Schema[req.ClientRequest]{
		Header: owtp.Header{
			Type: "query",
			Id:   cm.uuid,
		},
		Body: req.ClientRequest{
			Action: "GET_TEMPERATURE",
		},
		Timestamp: time.Now(),
	}
	response, err := socket.ClientTransporter(request)

	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return 0.0
	}

	return response.Body.Temperature
}

// receive the last stock level record
func (cm *ClientManager) ReceiveLastStockLevelRecord() float32 {
	request := owtp.Schema[req.ClientRequest]{
		Header: owtp.Header{
			Type: "query",
			Id:   cm.uuid,
		},
		Body: req.ClientRequest{
			Action: "GET_STOCK",
		},
		Timestamp: time.Now(),
	}
	response, err := socket.ClientTransporter(request)

	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return 0.0
	}

	return response.Body.StockLevel
}

// receive the last open port record
func (cm *ClientManager) ReceiveLastOpenPortRecord() bool {
	request := owtp.Schema[req.ClientRequest]{
		Header: owtp.Header{
			Type: "query",
			Id:   cm.uuid,
		},
		Body: req.ClientRequest{
			Action: "GET_PORT",
		},
		Timestamp: time.Now(),
	}
	response, err := socket.ClientTransporter(request)

	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return false
	}

	return response.Body.PortState
}

// Method to configure the temperature limit
func (cm *ClientManager) ConfigureTemperatureLimit(limit float64) {
	request := owtp.Schema[req.ClientConfigRequest]{
		Header: owtp.Header{
			Type: "query",
			Id:   cm.uuid,
		},
		Body: req.ClientConfigRequest{
			ClientRequest: req.ClientRequest{
				Action: "CONFIG_TEMP",
			},
			Config: limit,
		},
		Timestamp: time.Now(),
	}
	err := socket.ConfigTransporter(request)

	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return
	}
	fmt.Println("[SUCCESS] configured temperature limit")
}

// Method to configure the open port time
func (cm *ClientManager) ConfigureOpenPortTime(limit float64) {
	request := owtp.Schema[req.ClientConfigRequest]{
		Header: owtp.Header{
			Type: "query",
			Id:   cm.uuid,
		},
		Body: req.ClientConfigRequest{
			ClientRequest: req.ClientRequest{
				Action: "CONFIG_TIME",
			},
			Config: limit,
		},
		Timestamp: time.Now(),
	}
	err := socket.ConfigTransporter(request)

	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return
	}
	fmt.Println("[SUCCESS] configured temperature limit")
}
