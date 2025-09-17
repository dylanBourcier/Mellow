package messages

import (
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterMessageRoutes(mux *http.ServeMux, msgService services.MessageService, userSvc services.UserService, authService services.AuthService) {
	// GET|POST /messages
	mux.HandleFunc("/messages", utils.ChainHTTP(MessageRouter(msgService), middlewares.RequireAuthMiddleware(authService)).ServeHTTP)
	// GET|POST /messages/:userId
	mux.HandleFunc("/messages/", utils.ChainHTTP(MessageUserRouter(msgService, userSvc), middlewares.RequireAuthMiddleware(authService)).ServeHTTP)
	// TODO: Implement group + logic
	// GET|POST /messages/group/:groupId
	mux.HandleFunc("/messages/group/", utils.ChainHTTP(MessageGroupRouter(msgService, userSvc), middlewares.RequireAuthMiddleware(authService)).ServeHTTP)
}
