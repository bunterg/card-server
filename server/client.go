package server

import "github.com/gorilla/websocket"

type Client struct {
	hub *Hub
	// The websocket connection.
	conn *websocket.Conn

	room string
	name string
	id   int

	// Buffered channel of outbound messages.
	send chan []byte
}
