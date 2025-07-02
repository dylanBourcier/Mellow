package main

import (
	"log"
	"mellow/backend/database"
	"mellow/backend/middlewares"
	"mellow/backend/routes"
	"mellow/backend/utils"
	"net/http"
	"time"
)

func main() {
	dbPath := "backend/data/social.db"
	migrationsPath := "backend/database/migration/sqlite"

	err := database.ApplyMigrations(dbPath, migrationsPath)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ All migrations applied successfully.")

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
