package config

import "time"

var (
	CookieName     = "session_id"
	CookieSecure   = false // 🔁 change en prod
	CookieLifetime = 7 * 24 * time.Hour
)
