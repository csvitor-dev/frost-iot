package main

import (
	"fmt"

	"github.com/csvitor-dev/frost-iot/internal/socket/process"
	"github.com/csvitor-dev/frost-iot/src/manager"
	"github.com/csvitor-dev/socket.go/src/tcp"
)

func main() {
	server, err := tcp.NewTCPServer("localhost:8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	serverManager := manager.NewServerManager(0.0, 0.0)
	server.InstallProcess(&process.OwtpProcess{
		Interceptor: manager.NewServerInterceptor(serverManager),
	})
	server.ListenAndServe()
}
