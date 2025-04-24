package config

import "time"

type JWTConfig struct {
	SecretKey     string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
}

var JWT = JWTConfig{
	SecretKey:     "your-secret-key-change-in-production", // Change this in production
	TokenExpiry:   time.Hour * 24,                         // 24 hours
	RefreshExpiry: time.Hour * 24 * 7,                     // 7 days
} 