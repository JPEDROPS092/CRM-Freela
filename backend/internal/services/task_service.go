package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/jpcode092/crm-freela/internal/models"
	"github.com/jpcode092/crm-freela/internal/repository"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

// Erros comuns do serviço de tarefas
var (
	ErrTaskNotFound     = errors.New("tarefa não encontrada")
	ErrClientNotActive  = errors.New("cliente não está ativo")
)

// TaskService define a interface para o serviço de tarefas
type TaskService interface {
	Create(userID, clientID uint, title, description string, priority models.TaskPriority, 
		dueDate *time.Time, estimatedHours, hourlyRate float64) (*models.Task, error)
	GetByID(id, userID uint) (*models.Task, error)
	GetByUserID(userID uint, page, pageSize int) ([]models.Task, int64, error)
	GetByClientID(clientID, userID uint, page, pageSize int) ([]models.Task, int64, error)
	Update(id, userID, clientID uint, title, description string, status models.TaskStatus, 
		priority models.TaskPriority, dueDate *time.Time, estimatedHours, actualHours, hourlyRate float64) (*models.Task, error)
	Delete(id, userID uint) error
	ChangeStatus(id, userID uint, status models.TaskStatus) error
	GetUpcoming(userID uint, days int) ([]models.Task, error)
	GetByStatus(userID uint, status models.TaskStatus, page, pageSize int) ([]models.Task, int64, error)
}

// taskService implementa a interface TaskService
type taskService struct {
	taskRepo   repository.TaskRepository
	clientRepo repository.ClientRepository
	logger     logger.Logger
}

// NewTaskService cria uma nova instância de TaskService
func NewTaskService(taskRepo repository.TaskRepository, clientRepo repository.ClientRepository, logger logger.Logger) TaskService {
	return &taskService{
		taskRepo:   taskRepo,
		clientRepo: clientRepo,
		logger:     logger,
	}
}

// Create cria uma nova tarefa
func (s *taskService) Create(userID, clientID uint, title, description string, priority models.TaskPriority, 
	dueDate *time.Time, estimatedHours, hourlyRate float64) (*models.Task, error) {
	
	// Verifica se o cliente existe e está ativo
	client, err := s.clientRepo.GetByID(clientID)
	if err != nil {
		return nil, ErrClientNotFound
	}

	// Verifica se o cliente pertence ao usuário
	if client.UserID != userID {
		return nil, ErrClientNotFound
	}

	// Verifica se o cliente está ativo
	if client.Status != models.ClientActive {
		return nil, ErrClientNotActive
	}

	// Cria uma nova tarefa
	task := &models.Task{
		UserID:         userID,
		ClientID:       clientID,
		Title:          title,
		Description:    description,
		Status:         models.TaskTodo,
		Priority:       priority,
		DueDate:        dueDate,
		EstimatedHours: estimatedHours,
		HourlyRate:     hourlyRate,
	}

	// Salva a tarefa no banco de dados
	if err := s.taskRepo.Create(task); err != nil {
		s.logger.Error(fmt.Sprintf("Erro ao criar tarefa: %v", err))
		return nil, fmt.Errorf("erro ao criar tarefa: %w", err)
	}

	return task, nil
}

// GetByID busca uma tarefa pelo ID
func (s *taskService) GetByID(id, userID uint) (*models.Task, error) {
	task, err := s.taskRepo.GetByID(id)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	// Verifica se a tarefa pertence ao usuário
	if task.UserID != userID {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

// GetByUserID busca tarefas pelo ID do usuário com paginação
func (s *taskService) GetByUserID(userID uint, page, pageSize int) ([]models.Task, int64, error) {
	return s.taskRepo.GetByUserID(userID, page, pageSize)
}

// GetByClientID busca tarefas pelo ID do cliente com paginação
func (s *taskService) GetByClientID(clientID, userID uint, page, pageSize int) ([]models.Task, int64, error) {
	// Verifica se o cliente existe e pertence ao usuário
	client, err := s.clientRepo.GetByID(clientID)
	if err != nil {
		return nil, 0, ErrClientNotFound
	}

	if client.UserID != userID {
		return nil, 0, ErrClientNotFound
	}

	return s.taskRepo.GetByClientID(clientID, page, pageSize)
}

// Update atualiza uma tarefa existente
func (s *taskService) Update(id, userID, clientID uint, title, description string, status models.TaskStatus, 
	priority models.TaskPriority, dueDate *time.Time, estimatedHours, actualHours, hourlyRate float64) (*models.Task, error) {
	
	// Busca a tarefa pelo ID
	task, err := s.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	// Se o cliente foi alterado, verifica se o novo cliente existe e está ativo
	if task.ClientID != clientID {
		client, err := s.clientRepo.GetByID(clientID)
		if err != nil {
			return nil, ErrClientNotFound
		}

		// Verifica se o cliente pertence ao usuário
		if client.UserID != userID {
			return nil, ErrClientNotFound
		}

		// Verifica se o cliente está ativo
		if client.Status != models.ClientActive {
			return nil, ErrClientNotActive
		}

		task.ClientID = clientID
	}

	// Atualiza os campos da tarefa
	task.Title = title
	task.Description = description
	task.Status = status
	task.Priority = priority
	task.DueDate = dueDate
	task.EstimatedHours = estimatedHours
	task.ActualHours = actualHours
	task.HourlyRate = hourlyRate

	// Se a tarefa foi concluída, registra a data de término
	if status == models.TaskCompleted && task.EndDate == nil {
		now := time.Now()
		task.EndDate = &now
	}

	// Se a tarefa foi iniciada, registra a data de início
	if (status == models.TaskInProgress || status == models.TaskReview || status == models.TaskCompleted) && task.StartDate == nil {
		now := time.Now()
		task.StartDate = &now
	}

	// Salva as alterações no banco de dados
	if err := s.taskRepo.Update(task); err != nil {
		s.logger.Error(fmt.Sprintf("Erro ao atualizar tarefa: %v", err))
		return nil, fmt.Errorf("erro ao atualizar tarefa: %w", err)
	}

	return task, nil
}

// Delete remove uma tarefa (soft delete)
func (s *taskService) Delete(id, userID uint) error {
	// Verifica se a tarefa existe e pertence ao usuário
	if _, err := s.GetByID(id, userID); err != nil {
		return err
	}

	// Remove a tarefa
	if err := s.taskRepo.Delete(id); err != nil {
		s.logger.Error(fmt.Sprintf("Erro ao excluir tarefa: %v", err))
		return fmt.Errorf("erro ao excluir tarefa: %w", err)
	}

	return nil
}

// ChangeStatus altera o status de uma tarefa
func (s *taskService) ChangeStatus(id, userID uint, status models.TaskStatus) error {
	// Busca a tarefa pelo ID
	task, err := s.GetByID(id, userID)
	if err != nil {
		return err
	}

	// Atualiza o status da tarefa
	task.Status = status

	// Se a tarefa foi concluída, registra a data de término
	if status == models.TaskCompleted && task.EndDate == nil {
		now := time.Now()
		task.EndDate = &now
	}

	// Se a tarefa foi iniciada, registra a data de início
	if (status == models.TaskInProgress || status == models.TaskReview || status == models.TaskCompleted) && task.StartDate == nil {
		now := time.Now()
		task.StartDate = &now
	}

	// Salva as alterações no banco de dados
	if err := s.taskRepo.Update(task); err != nil {
		s.logger.Error(fmt.Sprintf("Erro ao atualizar status da tarefa: %v", err))
		return fmt.Errorf("erro ao atualizar status da tarefa: %w", err)
	}

	return nil
}

// GetUpcoming retorna as tarefas com prazo nos próximos X dias
func (s *taskService) GetUpcoming(userID uint, days int) ([]models.Task, error) {
	return s.taskRepo.GetUpcoming(userID, days)
}

// GetByStatus busca tarefas pelo status com paginação
func (s *taskService) GetByStatus(userID uint, status models.TaskStatus, page, pageSize int) ([]models.Task, int64, error) {
	return s.taskRepo.GetByStatus(userID, status, page, pageSize)
}
