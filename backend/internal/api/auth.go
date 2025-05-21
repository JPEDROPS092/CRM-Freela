package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpcode092/crm-freela/internal/errors"
	"github.com/jpcode092/crm-freela/internal/services"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

// AuthRequest representa os dados de requisição para autenticação
type AuthRequest struct {
	Name     string `json:"name" binding:"required,min=3" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"123456"`
}

// LoginRequest representa os dados de requisição para login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

// AuthResponse representa a resposta da autenticação
type AuthResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  struct {
		ID    uint   `json:"id" example:"1"`
		Name  string `json:"name" example:"John Doe"`
		Email string `json:"email" example:"john@example.com"`
	} `json:"user"`
}

// AuthHandler gerencia as requisições relacionadas a autenticação
type AuthHandler struct {
	authService services.AuthService
	logger      logger.Logger
}

// NewAuthHandler cria uma nova instância de AuthHandler
func NewAuthHandler(authService services.AuthService, logger logger.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

// Register godoc
// @Summary      Registrar novo usuário
// @Description  Registra um novo usuário no sistema
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body AuthRequest true "Dados do usuário"
// @Success      201  {object}  AuthResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      409  {object}  map[string]interface{} "E-mail já cadastrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno"
// @Router       /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	h.logger.Info("Recebida requisição de registro")
	
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Erro ao processar dados de registro: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}
	
	h.logger.Info("Processando registro para email: " + req.Email)

	user, err := h.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		if err == errors.ErrEmailInUse {
			h.logger.Warn("Tentativa de registro com email já cadastrado: " + req.Email)
			c.JSON(http.StatusConflict, gin.H{"error": "E-mail já cadastrado"})
			return
		}
		h.logger.Error("Erro ao registrar usuário: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar usuário"})
		return
	}

	// Gera um token para o usuário recém-registrado
	_, token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		h.logger.Error("Erro ao gerar token após registro: " + err.Error())
		// Mesmo que não consiga gerar o token, o registro foi bem-sucedido
		c.JSON(http.StatusCreated, gin.H{
			"message": "Usuário registrado com sucesso, mas não foi possível gerar o token",
			"user": gin.H{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			},
		})
		return
	}

	h.logger.Info("Usuário registrado com sucesso: " + user.Email)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário registrado com sucesso",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"plan":  user.Plan,
		},
	})
}

// Login godoc
// @Summary      Login de usuário
// @Description  Autentica um usuário no sistema
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Credenciais do usuário"
// @Success      200  {object}  AuthResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      401  {object}  map[string]interface{} "Credenciais inválidas"
// @Failure      500  {object}  map[string]interface{} "Erro interno"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	h.logger.Info("Recebida requisição de login")
	
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Erro ao processar dados de login: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}
	
	h.logger.Info("Processando login para email: " + req.Email)

	user, token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		switch err {
		case errors.ErrUserNotFound:
			h.logger.Warn("Tentativa de login com usuário não encontrado: " + req.Email)
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		case errors.ErrInvalidPassword:
			h.logger.Warn("Tentativa de login com senha inválida: " + req.Email)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Senha inválida"})
		case errors.ErrUserDeactivated:
			h.logger.Warn("Tentativa de login com usuário desativado: " + req.Email)
			c.JSON(http.StatusForbidden, gin.H{"error": "Usuário desativado"})
		default:
			h.logger.Error("Erro ao fazer login: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao fazer login"})
		}
		return
	}

	h.logger.Info("Login realizado com sucesso: " + user.Email)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login realizado com sucesso",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"plan":  user.Plan,
		},
	})
}

// RefreshToken godoc
// @Summary      Renovar token
// @Description  Renova o token JWT do usuário
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]interface{} "Token inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno"
// @Router       /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
		return
	}

	newToken, err := h.authService.RefreshToken(token)
	if err != nil {
		switch err {
		case errors.ErrInvalidToken:
			h.logger.Warn("Tentativa de renovar token inválido: " + token)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		case errors.ErrTokenExpired:
			h.logger.Warn("Tentativa de renovar token expirado: " + token)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
		case errors.ErrUserDeactivated:
			h.logger.Warn("Tentativa de renovar token com usuário desativado: " + token)
			c.JSON(http.StatusForbidden, gin.H{"error": "Usuário desativado"})
		default:
			h.logger.Error("Erro ao renovar token: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao renovar token"})
		}
		return
	}

	h.logger.Info("Token renovado com sucesso: " + newToken)
	c.JSON(http.StatusOK, gin.H{
		"message": "Token renovado com sucesso",
		"token":   newToken,
	})
}

// GetProfile godoc
// @Summary      Obter perfil do usuário
// @Description  Retorna os dados do perfil do usuário autenticado
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{} "Não autorizado"
// @Failure      500  {object}  map[string]interface{} "Erro interno"
// @Router       /user/profile [get]
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	user, err := h.authService.GetUserByID(userID.(uint))
	if err != nil {
		if err == errors.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		h.logger.Error("Erro ao buscar perfil: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar perfil"})
		return
	}

	h.logger.Info("Perfil do usuário obtido com sucesso: " + user.Email)
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"plan":  user.Plan,
		},
	})
}
