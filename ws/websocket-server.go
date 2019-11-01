package ws

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/fanmanpro/coordinator-server/coordinator"
	"github.com/fanmanpro/coordinator-server/worker"

	"github.com/google/uuid"

	"github.com/fanmanpro/coordinator-server/client"
	"github.com/fanmanpro/coordinator-server/gamedata"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/websocket"
)

type WSClient struct {
	client *client.Client
}

type gameServer struct {
	conn     *websocket.Conn
	id       string
	region   string
	seats    []*gamedata.ClientConnection
	capacity int32
}

var upgrader = websocket.Upgrader{}

//var clientPacketQueue []gamedata.Packet
var clientWebSocketMap map[string]*websocket.Conn
var clientMatchMap map[*websocket.Conn]string

var gameServerWebSocketMap map[string]*gameServer
var gameServerPacketQueue []gamedata.Packet

// WebSocketServer to open web socket ports
type WebSocketServer struct {
	coordinator *coordinator.Coordinator
	ip          string
	port        string
	workerPool  *worker.Pool
}

// NewServer initializes a new web socket server without starting it
func NewWebSocketServer(co *coordinator.Coordinator, ip string, port string, capacity int) *WebSocketServer {

	//clientPacketQueue = make([]gamedata.Packet, 0)
	clientWebSocketMap = make(map[string]*websocket.Conn)
	clientMatchMap = make(map[*websocket.Conn]string)
	ws := &WebSocketServer{co, ip, port, worker.NewPool(5, 100)}
	//for i := 0; i < cap(clients); i++ {
	//	go ws.processClient(<-clients)
	//}

	//gameServerPacketQueue = make([]gamedata.Packet, 0)
	//gameServerWebSocketMap = make(map[string]*gameServer)
	return ws
}

// Start starts the already intialized WebSocketServer
func (w *WebSocketServer) Start() {
	http.HandleFunc("/", w.connect)

	w.workerPool.Start()
	//go sending()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", w.ip, w.port), nil))
}

func (w *WebSocketServer) connect(wr http.ResponseWriter, rq *http.Request) {
	c, err := upgrader.Upgrade(wr, rq, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return

	}
	disconnected := make(chan bool, 1)

	disconnected <- false
	for {
		mt, data, err := c.ReadMessage()
		if err != nil {
			log.Printf("err: %v", err)
			disconnected <- true
		}
		if mt == websocket.BinaryMessage {
			w.workerPool.ScheduleJob(
				func() error {
					packet := &gamedata.Packet{}
					err = proto.Unmarshal(data, packet)
					if err != nil {
						panic(err)
					}
					log.Printf("recv: %s", packet.Header.OpCode)
					return w.handlePacket(c, packet)
				},
			)
		}
	}

	fmt.Println("end")
	<-disconnected
	c.Close()
	//clients <- c
	//disconnected := make(chan bool, 1)

	//go func() {
	//	disconnected <- w.processClient(c)
	//}()

	//<-disconnected
	fmt.Println("Client disconnected")
	//disconnected := make(chan bool, u.gameServer.Capacity())
	//simulation := make(chan *net.UDPAddr, 1)
	//receiving(c)
}
func (w *WebSocketServer) processClient(c *websocket.Conn) bool {
	c.SetCloseHandler(func(code int, text string) error {
		//return true
		return errors.New("asdas")
	})
	return true
}

func receiving(c *websocket.Conn) {
	defer func(c *websocket.Conn) {
		defer c.Close()
		fmt.Printf("%+v\n", c)
		if id, ok := clientMatchMap[c]; ok {
			delete(clientWebSocketMap, clientMatchMap[c])
			delete(clientMatchMap, c)
			log.Printf("disc: client: %v", id)
			return
		}
		// since its not a client it must be a game server. but PLEASE PLEASE don't do it this way
		// have a map for the conn of each game server and delete it like above, instead of the make
		gameServerWebSocketMap = make(map[string]*gameServer, 0)
		log.Printf("disc: gameserver")
	}(c)
}

func (w *WebSocketServer) send(p gamedata.Packet, c *websocket.Conn) error {
	data, err := proto.Marshal(&p)
	if err != nil {
		return err
	}

	err = c.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		return err
	} else {
		log.Printf("sent: %v", p.Header.OpCode)
		return nil
	}
}

//func sending() {
//	for {
//		time.Sleep(time.Millisecond)
//		for _, p := range clientPacketQueue {
//			if c, ok := clientWebSocketMap[p.Header.Cid]; ok {
//				send(p, c)
//				clientPacketQueue = nil
//			}
//		}
//
//		for _, p := range gameServerPacketQueue {
//			if c, ok := gameServerWebSocketMap[p.Header.Cid]; ok {
//				send(p, c.conn)
//				gameServerPacketQueue = nil
//			}
//		}
//	}
//}

func (w *WebSocketServer) handlePacket(c *websocket.Conn, p *gamedata.Packet) error {
	switch p.Header.OpCode {
	case gamedata.Header_GameServerOnline:
		{
			gameServerJoined := &gamedata.GameServerOnline{}
			// a reminder that the ClientJoined message comming from the client's first connection doesn't have a Data object of type Any
			err := ptypes.UnmarshalAny(p.Data, gameServerJoined)
			if err != nil {
				//log.Printf("err: invalid %v data. err: %v", gamedata.Header_GameServerOnline, err)
				return err
			}

			if gameServerJoined.Secret == "fanmanpro" {
				id, err := uuid.NewUUID()
				if err != nil {
					return err
				}
				packet := gamedata.Packet{
					Header: &gamedata.Header{
						OpCode: gamedata.Header_GameServerOnline,
						Cid:    id.String(),
					},
				}

				gameServerWebSocketMap[id.String()] = &gameServer{id: id.String(), conn: c, region: gameServerJoined.Region, capacity: gameServerJoined.Capacity, seats: make([]*gamedata.ClientConnection, 0)}
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
				return err
			}

			id, err := uuid.NewUUID()
			if err != nil {
				return err
			}
			packet := gamedata.Packet{
				Header: &gamedata.Header{
					OpCode: gamedata.Header_ClientOnline,
					Cid:    id.String(),
				},
				Data: data,
			}

			clientWebSocketMap[id.String()] = c

			return w.workerPool.ScheduleJob(func() error { return w.send(packet, c) })
			//clientPacketQueue = append(clientPacketQueue, packet)
		}
		break
	case gamedata.Header_ClientGameRequest:
		{
			if _, ok := clientWebSocketMap[p.Header.Cid]; ok {
				clientGameRequest := &gamedata.ClientGameRequest{}
				// a reminder that the ClientJoined message comming from the client's first connection doesn't have a Data object of type Any
				err := ptypes.UnmarshalAny(p.Data, clientGameRequest)
				if err != nil {
					return err
				}
				for _, g := range gameServerWebSocketMap {
					if g.region == clientGameRequest.Region {
						g.seats = append(g.seats, &gamedata.ClientConnection{ID: p.Header.Cid, Address: c.RemoteAddr().String()})
						if len(g.seats) == int(g.capacity) {
							gameServerStart := &gamedata.GameServerStart{}
							gameServerStart.Clients = g.seats

							data, err := ptypes.MarshalAny(gameServerStart)
							if err != nil {
								return err
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
					return err
				}

				for _, cc := range gameServerStart.Clients {
					if c, ok := clientWebSocketMap[cc.ID]; ok {
						clientMatchMap[c] = gameServerStart.ID

						clientGameFound := &gamedata.ClientGameFound{}
						clientGameFound.ID = gameServerStart.ID

						data, err := ptypes.MarshalAny(clientGameFound)
						if err != nil {
							return err
						}
						packet := gamedata.Packet{
							Header: &gamedata.Header{
								OpCode: gamedata.Header_ClientGameFound,
								Cid:    cc.ID,
							},
							Data: data,
						}
						//errchan := make(chan error, 1)
						//w.workerPool.ScheduleJob(worker.NewJob(func() error {
						//	errchan <- w.send(packet, c)
						//	return nil
						//}))
						return w.workerPool.ScheduleJob(func() error { return w.send(packet, c) })
						//return <-errchan
						//clientPacketQueue = append(clientPacketQueue, packet)
					}
				}
			}
		}
		break
	}
	return errors.New(fmt.Sprintf("Packet received but unknown handler for %v", p.Header.OpCode))
}
