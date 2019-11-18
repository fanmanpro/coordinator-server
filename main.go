package main

import (
	"github.com/fanmanpro/coordinator-server/coordinator"
	"github.com/fanmanpro/coordinator-server/ws"
)

//var (
//	clients      map[int32]*Client
//	packetQueue  []Packet
//	clientsMutex sync.RWMutex
//)

func main() {
	coordinator := coordinator.NewCoordinator()

	localIP := "127.0.0.1"
	localPort := "1540"
	wsServer := ws.NewWebSocketServer(coordinator, localIP, localPort)
	wsServer.Start()
}
