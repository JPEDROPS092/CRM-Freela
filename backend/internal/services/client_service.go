package services

import (
	"errors"
	"fmt"

	"github.com/jpcode092/crm-freela/internal/models"
	"github.com/jpcode092/crm-freela/internal/repository"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

// Erros comuns do serviço de clientes
var (
	ErrClientNotFound = errors.New("cliente não encontrado")
)

// ClientService define a interface para o serviço de clientes
type ClientService interface {
	Create(userID uint, name, email, phone, address string, status models.ClientStatus) (*models.Client, error)
	GetByID(id, userID uint) (*models.Client, error)
	GetByUserID(userID uint, page, pageSize int) ([]models.Client, int64, error)
	Update(id, userID uint, name, email, phone, address string, status models.ClientStatus) (*models.Client, error)
	Delete(id, userID uint) error
	CountByUser(userID uint) (int64, error)
}

// clientService implementa a interface ClientService
type clientService struct {
	clientRepo repository.ClientRepository
	planService PlanService
	logger     logger.Logger
}

// NewClientService cria uma nova instância de ClientService
func NewClientService(clientRepo repository.ClientRepository, planService PlanService, logger logger.Logger) ClientService {
	return &clientService{
		clientRepo: clientRepo,
		planService: planService,
		logger:     logger,
	}
}

// Create cria um novo cliente
func (s *clientService) Create(userID uint, name, email, phone, address string, status models.ClientStatus) (*models.Client, error) {
	// Verifica se o usuário pode criar mais clientes
	if err := s.planService.CanCreateClient(userID); err != nil {
		return nil, err
	}

	client := &models.Client{
		UserID:  userID,
		Name:    name,
		Email:   email,
		Phone:   phone,
		Address: address,
		Status:  status,
	}

	if err := s.clientRepo.Create(client); err != nil {
		return nil, fmt.Errorf("erro ao criar cliente: %w", err)
	}

	return client, nil
}

// GetByID busca um cliente pelo ID
func (s *clientService) GetByID(id, userID uint) (*models.Client, error) {
	client, err := s.clientRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Verifica se o cliente pertence ao usuário
	if client.UserID != userID {
		return nil, ErrClientNotFound
	}

	return client, nil
}

// GetByUserID busca clientes pelo ID do usuário com paginação
func (s *clientService) GetByUserID(userID uint, page, pageSize int) ([]models.Client, int64, error) {
	return s.clientRepo.GetByUserID(userID, page, pageSize)
}

// Update atualiza um cliente existente
func (s *clientService) Update(id, userID uint, name, email, phone, address string, status models.ClientStatus) (*models.Client, error) {
	client, err := s.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	// Atualiza os campos
	client.Name = name
	client.Email = email
	client.Phone = phone
	client.Address = address
	client.Status = status

	if err := s.clientRepo.Update(client); err != nil {
		return nil, fmt.Errorf("erro ao atualizar cliente: %w", err)
	}

	return client, nil
}

// Delete remove um cliente
func (s *clientService) Delete(id, userID uint) error {
	// Verifica se o cliente existe e pertence ao usuário
	if _, err := s.GetByID(id, userID); err != nil {
		return err
	}

	return s.clientRepo.Delete(id)
}

// CountByUser conta o número de clientes por usuário
func (s *clientService) CountByUser(userID uint) (int64, error) {
	return s.clientRepo.CountByUser(userID)
}
