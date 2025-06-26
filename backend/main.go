package main

import (
	"log"
	"mellow/middlewares"
	"net/http"
	"time"

	"mellow/routes"
	"mellow/utils"
)

func main() {
	mux := routes.SetupRoutes()

	// Appliquer middlewares globaux
	handler := utils.ChainHTTP(mux,
		middlewares.LoggingHTTP, // Logging des requêtes
		middlewares.CORS,        // Headers CORS
	)

	server := &http.Server{
		Addr:              ":3225",
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Println("✅ Serveur lancé sur http://localhost:3225")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Erreur serveur : ", err)
	}
}
