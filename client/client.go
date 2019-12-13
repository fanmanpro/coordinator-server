package client

import (
	"fmt"
	"net"

	"github.com/fanmanpro/coordinator-server/gamedata"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Client TODO
type Client struct {
	CID     string
	Addr    net.IP
	TCPConn TCPConnection
	UDPConn UDPConnection
	WSConn  WSConnection
}

// New TODO
func New() (*Client, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("err: could not generate uuid. %v", err)
	}
	return &Client{CID: id.String()}, nil
}

// NewTest TODO
func NewTest() (*Client, error) {
	return &Client{CID: "abc"}, nil
}

//func (c *Client) InitializeConnections(conn *net.TCPConn, u *net.UDPConn, s chan *gamedata.Packet) error {
//	c.Addr = conn.RemoteAddr()
//	c.TCPConn.Conn = conn
//	c.TCPConn.Send = s
//	c.UDPConn.Conn = u
//	c.UDPConn.Send = s
//	//u.
//	return nil
//}

// TCPConnection TODO
type TCPConnection struct {
	// TCP clients have active connections, the server uses this connection to send data with
	Conn *net.TCPConn
	Send chan *gamedata.Packet
}

// UDPConnection TODO
type UDPConnection struct {
	// UDP clients don't have active connections, the server just uses their addresses to send data with
	Addr *net.UDPAddr
	Send chan *gamedata.Packet
}

// WSConnection TODO
type WSConnection struct {
	Conn *websocket.Conn
	Send chan *gamedata.Packet
}

// AcceptPacket TODO
type AcceptTCPPacket struct {
	CID  string
	Conn *net.TCPConn
	Send chan *gamedata.Packet
}

type AcceptUDPPacket struct {
	CID   string
	Addr  *net.UDPAddr
	Send  chan *gamedata.Packet
	Start chan bool
}
