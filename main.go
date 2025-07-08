package main

import (
	"log"
	"mellow/backend/utils"
)

func main() {
	dbPath := "backend/data/social.db"
	migrationsPath := "backend/database/migration/sqlite"

	err := utils.ApplyMigrations(dbPath, migrationsPath)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ All migrations applied successfully.")
}
