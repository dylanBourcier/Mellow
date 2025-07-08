package database

import (
	"database/sql"
	"fmt"
	"io/fs"
	"mellow/repositories"
	"mellow/repositories/repoimpl"
	"mellow/services"
	"mellow/services/servimpl"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func ApplyMigrations(dbPath string, migrationsPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	var files []string
	err = filepath.WalkDir(migrationsPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".up.sql") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error reading migration files: %w", err)
	}

	sort.Strings(files)

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", file, err)
		}
		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("error executing migration %s: %w", file, err)
		}
		fmt.Printf("✅ Applied migration: %s\n", file)
	}

	return nil
}

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Check if the database is reachable
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	fmt.Println("✅ Database connection established successfully.")
	return db, nil
}

type Repositories struct {
	UserRepository repositories.UserRepository
}

func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: repoimpl.NewUserRepository(db),
	}
}

type Services struct {
	UserService services.UserService
}

func InitServices(repos *Repositories) *Services {
	return &Services{
		UserService: servimpl.NewUserService(repos.UserRepository),
	}
}
