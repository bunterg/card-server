package server

type message struct {
	data []byte
	room string
}

// Hub server control
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Registered connections.
	rooms map[string]map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Inbound messages from the clients.
	broadcastRoom chan message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// NewHub creates a Hub instance
func NewHub() *Hub {
	return &Hub{
		broadcast:     make(chan []byte),
		broadcastRoom: make(chan message),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		rooms:         make(map[string]map[*Client]bool),
	}
}

// Run runs Hub goroutines
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.broadcastRoom:
			for client := range h.rooms[message.room] {
				select {
				case client.send <- message.data:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
