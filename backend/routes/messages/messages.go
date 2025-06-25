package messages

import (
	"net/http"
)

func RegisterMessageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/messages/", MessageUserRouter) // GET|POST /messages/:userId

	mux.HandleFunc("/messages/group/", MessageGroupRouter) // GET|POST /messages/group/:groupId
}
