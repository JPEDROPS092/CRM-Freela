package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jpcode092/crm-freela/configs"
	"github.com/jpcode092/crm-freela/internal/models"
	"github.com/jpcode092/crm-freela/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Initialize logger
	logger := logger.NewLogger()
	logger.Info("Iniciando script de criação do usuário administrador")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logger.Fatal("Erro ao carregar arquivo .env")
	}

	// Initialize database connection
	dbConfig := &configs.Config{
		DB: configs.DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	}

	db, err := configs.NewDatabase(dbConfig, logger)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Erro ao conectar com o banco de dados: %v", err))
	}
	defer db.Close()

	// Create admin user
	adminPassword := "admin123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Erro ao gerar hash da senha:", err)
	}

	adminUser := &models.User{
		Name:     "Administrador",
		Email:    "admin@admin.com",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	result := db.DB.Create(adminUser)
	if result.Error != nil {
		if result.Error.Error() == "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)" {
			logger.Info("Usuário administrador já existe")
			return
		}
		logger.Fatal(fmt.Sprintf("Erro ao criar usuário administrador: %v", result.Error))
	}

	logger.Info("Usuário administrador criado com sucesso!")
	logger.Info(fmt.Sprintf("Email: %s", adminUser.Email))
	logger.Info(fmt.Sprintf("Senha: %s", adminPassword))
}
