package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/csvitor-dev/frost-iot/internal/owtp"
	"github.com/csvitor-dev/frost-iot/internal/socket"
	"github.com/csvitor-dev/frost-iot/internal/types"
)

func init() {
	owtp.Inject(SensorBody{})
}

type SensorBody struct {
	Temperature float64
}

func (s SensorBody) Validate() error {
	if s.Temperature < 0.0 || s.Temperature > 4.0 {
		return errors.New("[ERROR] invalid temperature")
	}
	return nil
}

func main() {
	schema := owtp.Schema[SensorBody]{
		Header: owtp.Header{
			Type: "sensor_data",
			Id:   types.NewUUID(),
			Role: "sensor",
		},
		Body: SensorBody{
			Temperature: 4.0,
		},
		Timestamp: time.Now(),
	}
	fmt.Printf("Created: %v\n", schema)

	bytes, err := socket.ParseSchemaToByte(schema)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Sending...\n%v\n", bytes)

	output, err := socket.ParseByteToSchema[SensorBody](bytes)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Parsing...\n%v\n", output)
}
