package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jpcode092/crm-freela/internal/models"
	"github.com/jpcode092/crm-freela/internal/services"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

// TaskRequest representa os dados de requisição para criação/atualização de tarefa
type TaskRequest struct {
	ClientID    uint              `json:"client_id" binding:"required"`
	Title       string            `json:"title" binding:"required"`
	Description string            `json:"description" binding:"required"`
	Priority    models.TaskPriority `json:"priority" binding:"required,oneof=low medium high"`
	DueDate     string            `json:"due_date" binding:"required"`
	EstimatedHours float64        `json:"estimated_hours" binding:"required"`
	HourlyRate    float64        `json:"hourly_rate" binding:"required"`
}

// TaskHandler gerencia as requisições relacionadas a tarefas
type TaskHandler struct {
	taskService services.TaskService
	logger      logger.Logger
}

// NewTaskHandler cria uma nova instância de TaskHandler
func NewTaskHandler(taskService services.TaskService, logger logger.Logger) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
		logger:      logger,
	}
}

// Create processa a requisição de criação de tarefa
func (h *TaskHandler) Create(c *gin.Context) {
	var req TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// Parse da data de vencimento
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data de vencimento inválida"})
		return
	}

	task, err := h.taskService.Create(
		userID.(uint),
		req.ClientID,
		req.Title,
		req.Description,
		req.Priority,
		&dueDate,
		req.EstimatedHours,
		req.HourlyRate,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar tarefa"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetByID processa a requisição de busca de tarefa por ID
func (h *TaskHandler) GetByID(c *gin.Context) {
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

	task, err := h.taskService.GetByID(uint(id), userID.(uint))
	if err != nil {
		if err == services.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefa"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// List processa a requisição de listagem de tarefas
func (h *TaskHandler) List(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	tasks, total, err := h.taskService.GetByUserID(userID.(uint), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar tarefas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

// Update processa a requisição de atualização de tarefa
func (h *TaskHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// Parse da data de vencimento
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data de vencimento inválida"})
		return
	}

	task, err := h.taskService.Update(
		uint(id),
		userID.(uint),
		req.ClientID,
		req.Title,
		req.Description,
		models.TaskInProgress, // Status padrão
		req.Priority,
		&dueDate,
		req.EstimatedHours,
		req.HourlyRate,
		0, // Horas trabalhadas
	)

	if err != nil {
		if err == services.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar tarefa"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Delete processa a requisição de exclusão de tarefa
func (h *TaskHandler) Delete(c *gin.Context) {
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

	err = h.taskService.Delete(uint(id), userID.(uint))
	if err != nil {
		if err == services.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir tarefa"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
