package messages

import (
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterMessageRoutes(mux *http.ServeMux, msgService services.MessageService, authService services.AuthService) {
	mux.HandleFunc("/messages/", utils.ChainHTTP(MessageUserRouter(msgService), middlewares.RequireAuthMiddleware(authService)).ServeHTTP) // GET|POST /messages/:userId
	// TODO: Implement group + logic
	mux.HandleFunc("/messages/group/", MessageGroupRouter) // GET|POST /messages/group/:groupId
}
