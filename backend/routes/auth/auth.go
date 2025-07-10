package auth

import (
	"mellow/controllers/auth"
	"mellow/services"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux, userService services.UserService, authService services.AuthService) {
	mux.HandleFunc("/auth/signup", auth.SignUpHandler(userService))
	mux.HandleFunc("/auth/login", auth.LoginHandler(authService))
	mux.HandleFunc("/auth/logout", auth.LogoutHandler(authService))
}
