package config

import "time"

var (
	CookieName     = "session_id"
	CookieSecure   = false // 🔁 change in production
	CookieLifetime = 7 * 24 * time.Hour
)
