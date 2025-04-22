package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/jpcode092/crm-freela/internal/models"
	"github.com/jpcode092/crm-freela/internal/repository"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

// Erros comuns do serviço de pagamentos
var (
	ErrPaymentNotFound = errors.New("pagamento não encontrado")
	ErrInvalidAmount   = errors.New("valor do pagamento inválido")
)

// PaymentService define a interface para o serviço de pagamentos
type PaymentService interface {
	Create(userID, clientID uint, taskID *uint, amount float64, currency string, 
		method models.PaymentMethod, description, invoiceNumber string, dueDate time.Time) (*models.Payment, error)
	GetByID(id, userID uint) (*models.Payment, error)
	GetByUserID(userID uint, page, pageSize int) ([]models.Payment, int64, error)
	GetByClientID(clientID, userID uint, page, pageSize int) ([]models.Payment, int64, error)
	GetByTaskID(taskID, userID uint) ([]models.Payment, error)
	Update(id, userID, clientID uint, taskID *uint, amount float64, currency string, 
		status models.PaymentStatus, method models.PaymentMethod, description, invoiceNumber string, 
		dueDate time.Time, paidDate *time.Time) (*models.Payment, error)
	Delete(id, userID uint) error
	MarkAsPaid(id, userID uint, paidDate time.Time) error
	GetOverdue(userID uint) ([]models.Payment, error)
	GetByStatus(userID uint, status models.PaymentStatus, page, pageSize int) ([]models.Payment, int64, error)
	GetSummaryByPeriod(userID uint, startDate, endDate time.Time) (float64, error)
	CheckAndUpdateOverduePayments(userID uint) (int, error)
}

// paymentService implementa a interface PaymentService
type paymentService struct {
	paymentRepo repository.PaymentRepository
	clientRepo  repository.ClientRepository
	taskRepo    repository.TaskRepository
	logger      logger.Logger
}

// NewPaymentService cria uma nova instância de PaymentService
func NewPaymentService(
	paymentRepo repository.PaymentRepository, 
	clientRepo repository.ClientRepository,
	taskRepo repository.TaskRepository,
	logger logger.Logger,
) PaymentService {
	return &paymentService{
		paymentRepo: paymentRepo,
		clientRepo:  clientRepo,
		taskRepo:    taskRepo,
		logger:      logger,
	}
}

// Create cria um novo pagamento
func (s *paymentService) Create(userID, clientID uint, taskID *uint, amount float64, currency string, 
	method models.PaymentMethod, description, invoiceNumber string, dueDate time.Time) (*models.Payment, error) {
	
	// Verifica se o valor é válido
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// Verifica se o cliente existe
	client, err := s.clientRepo.GetByID(clientID)
	if err != nil {
		return nil, ErrClientNotFound
	}

	// Verifica se o cliente pertence ao usuário
	if client.UserID != userID {
		return nil, ErrClientNotFound
	}

	// Verifica se a tarefa existe, se fornecida
	if taskID != nil {
		task, err := s.taskRepo.GetByID(*taskID)
		if err != nil {
			return nil, ErrTaskNotFound
		}

		// Verifica se a tarefa pertence ao usuário
		if task.UserID != userID {
			return nil, ErrTaskNotFound
		}

		// Verifica se a tarefa pertence ao cliente
		if task.ClientID != clientID {
			return nil, errors.New("a tarefa não pertence ao cliente especificado")
		}
	}

	// Cria um novo pagamento
	payment := &models.Payment{
		UserID:        userID,
		ClientID:      clientID,
		TaskID:        taskID,
		Amount:        amount,
		Currency:      currency,
		Status:        models.PaymentPending,
		Method:        method,
		Description:   description,
		InvoiceNumber: invoiceNumber,
		DueDate:       dueDate,
	}

	// Salva o pagamento no banco de dados
	if err := s.paymentRepo.Create(payment); err != nil {
		s.logger.Error(fmt.Sprintf("Erro ao criar pagamento: %v", err))
		return nil, fmt.Errorf("erro ao criar pagamento: %w", err)
	}

	return payment, nil
}

// GetByID busca um pagamento pelo ID
func (s *paymentService) GetByID(id, userID uint) (*models.Payment, error) {
	payment, err := s.paymentRepo.GetByID(id)
	if err != nil {
		return nil, ErrPaymentNotFound
	}

	// Verifica se o pagamento pertence ao usuário
	if payment.UserID != userID {
		return nil, ErrPaymentNotFound
	}

	return payment, nil
}

// GetByUserID busca pagamentos pelo ID do usuário com paginação
func (s *paymentService) GetByUserID(userID uint, page, pageSize int) ([]models.Payment, int64, error) {
	return s.paymentRepo.GetByUserID(userID, page, pageSize)
}

// GetByClientID busca pagamentos pelo ID do cliente com paginação
func (s *paymentService) GetByClientID(clientID, userID uint, page, pageSize int) ([]models.Payment, int64, error) {
	// Verifica se o cliente existe e pertence ao usuário
	client, err := s.clientRepo.GetByID(clientID)
	if err != nil {
		return nil, 0, ErrClientNotFound
	}

	if client.UserID != userID {
		return nil, 0, ErrClientNotFound
	}

	return s.paymentRepo.GetByClientID(clientID, page, pageSize)
}

// GetByTaskID busca pagamentos pelo ID da tarefa
func (s *paymentService) GetByTaskID(taskID, userID uint) ([]models.Payment, error) {
	// Verifica se a tarefa existe e pertence ao usuário
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	if task.UserID != userID {
		return nil, ErrTaskNotFound
	}

	return s.paymentRepo.GetByTaskID(taskID)
}

// Update atualiza um pagamento existente
func (s *paymentService) Update(id, userID, clientID uint, taskID *uint, amount float64, currency string, 
	status models.PaymentStatus, method models.PaymentMethod, description, invoiceNumber string, 
	dueDate time.Time, paidDate *time.Time) (*models.Payment, error) {
	
	// Verifica se o valor é válido
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// Busca o pagamento pelo ID
	payment, err := s.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	// Verifica se o cliente existe
	client, err := s.clientRepo.GetByID(clientID)
	if err != nil {
		return nil, ErrClientNotFound
	}

	// Verifica se o cliente pertence ao usuário
	if client.UserID != userID {
		return nil, ErrClientNotFound
	}

	// Verifica se a tarefa existe, se fornecida
	if taskID != nil {
		task, err := s.taskRepo.GetByID(*taskID)
		if err != nil {
			return nil, ErrTaskNotFound
		}

		// Verifica se a tarefa pertence ao usuário
		if task.UserID != userID {
			return nil, ErrTaskNotFound
		}

		// Verifica se a tarefa pertence ao cliente
		if task.ClientID != clientID {
			return nil, errors.New("a tarefa não pertence ao cliente especificado")
		}
	}

	// Atualiza os campos do pagamento
	payment.ClientID = clientID
	payment.TaskID = taskID
	payment.Amount = amount
	payment.Currency = currency
	payment.Status = status
	payment.Method = method
	payment.Description = description
	payment.InvoiceNumber = invoiceNumber
	payment.DueDate = dueDate
	payment.PaidDate = paidDate

	// Se o status for "pago" e não houver data de pagamento, define a data atual
	if status == models.PaymentPaid && payment.PaidDate == nil {
		now := time.Now()
		payment.PaidDate = &now
	}

	// Salva as alterações no banco de dados
	if err := s.paymentRepo.Update(payment); err != nil {
		s.logger.Error(fmt.Sprintf("Erro ao atualizar pagamento: %v", err))
		return nil, fmt.Errorf("erro ao atualizar pagamento: %w", err)
	}

	return payment, nil
}

// Delete remove um pagamento (soft delete)
func (s *paymentService) Delete(id, userID uint) error {
	// Verifica se o pagamento existe e pertence ao usuário
	if _, err := s.GetByID(id, userID); err != nil {
		return err
	}

	// Remove o pagamento
	if err := s.paymentRepo.Delete(id); err != nil {
		s.logger.Error(fmt.Sprintf("Erro ao excluir pagamento: %v", err))
		return fmt.Errorf("erro ao excluir pagamento: %w", err)
	}

	return nil
}

// MarkAsPaid marca um pagamento como pago
func (s *paymentService) MarkAsPaid(id, userID uint, paidDate time.Time) error {
	// Busca o pagamento pelo ID
	payment, err := s.GetByID(id, userID)
	if err != nil {
		return err
	}

	// Atualiza o status e a data de pagamento
	payment.Status = models.PaymentPaid
	payment.PaidDate = &paidDate

	// Salva as alterações no banco de dados
	if err := s.paymentRepo.Update(payment); err != nil {
		s.logger.Error(fmt.Sprintf("Erro ao marcar pagamento como pago: %v", err))
		return fmt.Errorf("erro ao marcar pagamento como pago: %w", err)
	}

	return nil
}

// GetOverdue retorna os pagamentos vencidos
func (s *paymentService) GetOverdue(userID uint) ([]models.Payment, error) {
	return s.paymentRepo.GetOverdue(userID)
}

// GetByStatus busca pagamentos pelo status com paginação
func (s *paymentService) GetByStatus(userID uint, status models.PaymentStatus, page, pageSize int) ([]models.Payment, int64, error) {
	return s.paymentRepo.GetByStatus(userID, status, page, pageSize)
}

// GetSummaryByPeriod retorna o total de pagamentos recebidos em um período
func (s *paymentService) GetSummaryByPeriod(userID uint, startDate, endDate time.Time) (float64, error) {
	return s.paymentRepo.GetSummaryByPeriod(userID, startDate, endDate)
}

// CheckAndUpdateOverduePayments verifica e atualiza o status de pagamentos vencidos
func (s *paymentService) CheckAndUpdateOverduePayments(userID uint) (int, error) {
	// Busca todos os pagamentos pendentes
	payments, count, err := s.paymentRepo.GetByStatus(userID, models.PaymentPending, 1, 1000)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Erro ao buscar pagamentos pendentes: %v", err))
		return 0, fmt.Errorf("erro ao buscar pagamentos pendentes: %w", err)
	}

	// Se não houver pagamentos pendentes, retorna 0
	if count == 0 {
		return 0, nil
	}

	// Verifica quais pagamentos estão vencidos
	now := time.Now()
	updatedCount := 0

	for _, payment := range payments {
		if payment.DueDate.Before(now) {
			// Atualiza o status para vencido
			payment.Status = models.PaymentOverdue

			// Salva as alterações no banco de dados
			if err := s.paymentRepo.Update(&payment); err != nil {
				s.logger.Error(fmt.Sprintf("Erro ao atualizar status de pagamento vencido: %v", err))
				continue
			}

			updatedCount++
		}
	}

	return updatedCount, nil
}
