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

var clientWebSocketMap map[string]*websocket.Conn

var gameServerWebSocketMap map[string]*gameServer
var gameServerPacketQueue []gamedata.Packet

// WebSocketServer to open web socket ports
type WebSocketServer struct {
	coordinator *coordinator.Coordinator
	ip          string
	port        string
	workerPool  *worker.Pool
}

// NewWebSocketServer initializes a new web socket server without starting it
func NewWebSocketServer(co *coordinator.Coordinator, ip string, port string) *WebSocketServer {
	clientWebSocketMap = make(map[string]*websocket.Conn)
	ws := &WebSocketServer{co, ip, port, worker.NewPool(5, 100)}
	gameServerWebSocketMap = make(map[string]*gameServer)
	return ws
}

// Start starts the already intialized WebSocketServer
func (w *WebSocketServer) Start() {
	http.HandleFunc("/", w.processClient)

	w.workerPool.Start()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", w.ip, w.port), nil))
}

func (w *WebSocketServer) processClient(wr http.ResponseWriter, rq *http.Request) {
	c, err := upgrader.Upgrade(wr, rq, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	disconnected := make(chan bool, 1)

	c.SetCloseHandler(func(code int, text string) error {
		for k, v := range gameServerWebSocketMap {
			if v.conn == c {
				log.Println(fmt.Sprintf("game server over connection %v disconnected", v.id))
				delete(gameServerWebSocketMap, k)
				log.Println(fmt.Sprintf("%v game servers online", len(gameServerWebSocketMap)))
			}
		}
		for k, v := range clientWebSocketMap {
			if v == c {
				delete(clientWebSocketMap, k)
				log.Println(fmt.Sprintf("%v clients online", len(clientWebSocketMap)))
			}
		}
		disconnected <- true
		return nil
	})

	defer c.Close()

	go func() {
		for {
			log.Printf("waiting for packet")
			mt, data, err := c.ReadMessage()
			if err != nil {
				_, ok := err.(*websocket.CloseError)
				if !ok {
					log.Printf("err: %v", mt)
				}
				break
			}
			if mt == websocket.BinaryMessage {
				w.workerPool.ScheduleJob(
					func(errChan chan error) {
						packet := &gamedata.Packet{}
						err = proto.Unmarshal(data, packet)
						if err != nil {
							panic(err)
						}
						log.Printf("receiving packet %v from %v", packet.Header.OpCode, packet.Header.Cid)
						log.Printf("recv: %s", packet.Header.OpCode)
						errChan <- w.handlePacket(c, packet)
					},
				)
			}
		}
	}()
	<-disconnected
}

func (w *WebSocketServer) send(p gamedata.Packet, c *websocket.Conn) error {
	log.Printf("sending packet %v over connection %v to %v from %v", p.Header.OpCode, p.Header.Cid, c.RemoteAddr().String(), c.LocalAddr().String())
	data, err := proto.Marshal(&p)
	if err != nil {
		return err
	}

	err = c.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		return err
	}
	log.Printf("sent: %v", p.Header.OpCode)
	return nil
}

func (w *WebSocketServer) handlePacket(c *websocket.Conn, p *gamedata.Packet) error {
	switch p.Header.OpCode {
	case gamedata.Header_GameServerOnline:
		{
			gameServerJoined := &gamedata.GameServerOnline{}
			// a reminder that the ClientJoined message comming from the client's first connection doesn't have a Data object of type Any
			err := ptypes.UnmarshalAny(p.Data, gameServerJoined)
			if err != nil {
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
				//fmt.Printf("%+v\n", &gameServer{id: id.String(), conn: c, region: gameServerJoined.Region, capacity: gameServerJoined.Capacity, seats: make([]*gamedata.ClientConnection, 0)})
				w.workerPool.ScheduleJob(func(errChan chan error) { errChan <- w.send(packet, c) })
				return nil
			}
			return errors.New(fmt.Sprintf("Invalid game server"))
		}
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

			w.workerPool.ScheduleJob(func(errChan chan error) { errChan <- w.send(packet, c) })
			return nil
		}
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
							//gameServerPacketQueue = append(gameServerPacketQueue, packet)
							w.workerPool.ScheduleJob(func(errChan chan error) { errChan <- w.send(packet, g.conn) })
							return nil
						}
					}
				}
			}
			return errors.New(fmt.Sprintf("No game servers available for client to join"))
		}
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
						w.workerPool.ScheduleJob(func(errChan chan error) { errChan <- w.send(packet, c) })
						return nil
					}
				}
			}
			return errors.New(fmt.Sprintf("Game server not available anymore"))
		}
	default:
		{
			return errors.New(fmt.Sprintf("Packet received but unknown handler for %v", p.Header.OpCode))
		}
	}
}
