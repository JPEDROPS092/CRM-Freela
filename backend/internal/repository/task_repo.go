package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/jpcode092/crm-freela/internal/models"
	"gorm.io/gorm"
)

// TaskRepository define a interface para operações de repositório de tarefas
type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id uint) (*models.Task, error)
	GetByUserID(userID uint, page, pageSize int) ([]models.Task, int64, error)
	GetByClientID(clientID uint, page, pageSize int) ([]models.Task, int64, error)
	Update(task *models.Task) error
	Delete(id uint) error
	List(page, pageSize int) ([]models.Task, int64, error)
	GetUpcoming(userID uint, days int) ([]models.Task, error)
	GetByStatus(userID uint, status models.TaskStatus, page, pageSize int) ([]models.Task, int64, error)
	CountByUserAndStatus(userID uint, status models.TaskStatus) (int64, error)
}

// taskRepository implementa a interface TaskRepository
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository cria uma nova instância de TaskRepository
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

// Create cria uma nova tarefa no banco de dados
func (r *taskRepository) Create(task *models.Task) error {
	result := r.db.Create(task)
	if result.Error != nil {
		return fmt.Errorf("erro ao criar tarefa: %w", result.Error)
	}
	return nil
}

// GetByID busca uma tarefa pelo ID
func (r *taskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	result := r.db.Preload("Client").First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("tarefa com ID %d não encontrada", id)
		}
		return nil, fmt.Errorf("erro ao buscar tarefa: %w", result.Error)
	}
	return &task, nil
}

// GetByUserID busca tarefas pelo ID do usuário com paginação
func (r *taskRepository) GetByUserID(userID uint, page, pageSize int) ([]models.Task, int64, error) {
	var tasks []models.Task
	var total int64

	// Conta o total de registros para o usuário
	if err := r.db.Model(&models.Task{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar tarefas do usuário: %w", err)
	}

	// Calcula o offset para paginação
	offset := (page - 1) * pageSize

	// Busca as tarefas do usuário com paginação
	result := r.db.Where("user_id = ?", userID).Preload("Client").Offset(offset).Limit(pageSize).Find(&tasks)
	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar tarefas do usuário: %w", result.Error)
	}

	return tasks, total, nil
}

// GetByClientID busca tarefas pelo ID do cliente com paginação
func (r *taskRepository) GetByClientID(clientID uint, page, pageSize int) ([]models.Task, int64, error) {
	var tasks []models.Task
	var total int64

	// Conta o total de registros para o cliente
	if err := r.db.Model(&models.Task{}).Where("client_id = ?", clientID).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar tarefas do cliente: %w", err)
	}

	// Calcula o offset para paginação
	offset := (page - 1) * pageSize

	// Busca as tarefas do cliente com paginação
	result := r.db.Where("client_id = ?", clientID).Preload("Client").Offset(offset).Limit(pageSize).Find(&tasks)
	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar tarefas do cliente: %w", result.Error)
	}

	return tasks, total, nil
}

// Update atualiza uma tarefa existente
func (r *taskRepository) Update(task *models.Task) error {
	result := r.db.Save(task)
	if result.Error != nil {
		return fmt.Errorf("erro ao atualizar tarefa: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("nenhuma tarefa foi atualizada")
	}
	return nil
}

// Delete remove uma tarefa pelo ID (soft delete)
func (r *taskRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Task{}, id)
	if result.Error != nil {
		return fmt.Errorf("erro ao excluir tarefa: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("tarefa com ID %d não encontrada", id)
	}
	return nil
}

// List retorna uma lista paginada de tarefas
func (r *taskRepository) List(page, pageSize int) ([]models.Task, int64, error) {
	var tasks []models.Task
	var total int64

	// Conta o total de registros
	if err := r.db.Model(&models.Task{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar tarefas: %w", err)
	}

	// Calcula o offset para paginação
	offset := (page - 1) * pageSize

	// Busca as tarefas com paginação
	result := r.db.Preload("Client").Offset(offset).Limit(pageSize).Find(&tasks)
	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar tarefas: %w", result.Error)
	}

	return tasks, total, nil
}

// GetUpcoming retorna as tarefas com prazo nos próximos X dias
func (r *taskRepository) GetUpcoming(userID uint, days int) ([]models.Task, error) {
	var tasks []models.Task
	now := time.Now()
	deadline := now.AddDate(0, 0, days)

	result := r.db.Where("user_id = ? AND due_date BETWEEN ? AND ? AND status != ?", 
		userID, now, deadline, models.TaskCompleted).
		Preload("Client").
		Order("due_date ASC").
		Find(&tasks)

	if result.Error != nil {
		return nil, fmt.Errorf("erro ao buscar tarefas próximas do prazo: %w", result.Error)
	}

	return tasks, nil
}

// GetByStatus busca tarefas pelo status com paginação
func (r *taskRepository) GetByStatus(userID uint, status models.TaskStatus, page, pageSize int) ([]models.Task, int64, error) {
	var tasks []models.Task
	var total int64

	// Conta o total de registros para o status
	if err := r.db.Model(&models.Task{}).Where("user_id = ? AND status = ?", userID, status).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("erro ao contar tarefas por status: %w", err)
	}

	// Calcula o offset para paginação
	offset := (page - 1) * pageSize

	// Busca as tarefas por status com paginação
	result := r.db.Where("user_id = ? AND status = ?", userID, status).
		Preload("Client").
		Offset(offset).
		Limit(pageSize).
		Find(&tasks)

	if result.Error != nil {
		return nil, 0, fmt.Errorf("erro ao listar tarefas por status: %w", result.Error)
	}

	return tasks, total, nil
}

// CountByUserAndStatus conta o número de tarefas por usuário e status
func (r *taskRepository) CountByUserAndStatus(userID uint, status models.TaskStatus) (int64, error) {
	var count int64
	result := r.db.Model(&models.Task{}).Where("user_id = ? AND status = ?", userID, status).Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("erro ao contar tarefas por usuário e status: %w", result.Error)
	}
	return count, nil
}
