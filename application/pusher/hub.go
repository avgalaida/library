package main

import (
	"encoding/json"
	"github.com/avgalaida/library/domain"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients    []*Client
	nextID     int
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

func newHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		nextID:     0,
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (hub *Hub) run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) broadcast(message domain.IDelta) {
	data, _ := json.Marshal(message)
	typeString := []byte(`,"type":"` + message.Key() + `"}`)
	data = data[:len(data)-1]
	data = append(data, typeString...)

	for _, c := range hub.clients {
		c.outbound <- data
	}
}

func (hub *Hub) send(message interface{}, client *Client) {
	data, _ := json.Marshal(message)
	client.outbound <- data
}

func (hub *Hub) disconnect(client *Client) {
	log.Println("отключен клиент: ", client.socket.RemoteAddr())
	client.close()
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	i := -1
	for j, c := range hub.clients {
		if c.id == client.id {
			i = j
			break
		}
	}

	copy(hub.clients[i:], hub.clients[i+1:])
	hub.clients[len(hub.clients)-1] = nil
	hub.clients = hub.clients[:len(hub.clients)-1]
}

func (hub *Hub) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, _ := upgrader.Upgrade(w, r, nil)
	client := newClient(hub, socket)
	hub.register <- client
	go client.write()
}

func (hub *Hub) onConnect(client *Client) {
	log.Println("подлючился клиент: ", client.socket.RemoteAddr())

	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	client.id = hub.nextID
	hub.nextID++
	hub.clients = append(hub.clients, client)
}

func (hub *Hub) onDisconnect(client *Client) {
	log.Println("отключился клиент: ", client.socket.RemoteAddr())
	hub.disconnect(client)
}
