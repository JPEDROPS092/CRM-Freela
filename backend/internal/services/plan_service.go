package services

import (
	"errors"

	"github.com/jpcode092/crm-freela/internal/models"
	"github.com/jpcode092/crm-freela/internal/repository"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

const (
	FreePlanClientLimit = 5
	FreePlanTaskLimit   = 10
)

var (
	ErrClientLimitExceeded = errors.New("limite de clientes do plano gratuito excedido")
	ErrTaskLimitExceeded   = errors.New("limite de tarefas do plano gratuito excedido")
)

// PlanService define a interface para o serviço de planos
type PlanService interface {
	CanCreateClient(userID uint) error
	CanCreateTask(userID uint) error
}

type planService struct {
	clientRepo repository.ClientRepository
	taskRepo   repository.TaskRepository
	logger     logger.Logger
}

// NewPlanService cria uma nova instância de PlanService
func NewPlanService(clientRepo repository.ClientRepository, taskRepo repository.TaskRepository, logger logger.Logger) PlanService {
	return &planService{
		clientRepo: clientRepo,
		taskRepo:   taskRepo,
		logger:     logger,
	}
}

// CanCreateClient verifica se o usuário pode criar mais clientes
func (s *planService) CanCreateClient(userID uint) error {
	// TODO: Implementar verificação de plano premium
	// Por enquanto, assume que todos os usuários estão no plano gratuito
	count, err := s.clientRepo.CountByUser(userID)
	if err != nil {
		return err
	}

	if count >= FreePlanClientLimit {
		return ErrClientLimitExceeded
	}

	return nil
}

// CanCreateTask verifica se o usuário pode criar mais tarefas
func (s *planService) CanCreateTask(userID uint) error {
	// TODO: Implementar verificação de plano premium
	// Por enquanto, assume que todos os usuários estão no plano gratuito
	count, err := s.taskRepo.CountByUserAndStatus(userID, models.TaskTodo)
	if err != nil {
		return err
	}

	if count >= FreePlanTaskLimit {
		return ErrTaskLimitExceeded
	}

	return nil
}
