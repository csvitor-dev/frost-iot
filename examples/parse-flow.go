package main

import (
	"fmt"
	"time"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/socket"
	"github.com/csvitor-dev/frost-iot/internal/types"
)

func init() {
	socket.InjectBodyMessage(SensorBody{})
}

type SensorBody struct {
	Temperature float64
	Humidity    float32
}

func main() {
	schema := &owtp.Schema{
		Header: owtp.Header{
			Type: "sensor_data",
			Id: types.NewUUID(),
			Role: "sensor",
		},
		BodyMessage: SensorBody{
			Temperature: 4.0,
			Humidity: 65,
		},
		Timestamp: time.Now(),
	}
	fmt.Printf("Created: %v\n", schema)

	bytes, err := socket.ParseSchemaToByte(schema)

	if err!= nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Sending...\n%v\n", bytes)
	
	var output owtp.Schema

	err = socket.ParseByteToSchema(bytes, &output)

	if err!= nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Parsing...\n%v\n", output)
}
