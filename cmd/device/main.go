package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/csvitor-dev/frost-iot/src/device"
)

func main() {
	device, err := device.NewRefrigeratorDevice(8.0, 0.1)

	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			time.Sleep(time.Second * 10)
			device.Trigger()
		}
	}()

	for {

		device.ThrowPortEvent()
		device.ThrowStockEvent()
		device.ThrowTemperatureEvent()

		if SetAvaliable() {
			device.SetTemperature(rand.Float64() * 8.0)
		}
		if SetAvaliable() && device.GetPortState() {
			device.SetStock(rand.Float32())
		}
		if SetAvaliable() {
			device.SetPortState()
		}

		time.Sleep(time.Second * 2)
	}

}

func SetAvaliable() bool {
	return rand.Float32() > 0.5
}
