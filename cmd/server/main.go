package main

import (
	"github.com/csvitor-dev/socket.go/src/types/tcp"
	"github.com/csvitor-dev/socket.go/src/runner"
)

func main() {
	runner.RunServer(&tcp.TCPServer{}, "localhost:8080")
}
