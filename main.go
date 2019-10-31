package main

import "github.com/fanmanpro/coordinator-server/ws"

//var (
//	clients      map[int32]*Client
//	packetQueue  []Packet
//	clientsMutex sync.RWMutex
//)

func main() {
	localIP := "127.0.0.1"
	localPort := "1540"
	wsServer := ws.New(localIP, localPort)
	wsServer.Start()
}
