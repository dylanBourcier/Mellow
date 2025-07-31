package websocket

import ()

// Hub structure centrale qui gère tous les clients connectés et la diffusion des messages.
type Hub struct {
	UsersConnected map[uuid.UUID]*models.User
	Mutex sync.Mutex
}

// NewHub initialise et retourne une nouvelle instance de Hub.
// À utiliser pour créer le gestionnaire principal des connexions websocket.
func NewHub() *Hub {
	return &Hub{
		UsersConnected: make(map[uuid.UUID]*models.User),
	}
}
// AddClient ajoute un nouveau client connecté au Hub.
// À appeler lorsqu'un utilisateur ouvre une connexion websocket.
func (hub *Hub) AddClient() {
hub.Mutex.Lock()
	defer hub.Mutex.Unlock()
}

// RemoveClient retire un client du Hub.
// À appeler lorsqu'un utilisateur ferme sa connexion websocket ou se déconnecte.
func (hub *Hub) RemoveClient() {
	hub.Mutex.Lock()
	defer hub.Mutex.Unlock()
}

// Broadcast envoie un message à tous les clients connectés au Hub.
// À utiliser pour diffuser des messages à tous les utilisateurs en temps réel.
func (hub *Hub) Broadcast() {
	hub.Mutex.Lock()
	defer hub.Mutex.Unlock()
}

