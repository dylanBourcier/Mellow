package config

import (
	"mellow/utils"
	"time"

	"github.com/joho/godotenv" // si tu utilises godotenv
)

// var (
// 	CookieName     = "session_id"
// 	CookieSecure   = false // 🔁 change in production
// 	CookieLifetime = 7 * 24 * time.Hour
// )

var (
	CookieName     string
	CookieSecure   bool
	CookieLifetime time.Duration
)

func Load() {
	// Charge les variables d'environnement depuis le .env (optionnel)
	_ = godotenv.Load()

	// Nom du cookie
	CookieName = utils.GetEnv("COOKIE_NAME", "session")

	// Est-ce que le cookie doit être sécurisé (HTTPS only)
	CookieSecure = utils.GetEnvAsBool("COOKIE_SECURE", false)

	// Durée de vie du cookie
	CookieLifetime = utils.GetEnvAsDuration("COOKIE_LIFETIME", time.Hour*24*7)
}
