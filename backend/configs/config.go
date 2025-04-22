package configs

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config representa as configurações da aplicação
type Config struct {
	Server   ServerConfig
	DB       DBConfig
	JWT      JWTConfig
}

// ServerConfig representa as configurações do servidor
type ServerConfig struct {
	Port string
}

// DBConfig representa as configurações do banco de dados
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// JWTConfig representa as configurações do JWT
type JWTConfig struct {
	Secret        string
	AccessTokenTTL time.Duration
}

// LoadConfig carrega as configurações da aplicação
func LoadConfig() (*Config, error) {
	// Carrega o arquivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "crm_freela"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:        getEnv("JWT_SECRET", "your-256-bit-secret"),
			AccessTokenTTL: time.Hour * 24, // 24 horas
		},
	}, nil
}

// getEnv retorna o valor de uma variável de ambiente ou um valor padrão
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
