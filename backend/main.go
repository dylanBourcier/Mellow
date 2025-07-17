package main

import (
	"log"
	"mellow/bootstrap"
	"mellow/config"
	"mellow/database"
	"mellow/middlewares"
	"mellow/routes"
	"mellow/utils"
	"net/http"
	"os"
	"time"
)

func main() {
	config.Load()
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/social.db" // Default value
	}

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath == "" {
		migrationsPath = "database/migration/sqlite" // Default value
	}
	err := database.ApplyMigrations(dbPath, migrationsPath)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ All migrations applied successfully.")

	// Initialiser la base de données
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Fatalf("❌ Failed to connect to the database: %v", err)
	}
	defer db.Close()
	repos := bootstrap.InitRepositories(db)
	services := bootstrap.InitServices(repos)
	mux := routes.SetupRoutes(services)
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads/"))))

	// Appliquer middlewares globaux
	handler := utils.ChainHTTP(mux,
		middlewares.LoggingHTTP, // Logging des requêtes
		middlewares.CORS,        // Headers CORS
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3225" // Default port
	}


	server := &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Println("✅ Server started at http://localhost:" + port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error: ", err)
	}
}
