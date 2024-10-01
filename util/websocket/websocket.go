package websocketutil

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Message struct {
	VideoTitle  string `json:"video_title"`
	EmailSender string `json:"email_sender"`
}

type WebSocket struct {
	Broadcast chan Message
	Clients   map[*websocket.Conn]bool
	Upgrader  *websocket.Upgrader
}

func New() *WebSocket {
	return &WebSocket{
		Broadcast: make(chan Message),
		Clients:   make(map[*websocket.Conn]bool),
		Upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (s *WebSocket) HandleConnection(c echo.Context) error {
	ws, err := s.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	s.Clients[ws] = true
	return nil
}

func (s *WebSocket) HandleMessage() {
	defer func() {
		for client := range s.Clients {
			client.Close()
		}
		close(s.Broadcast)
	}()

	for {
		msg := <-s.Broadcast
		fmt.Println("Broadcasting message: ", msg)
		for client := range s.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(s.Clients, client)
			}
			fmt.Println("Message sent for client: ", client.RemoteAddr())
		}
	}
}

func (s *WebSocket) BroadcastMessage(msg Message) {
	s.Broadcast <- msg
}
