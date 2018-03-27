package register

import (
	"github.com/satori/go.uuid"
	"github.com/gorilla/websocket"
	"encoding/json"
	"message"
)

type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

type Client struct {
	Uid    string
	MUid   int
	Socket *websocket.Conn
	Send   chan []byte
}

func GetUid() string {
	sUuid, _ := uuid.NewV4();
	cUuid := uuid.NewV5(sUuid, "fdkljklfd").String()
	return cUuid
}

func SetUserInfo(conn *websocket.Conn) *Client {
	client := &Client{Uid: GetUid(), MUid: 0, Socket: conn, Send: make(chan []byte)}
	return client
}
func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&message.Message{Content: "/A new socket has connected."})
<<<<<<< HEAD
			manager.send(jsonMessage, conn)
			manager.Send()
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
=======
			message.Send(jsonMessage, conn, manager)
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&message.Message{Content: "/A socket has disconnected."})
				message.Send(jsonMessage, conn, manager)
>>>>>>> 44e1face9d923888421f132172ce11ce8365e7b0
			}
		case messageContent := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- messageContent:
				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}
