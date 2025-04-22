package models

import "errors"

// ErrRecordNotFound é retornado quando um registro não é encontrado
var ErrRecordNotFound = errors.New("registro não encontrado")

// UserRepository define a interface para operações de persistência de usuários
type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	List(page, pageSize int) ([]User, int64, error)
}

// ClientRepository define a interface para operações de persistência de clientes
type ClientRepository interface {
	Create(client *Client) error
	Update(client *Client) error
	Delete(id uint) error
	GetByID(id uint) (*Client, error)
	GetByUserID(userID uint, page, pageSize int) ([]Client, int64, error)
	CountByUserID(userID uint) (int64, error)
}

// TaskRepository define a interface para operações de persistência de tarefas
type TaskRepository interface {
	Create(task *Task) error
	Update(task *Task) error
	Delete(id uint) error
	GetByID(id uint) (*Task, error)
	GetByUserID(userID uint, page, pageSize int) ([]Task, int64, error)
	GetByClientID(clientID uint, page, pageSize int) ([]Task, int64, error)
	CountByUserID(userID uint) (int64, error)
}

// PaymentRepository define a interface para operações de persistência de pagamentos
type PaymentRepository interface {
	Create(payment *Payment) error
	Update(payment *Payment) error
	Delete(id uint) error
	GetByID(id uint) (*Payment, error)
	GetByUserID(userID uint, page, pageSize int) ([]Payment, int64, error)
	GetByClientID(clientID uint, page, pageSize int) ([]Payment, int64, error)
}
