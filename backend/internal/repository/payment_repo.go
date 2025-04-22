package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/jpcode092/crm-freela/internal/models"
	"gorm.io/gorm"
)

// PaymentRepository define a interface para operações de repositório de pagamentos
type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetByID(id uint) (*models.Payment, error)
	GetByUserID(userID uint, page, pageSize int) ([]models.Payment, int64, error)
	GetByClientID(clientID uint, page, pageSize int) ([]models.Payment, int64, error)
	GetByTaskID(taskID uint) ([]models.Payment, error)
	Update(payment *models.Payment) error
	Delete(id uint) error
	List(page, pageSize int) ([]models.Payment, int64, error)
	GetOverdue(userID uint) ([]models.Payment, error)
	GetByStatus(userID uint, status models.PaymentStatus, page, pageSize int) ([]models.Payment, int64, error)
	GetSummaryByPeriod(userID uint, startDate, endDate time.Time) (float64, error)
}

// paymentRepository implementa a interface PaymentRepository
type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository cria uma nova instância de PaymentRepository
func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

// Create cria um novo pagamento no banco de dados
func (r *paymentRepository) Create(payment *models.Payment) error {
	result := r.db.Create(payment)
	if result.Error != nil {
		return fmt.Errorf("erro ao criar pagamento: %w", result.Error)
	}
	return nil
}

// GetByID busca um pagamento pelo ID
func (r *paymentRepository) GetByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	result := r.db.Preload("Client").Preload("Task").First(&payment, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("pagamento com ID %d não encontrado", id)
		}
		return nil, fmt.Errorf("erro ao buscar pagamento: %w", result.Error)
	}
	return &payment, nil
}

// GetByUserID busca pagamentos pelo ID do usuário com paginação
func (r *paymentRepository) GetByUserID(userID uint, page, pageSize int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	// Conta o total de registros para o usuário
	if err := r.db.Model(&models.Payment{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar pagamentos do usuário: %w", err)
	}

	// Calcula o offset para paginação
	offset := (page - 1) * pageSize

	// Busca os pagamentos do usuário com paginação
	result := r.db.Where("user_id = ?", userID).
		Preload("Client").
		Preload("Task").
		Offset(offset).
		Limit(pageSize).
		Order("due_date DESC").
		Find(&payments)

	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar pagamentos do usuário: %w", result.Error)
	}

	return payments, total, nil
}

// GetByClientID busca pagamentos pelo ID do cliente com paginação
func (r *paymentRepository) GetByClientID(clientID uint, page, pageSize int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	// Conta o total de registros para o cliente
	if err := r.db.Model(&models.Payment{}).Where("client_id = ?", clientID).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar pagamentos do cliente: %w", err)
	}

	// Calcula o offset para paginação
	offset := (page - 1) * pageSize

	// Busca os pagamentos do cliente com paginação
	result := r.db.Where("client_id = ?", clientID).
		Preload("Task").
		Offset(offset).
		Limit(pageSize).
		Order("due_date DESC").
		Find(&payments)

	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar pagamentos do cliente: %w", result.Error)
	}

	return payments, total, nil
}

// GetByTaskID busca pagamentos pelo ID da tarefa
func (r *paymentRepository) GetByTaskID(taskID uint) ([]models.Payment, error) {
	var payments []models.Payment

	result := r.db.Where("task_id = ?", taskID).
		Preload("Client").
		Order("due_date DESC").
		Find(&payments)

	if result.Error != nil {
		return nil, fmt.Errorf("erro ao listar pagamentos da tarefa: %w", result.Error)
	}

	return payments, nil
}

// Update atualiza um pagamento existente
func (r *paymentRepository) Update(payment *models.Payment) error {
	result := r.db.Save(payment)
	if result.Error != nil {
		return fmt.Errorf("erro ao atualizar pagamento: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("nenhum pagamento foi atualizado")
	}
	return nil
}

// Delete remove um pagamento pelo ID (soft delete)
func (r *paymentRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Payment{}, id)
	if result.Error != nil {
		return fmt.Errorf("erro ao excluir pagamento: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("pagamento com ID %d não encontrado", id)
	}
	return nil
}

// List retorna uma lista paginada de pagamentos
func (r *paymentRepository) List(page, pageSize int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	// Conta o total de registros
	if err := r.db.Model(&models.Payment{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar pagamentos: %w", err)
	}

	// Calcula o offset para paginação
	offset := (page - 1) * pageSize

	// Busca os pagamentos com paginação
	result := r.db.Preload("Client").
		Preload("Task").
		Offset(offset).
		Limit(pageSize).
		Order("due_date DESC").
		Find(&payments)

	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar pagamentos: %w", result.Error)
	}

	return payments, total, nil
}

// GetOverdue retorna os pagamentos vencidos
func (r *paymentRepository) GetOverdue(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	now := time.Now()

	result := r.db.Where("user_id = ? AND due_date < ? AND status = ?", 
		userID, now, models.PaymentPending).
		Preload("Client").
		Preload("Task").
		Order("due_date ASC").
		Find(&payments)

	if result.Error != nil {
		return nil, fmt.Errorf("erro ao buscar pagamentos vencidos: %w", result.Error)
	}

	return payments, nil
}

// GetByStatus busca pagamentos pelo status com paginação
func (r *paymentRepository) GetByStatus(userID uint, status models.PaymentStatus, page, pageSize int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	// Conta o total de registros para o status
	if err := r.db.Model(&models.Payment{}).Where("user_id = ? AND status = ?", userID, status).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar pagamentos por status: %w", err)
	}

	// Calcula o offset para paginação
	offset := (page - 1) * pageSize

	// Busca os pagamentos por status com paginação
	result := r.db.Where("user_id = ? AND status = ?", userID, status).
		Preload("Client").
		Preload("Task").
		Offset(offset).
		Limit(pageSize).
		Order("due_date DESC").
		Find(&payments)

	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar pagamentos por status: %w", result.Error)
	}

	return payments, total, nil
}

// GetSummaryByPeriod retorna o total de pagamentos recebidos em um período
func (r *paymentRepository) GetSummaryByPeriod(userID uint, startDate, endDate time.Time) (float64, error) {
	var total float64

	result := r.db.Model(&models.Payment{}).
		Select("COALESCE(SUM(amount), 0) as total").
		Where("user_id = ? AND status = ? AND paid_date BETWEEN ? AND ?", 
			userID, models.PaymentPaid, startDate, endDate).
		Scan(&total)

	if result.Error != nil {
		return 0, fmt.Errorf("erro ao calcular total de pagamentos no período: %w", result.Error)
	}

	return total, nil
}
