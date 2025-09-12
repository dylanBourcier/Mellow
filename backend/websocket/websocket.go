package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// ⚠️ sécuriser : vérifier domaine ou cookie de session
		return true
	},
}

type Client struct {
	conn *websocket.Conn
	room string
	send chan []byte
}

var rooms = make(map[string]map[*Client]bool) // roomID -> clients

// Structure du message envoyé aux clients
type WSMessage struct {
	ID             string  `json:"message_id"`
	SenderID       string  `json:"sender_id"`
	SenderUsername *string `json:"username,omitempty"`
	SenderImageUrl *string `json:"image_url,omitempty"`
	Content        string  `json:"content"`
	Timestamp      string  `json:"creation_date"`
	Room           string  `json:"room"`
	Type           string  `json:"type"` // "private" ou "group"
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	room := r.URL.Query().Get("room")
	if room == "" {
		log.Println("missing room id")
		conn.Close()
		return
	}

	client := &Client{
		conn: conn,
		room: room,
		send: make(chan []byte, 256),
	}

	// Ajouter le client dans la room
	if rooms[room] == nil {
		rooms[room] = make(map[*Client]bool)
	}
	rooms[room][client] = true

	go client.writePump()
	client.readPump()
}

func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
		delete(rooms[c.room], c)
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		// Ici on n’insère pas en DB (c’est REST qui gère la persistance)
		// Mais tu pourrais log ou gérer si tu voulais.
		log.Printf("Message reçu mais ignoré (REST fait la persistance): %s", msg)
	}
}

func (c *Client) writePump() {
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}

// Diffuser un message à une room (appelé depuis ton handler REST après insertion DB)
func BroadcastMessage(room string, message WSMessage) {
	fmt.Println("Broadcasting to room:", room, "message:", message)
	data, _ := json.Marshal(message)
	for client := range rooms[room] {
		select {
		case client.send <- data:
		default:
			close(client.send)
			delete(rooms[room], client)
		}
	}
}

func ListRooms() []string {
	keys := make([]string, 0, len(rooms))
	for room := range rooms {
		keys = append(keys, room)
	}
	return keys
}
func ListClients(room string) []string {
	var clientsList []string
	for client := range rooms[room] {
		// Ici tu peux mettre l’ID du sender si tu l’as stocké dans Client
		clientsList = append(clientsList, client.conn.RemoteAddr().String())
	}
	return clientsList
}
func RegisterDebugRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ws/debug", func(w http.ResponseWriter, r *http.Request) {
		type RoomInfo struct {
			Room    string   `json:"room"`
			Clients []string `json:"clients"`
		}
		var infos []RoomInfo
		for room := range rooms {
			infos = append(infos, RoomInfo{
				Room:    room,
				Clients: ListClients(room),
			})
		}
		json.NewEncoder(w).Encode(infos)
	})
}

func MakePrivateRoom(user1, user2 string) string {
	if user1 < user2 {
		return "private:" + user1 + ":" + user2
	}
	return "private:" + user2 + ":" + user1
}
