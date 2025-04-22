package repository

import (
	"errors"
	"fmt"

	"github.com/jpcode092/crm-freela/internal/models"
	"gorm.io/gorm"
)

// ClientRepository define a interface para operações de repositório de clientes
type ClientRepository interface {
	Create(client *models.Client) error
	GetByID(id uint) (*models.Client, error)
	GetByUserID(userID uint, page, pageSize int) ([]models.Client, int64, error)
	Update(client *models.Client) error
	Delete(id uint) error
	List(page, pageSize int) ([]models.Client, int64, error)
	CountByUser(userID uint) (int64, error)
}

// clientRepository implementa a interface ClientRepository
type clientRepository struct {
	db *gorm.DB
}

// NewClientRepository cria uma nova instância de ClientRepository
func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{
		db: db,
	}
}

// Create cria um novo cliente no banco de dados
func (r *clientRepository) Create(client *models.Client) error {
	result := r.db.Create(client)
	if result.Error != nil {
		return fmt.Errorf("erro ao criar cliente: %w", result.Error)
	}
	return nil
}

// GetByID busca um cliente pelo ID
func (r *clientRepository) GetByID(id uint) (*models.Client, error) {
	var client models.Client
	result := r.db.First(&client, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("cliente com ID %d não encontrado", id)
		}
		return nil, fmt.Errorf("erro ao buscar cliente: %w", result.Error)
	}
	return &client, nil
}

// GetByUserID busca clientes pelo ID do usuário com paginação
func (r *clientRepository) GetByUserID(userID uint, page, pageSize int) ([]models.Client, int64, error) {
	var clients []models.Client
	var total int64

	offset := (page - 1) * pageSize

	// Conta o total de registros
	if err := r.db.Model(&models.Client{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar clientes: %w", err)
	}

	// Busca os registros com paginação
	result := r.db.Where("user_id = ?", userID).Offset(offset).Limit(pageSize).Find(&clients)
	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar clientes: %w", result.Error)
	}

	return clients, total, nil
}

// Update atualiza um cliente existente
func (r *clientRepository) Update(client *models.Client) error {
	result := r.db.Save(client)
	if result.Error != nil {
		return fmt.Errorf("erro ao atualizar cliente: %w", result.Error)
	}
	return nil
}

// Delete remove um cliente pelo ID (soft delete)
func (r *clientRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Client{}, id)
	if result.Error != nil {
		return fmt.Errorf("erro ao remover cliente: %w", result.Error)
	}
	return nil
}

// List retorna uma lista paginada de clientes
func (r *clientRepository) List(page, pageSize int) ([]models.Client, int64, error) {
	var clients []models.Client
	var total int64

	offset := (page - 1) * pageSize

	// Conta o total de registros
	if err := r.db.Model(&models.Client{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar clientes: %w", err)
	}

	// Busca os registros com paginação
	result := r.db.Offset(offset).Limit(pageSize).Find(&clients)
	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar clientes: %w", result.Error)
	}

	return clients, total, nil
}

// CountByUser conta o número de clientes por usuário
func (r *clientRepository) CountByUser(userID uint) (int64, error) {
	var count int64
	result := r.db.Model(&models.Client{}).Where("user_id = ?", userID).Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("erro ao contar clientes por usuário: %w", result.Error)
	}
	return count, nil
}
