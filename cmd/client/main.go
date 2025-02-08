package main

import (
	"github.com/csvitor-dev/frost-iot/src/client"
)

func main() {
	client := client.NewClientManager()

	client.InitializeView()
}
