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

type PaymentHandler struct {
	paymentService services.PaymentService
	logger         logger.Logger
}

func NewPaymentHandler(paymentService services.PaymentService, logger logger.Logger) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		logger:        logger,
	}
}

type CreatePaymentRequest struct {
	ClientID      uint                 `json:"client_id" binding:"required"`
	TaskID        *uint                `json:"task_id"`
	Amount        float64              `json:"amount" binding:"required"`
	Description   string               `json:"description" binding:"required"`
	Method        models.PaymentMethod `json:"method" binding:"required"`
	Status        string               `json:"status" binding:"required"`
	DueDate       string               `json:"due_date" binding:"required"`
	PaymentDate   string               `json:"payment_date"`
}

// CreatePayment handles payment creation requests
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	userID := c.GetUint("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid due date format"})
		return
	}

	createdPayment, err := h.paymentService.Create(
		userID,
		req.ClientID,
		req.TaskID,
		req.Amount,
		"USD", // Default currency
		req.Method,
		req.Description,
		"", // Invoice number can be empty
		dueDate,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPayment)
}

type UpdatePaymentRequest struct {
	Amount      float64              `json:"amount"`
	Description string               `json:"description"`
	Method      models.PaymentMethod `json:"method"`
	Status      string               `json:"status"`
	DueDate     string               `json:"due_date"`
	PaymentDate string               `json:"payment_date"`
	TaskID      *uint                `json:"task_id"`
}

// UpdatePayment handles payment update requests
func (h *PaymentHandler) UpdatePayment(c *gin.Context) {
	userID := c.GetUint("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment ID"})
		return
	}

	var req UpdatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid due date format"})
		return
	}

	var paidDate *time.Time
	if req.PaymentDate != "" {
		pd, err := time.Parse("2006-01-02", req.PaymentDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment date format"})
			return
		}
		paidDate = &pd
	}

	updatedPayment, err := h.paymentService.Update(
		uint(id),
		userID,
		0, // ClientID is not updated
		req.TaskID,
		req.Amount,
		"USD", // Default currency
		models.PaymentStatus(req.Status),
		req.Method,
		req.Description,
		"", // Invoice number can be empty
		dueDate,
		paidDate,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPayment)
}

// ListPayments handles requests to list all payments
func (h *PaymentHandler) ListPayments(c *gin.Context) {
	userID := c.GetUint("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	payments, total, err := h.paymentService.GetByUserID(userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payments": payments,
		"total":    total,
	})
}

// GetPaymentByClientID handles requests to get payments by client ID
func (h *PaymentHandler) GetPaymentByClientID(c *gin.Context) {
	userID := c.GetUint("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	clientID, err := strconv.ParseUint(c.Param("clientId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	payments, total, err := h.paymentService.GetByClientID(uint(clientID), userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payments": payments,
		"total":    total,
	})
}

// GetPayment handles requests to get a specific payment by ID
func (h *PaymentHandler) GetPayment(c *gin.Context) {
	userID := c.GetUint("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment ID"})
		return
	}

	payment, err := h.paymentService.GetByID(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// DeletePayment handles payment deletion requests
func (h *PaymentHandler) DeletePayment(c *gin.Context) {
	userID := c.GetUint("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment ID"})
		return
	}

	if err := h.paymentService.Delete(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment deleted successfully"})
}
