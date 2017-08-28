package config

import (
	"os"
	"fmt"
)

var (
	Http = httpConfig{
		ListenPort: getenv("HTTP_LISTEN_PORT", "80"),
		AppKey: os.Getenv("HTTP_APP_KEY"),
		SecretKey: os.Getenv("HTTP_SECRET_KEY"),
	}

	Redis = redisConfig{
		Address: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), getenv("REDIS_PORT", "6379")),
		Password: getenv("REDIS_PASSWORD", ""),
		DB: getenv("REDIS_DB", "0"),
	}
)

type httpConfig struct {
	ListenPort string
	AppKey  string
	SecretKey  string
}

type redisConfig struct {
	Address  string
	Password string
	DB       string
}


func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}
