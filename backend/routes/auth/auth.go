package auth

import (
	"mellow/controllers/auth"
	"mellow/services"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux, userService services.UserService) {
	mux.HandleFunc("/auth/signup", auth.SignUpHandler(userService))
	mux.HandleFunc("/auth/login", auth.LoginHandler)
	mux.HandleFunc("/auth/logout", auth.LogoutHandler)
}
