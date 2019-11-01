package client

import (
	"net"

	"github.com/gorilla/websocket"
)

type Client struct {
	CID    string
	IPAddr string
}

type UDPClient struct {
	Client  *Client
	UDPAddr *net.UDPAddr
}
type WSClient struct {
	Client *Client
	Conn   *websocket.Conn
}
