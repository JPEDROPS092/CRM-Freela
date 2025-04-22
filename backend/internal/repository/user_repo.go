package repository

import (
	"errors"
	"fmt"

	"github.com/jpcode092/crm-freela/internal/models"
	"gorm.io/gorm"
)

// UserRepository define a interface para operações de repositório de usuários
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByResetToken(token string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	List(page, pageSize int) ([]models.User, int64, error)
	CountByPlan(plan models.PlanType) (int64, error)
}

// userRepository implementa a interface UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância de UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create cria um novo usuário no banco de dados
func (r *userRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("erro ao criar usuário: %w", result.Error)
	}
	return nil
}

// GetByID busca um usuário pelo ID
func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("usuário com ID %d não encontrado", id)
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", result.Error)
	}
	return &user, nil
}

// GetByEmail busca um usuário pelo email
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("usuário com email %s não encontrado", email)
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", result.Error)
	}
	return &user, nil
}

// GetByResetToken busca um usuário pelo token de recuperação de senha
func (r *userRepository) GetByResetToken(token string) (*models.User, error) {
	var user models.User
	result := r.db.Where("reset_token = ?", token).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("token inválido ou expirado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", result.Error)
	}
	return &user, nil
}

// Update atualiza um usuário existente
func (r *userRepository) Update(user *models.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("nenhum usuário foi atualizado")
	}
	return nil
}

// Delete remove um usuário pelo ID (soft delete)
func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return fmt.Errorf("erro ao excluir usuário: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("usuário com ID %d não encontrado", id)
	}
	return nil
}

// List retorna uma lista paginada de usuários
func (r *userRepository) List(page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Conta o total de registros
	if err := r.db.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar usuários: %w", err)
	}

	// Calcula o offset para paginação
	offset := (page - 1) * pageSize

	// Busca os usuários com paginação
	result := r.db.Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar usuários: %w", result.Error)
	}

	return users, total, nil
}

// CountByPlan conta o número de usuários por tipo de plano
func (r *userRepository) CountByPlan(plan models.PlanType) (int64, error) {
	var count int64
	result := r.db.Model(&models.User{}).Where("plan = ?", plan).Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("erro ao contar usuários por plano: %w", result.Error)
	}
	return count, nil
}
