package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jpcode092/crm-freela/configs"
	"github.com/jpcode092/crm-freela/internal/middleware"
	"github.com/jpcode092/crm-freela/internal/services"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

// Router representa o roteador da API
type Router struct {
	engine      *gin.Engine
	config      *configs.Config
	authService services.AuthService
	logger      logger.Logger
}

// NewRouter cria uma nova instância do roteador
func NewRouter(config *configs.Config, authService services.AuthService, logger logger.Logger) *Router {
	return &Router{
		engine:      gin.Default(),
		config:      config,
		authService: authService,
		logger:      logger,
	}
}

// SetupRoutes configura as rotas da API
func (r *Router) SetupRoutes(
	authHandler *AuthHandler,
	clientHandler *ClientHandler,
	taskHandler *TaskHandler,
	paymentHandler *PaymentHandler,
) {
	// Middleware global para CORS
	r.engine.Use(middleware.CORSMiddleware())

	// Grupo de rotas públicas
	public := r.engine.Group("/api/v1")
	{
		// Rotas de autenticação
		public.POST("/auth/register", authHandler.Register)
		public.POST("/auth/login", authHandler.Login)
	}

	// Grupo de rotas protegidas
	protected := r.engine.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(r.config))
	{
		// Rotas de autenticação
		protected.POST("/auth/refresh", authHandler.RefreshToken)

		// Rotas de usuário
		protected.GET("/user/profile", authHandler.GetProfile)

		// Rotas de clientes
		protected.POST("/clients", clientHandler.Create)
		protected.GET("/clients", clientHandler.List)
		protected.GET("/clients/:id", clientHandler.GetByID)
		protected.PUT("/clients/:id", clientHandler.Update)
		protected.DELETE("/clients/:id", clientHandler.Delete)

		// Rotas de tarefas
		protected.POST("/tasks", taskHandler.Create)
		protected.GET("/tasks", taskHandler.List)
		protected.GET("/tasks/:id", taskHandler.GetByID)
		protected.PUT("/tasks/:id", taskHandler.Update)
		protected.DELETE("/tasks/:id", taskHandler.Delete)

		// Rotas de pagamentos
		protected.POST("/payments", paymentHandler.CreatePayment)
		protected.GET("/payments", paymentHandler.ListPayments)
		protected.GET("/payments/:id", paymentHandler.GetPayment)
		protected.PUT("/payments/:id", paymentHandler.UpdatePayment)
		protected.DELETE("/payments/:id", paymentHandler.DeletePayment)
		protected.GET("/payments/client/:clientId", paymentHandler.GetPaymentByClientID)
	}
}

// Run inicia o servidor HTTP
func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
