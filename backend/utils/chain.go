package utils

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Fonction permettant de chainer plusieurs middleware
func Chain(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// Middleware pour http.Handler
type HTTPMiddleware func(http.Handler) http.Handler

func ChainHTTP(h http.Handler, middlewares ...HTTPMiddleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
