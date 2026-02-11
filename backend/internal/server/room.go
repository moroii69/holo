package server

import (
	"sync"
	"time"
)

// Room is an in-memory collection of clients.
type Room struct {
	id           string
	clients      map[*Client]struct{}
	mu           sync.RWMutex
	lastActivity time.Time
}

func NewRoom(id string) *Room {
	return &Room{
		id:           id,
		clients:      make(map[*Client]struct{}),
		lastActivity: time.Now(),
	}
}

func (r *Room) touch() {
	r.lastActivity = time.Now()
}

func (r *Room) AddClient(c *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.clients[c] = struct{}{}
	r.touch()
}

func (r *Room) RemoveClient(c *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.clients, c)
	r.touch()
}

func (r *Room) Broadcast(sender *Client, msg []byte) {
	// Take a full lock while walking and potentially pruning the client set.
	r.mu.Lock()
	defer r.mu.Unlock()

	r.touch()
	for c := range r.clients {
		if c == sender {
			continue
		}
		select {
		case c.send <- msg:
		default:
			// slow consumer – drop connection
			close(c.send)
			_ = c.conn.Close()
			delete(r.clients, c)
		}
	}
}
