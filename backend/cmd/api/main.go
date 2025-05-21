package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jpcode092/crm-freela/configs"
	"github.com/jpcode092/crm-freela/docs"
	"github.com/jpcode092/crm-freela/internal/api"
	"github.com/jpcode092/crm-freela/internal/middleware"
	"github.com/jpcode092/crm-freela/internal/repository"
	"github.com/jpcode092/crm-freela/internal/services"
	"github.com/jpcode092/crm-freela/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           CRM Freela API
// @version         1.0
// @description     API para gerenciamento de freelancers e projetos
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Initialize logger
	logger := logger.NewLogger()
	logger.Info("Starting CRM Freelancer API")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logger.Warn("No .env file found, using environment variables")
	}

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Swagger docs
	docs.SwaggerInfo.Title = "CRM Freela API"
	docs.SwaggerInfo.Description = "API para gerenciamento de freelancers e projetos"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Initialize database
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
		logger.Fatal(fmt.Sprintf("Failed to initialize database: %v", err))
	}
	defer db.Close()

	// Initialize router
	router := gin.Default()

	// Add custom logger middleware
	router.Use(middleware.LoggerMiddleware(logger))

	// Configure CORS
	config := cors.DefaultConfig()
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000" // Default frontend URL if not set
	}
	config.AllowOrigins = []string{frontendURL}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// Initialize repositories, services and handlers
	userRepo := repository.NewUserRepository(db.DB)
	authService := services.NewAuthService(userRepo, logger, dbConfig)
	authHandler := api.NewAuthHandler(authService, logger)

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	// Public routes
	router.POST("/api/auth/register", authHandler.Register)
	router.POST("/api/auth/login", authHandler.Login)
	router.POST("/api/auth/refresh", authHandler.RefreshToken)

	// Protected routes
	authRoutes := router.Group("/api")
	authRoutes.Use(middleware.JWTAuth())
	{
		// User routes
		authRoutes.GET("/user/profile", authHandler.GetProfile)
		
		// TODO: Implement these routes
		// Client routes
		// Task routes
		// Payment routes
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info(fmt.Sprintf("Server running on port %s", port))
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
