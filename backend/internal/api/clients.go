package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jpcode092/crm-freela/internal/models"
	"github.com/jpcode092/crm-freela/internal/services"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

// ClientRequest representa os dados de requisição para criação/atualização de cliente
type ClientRequest struct {
	Name    string `json:"name" binding:"required,min=2"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address" binding:"required"`
	Status  string `json:"status" binding:"omitempty,oneof=active inactive blocked"`
}

// ClientHandler gerencia as requisições relacionadas a clientes
type ClientHandler struct {
	clientService services.ClientService
	logger        logger.Logger
}

// NewClientHandler cria uma nova instância de ClientHandler
func NewClientHandler(clientService services.ClientService, logger logger.Logger) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
		logger:        logger,
	}
}

// Create processa a requisição de criação de cliente
func (h *ClientHandler) Create(c *gin.Context) {
	var req ClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// Define o status como ativo por padrão
	status := models.ClientActive
	if req.Status != "" {
		status = models.ClientStatus(req.Status)
	}

	client, err := h.clientService.Create(
		userID.(uint),
		req.Name,
		req.Email,
		req.Phone,
		req.Address,
		status,
	)

	if err != nil {
		if err == services.ErrClientLimitExceeded {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
		return
	}

	c.JSON(http.StatusCreated, client)
}

// GetByID processa a requisição de busca de cliente por ID
func (h *ClientHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	client, err := h.clientService.GetByID(uint(id), userID.(uint))
	if err != nil {
		if err == services.ErrClientNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar cliente"})
		return
	}

	c.JSON(http.StatusOK, client)
}

// List processa a requisição de listagem de clientes
func (h *ClientHandler) List(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	clients, total, err := h.clientService.GetByUserID(userID.(uint), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar clientes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": clients,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

// Update processa a requisição de atualização de cliente
func (h *ClientHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req ClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// Define o status como ativo por padrão
	status := models.ClientActive
	if req.Status != "" {
		status = models.ClientStatus(req.Status)
	}

	client, err := h.clientService.Update(
		uint(id),
		userID.(uint),
		req.Name,
		req.Email,
		req.Phone,
		req.Address,
		status,
	)

	if err != nil {
		if err == services.ErrClientNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
			return
		}
		if err == services.ErrClientLimitExceeded {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente"})
		return
	}

	c.JSON(http.StatusOK, client)
}

// Delete processa a requisição de exclusão de cliente
func (h *ClientHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	err = h.clientService.Delete(uint(id), userID.(uint))
	if err != nil {
		if err == services.ErrClientNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir cliente"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
