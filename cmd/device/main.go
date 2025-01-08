package main

import (
	"fmt"

	"github.com/csvitor-dev/frost-iot/src/device"
)

func main() {
	d, err := device.NewRefrigeratorDevice(8.0, 0.1)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	d.SetPortState()
	fmt.Println(d)
}
