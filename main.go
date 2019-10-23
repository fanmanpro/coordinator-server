package main

import "github.com/fanmanpro/coordinator-server/ws"

//var (
//	clients      map[int32]*Client
//	packetQueue  []Packet
//	clientsMutex sync.RWMutex
//)

func main() {
	ip := "127.0.0.1"
	port := "6000"
	wsServer := ws.New(ip, port)
	wsServer.Start()
}
