package messages

import (
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterMessageRoutes(mux *http.ServeMux, msgService services.MessageService, authService services.AuthService) {
	// GET|POST /messages/:userId
	mux.HandleFunc("/messages/", utils.ChainHTTP(MessageUserRouter(msgService), middlewares.RequireAuthMiddleware(authService)).ServeHTTP)
	// TODO: Implement group + logic
	// GET|POST /messages/group/:groupId
	mux.HandleFunc("/messages/group/", MessageGroupRouter)
}
