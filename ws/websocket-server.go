package ws

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/fanmanpro/coordinator-server/gamedata"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/websocket"
)

type gameServer struct {
	conn     *websocket.Conn
	id       string
	region   string
	seats    []string
	capacity int32
}

type client struct {
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{}

var clientPacketQueue []gamedata.Packet
var clientWebSocketMap map[string]*client
var clientMatchMap map[*client]string

var gameServerWebSocketMap map[string]*gameServer
var gameServerPacketQueue []gamedata.Packet

// WebSocketServer to open web socket ports
type WebSocketServer struct {
	ip   string
	port string
}

// New initializes a new web socket server without starting it
func New(ip string, port string) *WebSocketServer {
	clientPacketQueue = make([]gamedata.Packet, 0)
	clientWebSocketMap = make(map[string]*client)
	clientMatchMap = make(map[*client]string)

	gameServerPacketQueue = make([]gamedata.Packet, 0)
	gameServerWebSocketMap = make(map[string]*gameServer)
	return &WebSocketServer{ip, port}
}

// Start starts the already intialized WebSocketServer
func (w *WebSocketServer) Start() {
	http.HandleFunc("/", echo)
	go sending()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", w.ip, w.port), nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	receiving(c)
}

func receiving(c *websocket.Conn) {
	for {
		time.Sleep(time.Millisecond)
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("err: %v", err)
			break
		}
		if mt == websocket.BinaryMessage {
			packet := &gamedata.Packet{}
			err = proto.Unmarshal(message, packet)
			if err != nil {
				panic(err)
			}
			log.Printf("recv: %s", packet.Header.OpCode)
			handlePacket(c, packet)
		}
	}
}

func send(p gamedata.Packet, c *websocket.Conn) {
	data, err := proto.Marshal(&p)
	if err != nil {
		log.Printf("err: failed to marshal packet. %v", err)
		return
	}

	err = c.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		log.Println("err:", err)
	} else {
		log.Printf("sent: %v", p.Header.OpCode)
	}
}

func sending() {
	for {
		time.Sleep(time.Millisecond)
		for _, p := range clientPacketQueue {
			if c, ok := clientWebSocketMap[p.Header.Cid]; ok {
				send(p, c.conn)
				clientPacketQueue = nil
			}
		}

		for _, p := range gameServerPacketQueue {
			if c, ok := gameServerWebSocketMap[p.Header.Cid]; ok {
				send(p, c.conn)
				fmt.Println(p)
				gameServerPacketQueue = nil
			}
		}
	}
}

func handlePacket(c *websocket.Conn, p *gamedata.Packet) {
	switch p.Header.OpCode {
	case gamedata.Header_GameServerOnline:
		{
			gameServerJoined := &gamedata.GameServerOnline{}
			// a reminder that the ClientJoined message comming from the client's first connection doesn't have a Data object of type Any
			err := ptypes.UnmarshalAny(p.Data, gameServerJoined)
			if err != nil {
				log.Printf("err: invalid %v data. err: %v", gamedata.Header_GameServerOnline, err)
				break
			}

			if gameServerJoined.Secret == "fanmanpro" {
				id, err := uuid.NewUUID()
				if err != nil {
					log.Printf("err: could not generate uuid. %v", err)
					break
				}
				packet := gamedata.Packet{
					Header: &gamedata.Header{
						OpCode: gamedata.Header_GameServerOnline,
						Cid:    id.String(),
					},
				}

				gameServerWebSocketMap[id.String()] = &gameServer{id: id.String(), conn: c, region: gameServerJoined.Region, capacity: gameServerJoined.Capacity, seats: make([]string, 0)}
				gameServerPacketQueue = append(gameServerPacketQueue, packet)
			}
		}
		break
	case gamedata.Header_ClientOnline:
		{
			clientOnline := &gamedata.ClientOnline{}
			// get the player info from the database
			clientOnline.Name = "FanManPro"

			data, err := ptypes.MarshalAny(clientOnline)
			if err != nil {
				log.Printf("err: could not generate uuid. %v", err)
				break
			}

			id, err := uuid.NewUUID()
			if err != nil {
				log.Printf("err: could not generate uuid. %v", err)
				break
			}
			packet := gamedata.Packet{
				Header: &gamedata.Header{
					OpCode: gamedata.Header_ClientOnline,
					Cid:    id.String(),
				},
				Data: data,
			}

			clientWebSocketMap[id.String()] = &client{conn: c}
			clientPacketQueue = append(clientPacketQueue, packet)
		}
		break
	case gamedata.Header_ClientGameRequest:
		{
			if _, ok := clientWebSocketMap[p.Header.Cid]; ok {
				clientGameRequest := &gamedata.ClientGameRequest{}
				// a reminder that the ClientJoined message comming from the client's first connection doesn't have a Data object of type Any
				err := ptypes.UnmarshalAny(p.Data, clientGameRequest)
				if err != nil {
					log.Printf("err: invalid %v data. err: %v", gamedata.Header_GameServerOnline, err)
					break
				}
				for _, g := range gameServerWebSocketMap {
					if g.region == clientGameRequest.Region {
						g.seats = append(g.seats, p.Header.Cid)
						if len(g.seats) == int(g.capacity) {
							gameServerStart := &gamedata.GameServerStart{}
							gameServerStart.Clients = g.seats

							data, err := ptypes.MarshalAny(gameServerStart)
							if err != nil {
								log.Printf("err: could not generate uuid. %v", err)
								break
							}
							packet := gamedata.Packet{
								Header: &gamedata.Header{
									OpCode: gamedata.Header_GameServerStart,
									Cid:    g.id,
								},
								Data: data,
							}
							gameServerPacketQueue = append(gameServerPacketQueue, packet)
						}
					}
				}

				//if gameServerJoined.Secret == "fanmanpro" {
				//id, err := uuid.NewUUID()
				//if err != nil {
				//	log.Printf("err: could not generate uuid. %v", err)
				//	break
				//}
				//packet := gamedata.Packet{
				//	Header: &gamedata.Header{
				//		OpCode: opCode,
				//		Cid:    id.String(),
				//	},
				//}

				//gameServerWebSocketMap[id.String()] = c
				//gameServerPacketQueue = append(gameServerPacketQueue, packet)
				//}
			}
		}
		break
	case gamedata.Header_GameServerStart:
		{
			if _, ok := gameServerWebSocketMap[p.Header.Cid]; ok {
				gameServerStart := &gamedata.GameServerStart{}
				// a reminder that the ClientJoined message comming from the client's first connection doesn't have a Data object of type Any
				err := ptypes.UnmarshalAny(p.Data, gameServerStart)
				if err != nil {
					log.Printf("err: invalid %v data. err: %v", gamedata.Header_GameServerOnline, err)
					break
				}

				for _, cid := range gameServerStart.Clients {
					if c, ok := clientWebSocketMap[cid]; ok {
						clientMatchMap[c] = gameServerStart.ID

						clientGameFound := &gamedata.ClientGameFound{}
						clientGameFound.ID = gameServerStart.ID

						data, err := ptypes.MarshalAny(clientGameFound)
						if err != nil {
							log.Printf("err: could not generate uuid. %v", err)
							break
						}
						packet := gamedata.Packet{
							Header: &gamedata.Header{
								OpCode: gamedata.Header_ClientGameFound,
								Cid:    cid,
							},
							Data: data,
						}
						clientPacketQueue = append(clientPacketQueue, packet)
					}
				}
			}
		}
		break
	}
}
