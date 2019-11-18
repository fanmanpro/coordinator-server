package client

import (
	"net"

	"github.com/fanmanpro/coordinator-server/gamedata"
	"github.com/gorilla/websocket"
)

type Client struct {
	CID    string
	IPAddr string
}

type UDPClient struct {
	Client  *Client
	UDPAddr *net.UDPAddr
	Send    chan gamedata.Packet
}
type WSClient struct {
	Client *Client
	Conn   *websocket.Conn
}
