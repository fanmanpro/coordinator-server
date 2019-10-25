package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fanmanpro/coordinator-server/gamedata"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// WebSocketServer to open web socket ports
type WebSocketServer struct {
	ip   string
	port string
}

// New initializes a new web socket server without starting it
func New(ip string, port string) *WebSocketServer {
	return &WebSocketServer{ip, port}
}

// Start starts the already intialized WebSocketServer
func (w *WebSocketServer) Start() {
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", w.ip, w.port), nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	go receiving(c)
	sending(c)
}
func receiving(c *websocket.Conn) {
	for {
		mt, message, err := c.ReadMessage()
		log.Println("mt: ", mt)
		if err != nil {
			log.Println("read: ", err)
			break
		}
		if mt == websocket.BinaryMessage {
			newMessage := &gamedata.Packet{}
			err = proto.Unmarshal(message, newMessage)
			if err != nil {
				panic(err)
			}
			log.Printf("recv: %s", newMessage.Header.OpCode)
		}
	}
}
func sending(c *websocket.Conn) {
	for {
		err := c.WriteMessage(websocket.BinaryMessage, []byte{})
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
