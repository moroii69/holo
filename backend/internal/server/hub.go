package server

import (
	"sync"
	"time"
)

// Hub manages rooms and garbage collection.
type Hub struct {
	rooms       map[string]*Room
	mu          sync.RWMutex
	register    chan *Client
	unregister  chan *Client
	inactivity  time.Duration
	gcTickerDur time.Duration
}

func NewHub(inactivity time.Duration) *Hub {
	return &Hub{
		rooms:       make(map[string]*Room),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		inactivity:  inactivity,
		gcTickerDur: time.Minute,
	}
}

func (h *Hub) Run() {
	ticker := time.NewTicker(h.gcTickerDur)
	defer ticker.Stop()

	for {
		select {
		case c := <-h.register:
			h.addClient(c)
		case c := <-h.unregister:
			h.removeClient(c)
		case <-ticker.C:
			h.gcRooms()
		}
	}
}

func (h *Hub) addClient(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	room, ok := h.rooms[c.roomID]
	if !ok {
		room = NewRoom(c.roomID)
		h.rooms[c.roomID] = room
	}
	room.AddClient(c)
	logInfo("client_joined", logFields{
		"roomId":      c.roomID,
		"clientId":    c.id,
		"clientCount": len(room.clients),
	})
}

func (h *Hub) removeClient(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	room, ok := h.rooms[c.roomID]
	if !ok {
		return
	}
	room.RemoveClient(c)
	logInfo("client_left", logFields{
		"roomId":      c.roomID,
		"clientId":    c.id,
		"clientCount": len(room.clients),
	})
	if len(room.clients) == 0 {
		delete(h.rooms, c.roomID)
	}
}

func (h *Hub) Broadcast(roomID string, sender *Client, message []byte) {
	h.mu.RLock()
	room, ok := h.rooms[roomID]
	h.mu.RUnlock()
	if !ok {
		return
	}
	room.Broadcast(sender, message)
}

func (h *Hub) gcRooms() {
	now := time.Now()
	h.mu.Lock()
	defer h.mu.Unlock()

	for id, room := range h.rooms {
		if now.Sub(room.lastActivity) > h.inactivity || len(room.clients) == 0 {
			logInfo("room_gc", logFields{
				"roomId":      id,
				"clientCount": len(room.clients),
				"lastActive":  room.lastActivity,
			})
			for c := range room.clients {
				close(c.send)
				_ = c.conn.Close()
			}
			delete(h.rooms, id)
		}
	}
}
