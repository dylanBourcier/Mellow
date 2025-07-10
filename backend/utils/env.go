package utils

import (
	"log"
	"os"
	"strconv"
	"time"
)

func GetEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func GetEnvAsBool(key string, defaultVal bool) bool {
	if valStr, ok := os.LookupEnv(key); ok {
		val, err := strconv.ParseBool(valStr)
		if err == nil {
			return val
		}
	}
	return defaultVal
}

func GetEnvAsDuration(key string, defaultVal time.Duration) time.Duration {
	if valStr, ok := os.LookupEnv(key); ok {
		val, err := time.ParseDuration(valStr)
		if err == nil {
			return val
		}
		log.Printf("Erreur parsing duration %s: %v", key, err)
	}
	return defaultVal
}
