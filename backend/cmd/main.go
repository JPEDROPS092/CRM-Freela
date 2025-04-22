package main

import (
	"log"

	"github.com/jpcode092/crm-freela/configs"
	"github.com/jpcode092/crm-freela/internal/api"
	"github.com/jpcode092/crm-freela/internal/models"
	"github.com/jpcode092/crm-freela/internal/repository"
	"github.com/jpcode092/crm-freela/internal/services"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

func main() {
	// Carrega as configurações
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Inicializa o logger
	logger := logger.NewLogger()

	// Inicializa o banco de dados
	db, err := configs.NewDatabase(config, logger)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Auto-migração dos modelos
	err = db.DB.AutoMigrate(
		&models.User{},
		&models.Client{},
		&models.Task{},
		&models.Payment{},
	)
	if err != nil {
		log.Fatalf("Erro ao migrar modelos: %v", err)
	}

	// Inicializa os repositórios
	userRepo := repository.NewUserRepository(db.DB)
	clientRepo := repository.NewClientRepository(db.DB)
	taskRepo := repository.NewTaskRepository(db.DB)
	paymentRepo := repository.NewPaymentRepository(db.DB)

	// Inicializa os serviços
	planService := services.NewPlanService(clientRepo, taskRepo, logger)
	authService := services.NewAuthService(userRepo, logger, config)
	clientService := services.NewClientService(clientRepo, planService, logger)
	taskService := services.NewTaskService(taskRepo, clientRepo, logger)
	paymentService := services.NewPaymentService(paymentRepo, clientRepo, taskRepo, logger)

	// Inicializa os handlers
	authHandler := api.NewAuthHandler(authService, logger)
	clientHandler := api.NewClientHandler(clientService, logger)
	taskHandler := api.NewTaskHandler(taskService, logger)
	paymentHandler := api.NewPaymentHandler(paymentService, logger)

	// Inicializa o router
	router := api.NewRouter(config, authService, logger)
	router.SetupRoutes(authHandler, clientHandler, taskHandler, paymentHandler)

	// Inicia o servidor
	logger.Info("Servidor iniciando na porta " + config.Server.Port)
	if err := router.Run(":" + config.Server.Port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
