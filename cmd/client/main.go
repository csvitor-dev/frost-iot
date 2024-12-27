package main

import (
	"github.com/csvitor-dev/socket.go/src/runner"
	"github.com/csvitor-dev/socket.go/src/types/tcp"
)

func main() {
	runner.RunClient(&tcp.TCPConnection{}, ":8080")
}
