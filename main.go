package main

import (
	"log"
	"mellow/backend/database"
)

func main() {
	dbPath := "backend/data/social.db"
	migrationsPath := "backend/database/migration/sqlite"

	err := database.ApplyMigrations(dbPath, migrationsPath)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ All migrations applied successfully.")
}
