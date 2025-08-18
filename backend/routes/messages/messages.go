package messages

import (
	"mellow/services"
	"net/http"
)

func RegisterMessageRoutes(mux *http.ServeMux, messageService services.MessageService, authService services.AuthService) {
	mux.Handle("/messages/", MessageUserRouter(messageService, authService)) // GET|POST /messages/:userId

	mux.HandleFunc("/messages/group/", MessageGroupRouter) // GET|POST /messages/group/:groupId
}
