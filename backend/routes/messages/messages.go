package messages

import (
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterMessageRoutes(mux *http.ServeMux, msgService services.MessageService, authService services.AuthService) {
<<<<<<< HEAD
	mux.HandleFunc("/messages/", utils.ChainHTTP(MessageUserRouter(msgService), middlewares.RequireAuthMiddleware(authService)).ServeHTTP) // GET|POST /messages/:userId
=======
	mux.HandleFunc("/messages/", utils.ChainHTTP(MessageUserRouter(msgService), middlewares.RequireAuthMiddleware(authService))) // GET|POST /messages/:userId
>>>>>>> dece5e0 (added necessary message route)
	// TODO: Implement group + logic
	mux.HandleFunc("/messages/group/", MessageGroupRouter) // GET|POST /messages/group/:groupId
}
