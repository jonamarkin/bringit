package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env            string
	HTTPAddr       string
	DatabaseURL    string
	FrontendOrigin string
	PublicBaseURL  string
	JWTSecret      string
	CookieDomain   string
	ResendAPIKey   string
	MailSender     string
}

func Load() Config {
	_ = godotenv.Load()

	resendAPIKey := getEnv("RESEND_API_KEY", "")
	if resendAPIKey == "" {
		resendAPIKey = getEnv("SMTP_PASS", "")
	}

	return Config{
		Env:            getEnv("APP_ENV", "development"),
		HTTPAddr:       getEnv("HTTP_ADDR", ":8080"),
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://bringit:bringit@localhost:5432/bringit?sslmode=disable"),
		FrontendOrigin: getEnv("FRONTEND_ORIGIN", "http://localhost:3000"),
		PublicBaseURL:  getEnv("PUBLIC_BASE_URL", "http://localhost:3000"),
		JWTSecret:      getEnv("JWT_SECRET", "dev-only-change-this-bringit-secret"),
		CookieDomain:   getEnv("COOKIE_DOMAIN", ""),
		ResendAPIKey:   resendAPIKey,
		MailSender:     getEnv("SMTP_SENDER", "BringIt <updates@example.com>"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
