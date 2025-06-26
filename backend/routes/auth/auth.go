package auth

import (
	"mellow/controllers/auth"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/auth/signup", auth.SignUpHandler)
	mux.HandleFunc("/auth/login", auth.LoginHandler)
	mux.HandleFunc("/auth/logout", auth.LogoutHandler)
}
