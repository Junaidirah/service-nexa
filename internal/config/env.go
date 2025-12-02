package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	dbUrl            string
	jwtSecret        string
	jwtRefreshSecret string
	jwtExpire        string
	jwtRefreshExpire string
	port             string
}
type SmtpConfig struct {
	host     string
	port     string
	username string
	password string
}

func LoadEnv() (*Config, *SmtpConfig) {
	_ = godotenv.Load()

	config := &Config{
		dbUrl:            os.Getenv("DB_URL"),
		jwtSecret:        os.Getenv("JWT_SECRET"),
		jwtRefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
		jwtExpire:        os.Getenv("JWT_EXPIRE"),
		jwtRefreshExpire: os.Getenv("JWT_REFRESH_EXPIRE"),
		port:             os.Getenv("PORT"),
	}
	if config.dbUrl == "" {
		panic("DB_URL not set")
	}
	if config.jwtSecret == "" {
		panic("JWT_SECRET not set")
	}
	if config.jwtRefreshSecret == "" {
		panic("JWT_REFRESH_SECRET not set")
	}
	if config.jwtExpire == "" {
		panic("JWT_EXPIRE not set")
	}
	if config.jwtRefreshExpire == "" {
		panic("JWT_REFRESH_EXPIRE not set")
	}
	if config.port == "" {
		panic("PORT not set")
	}
	smtpConfig := &SmtpConfig{
		host:     os.Getenv("SMTP_HOST"),
		port:     os.Getenv("SMTP_PORT"),
		username: os.Getenv("SMTP_USERNAME"),
		password: os.Getenv("SMTP_PASSWORD"),
	}
	if smtpConfig.host == "" || smtpConfig.port == "" {
		panic("SMTP_HOST dan SMTP_PORT wajib di-set")
	}
	return config, smtpConfig
}
