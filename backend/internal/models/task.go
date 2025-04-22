package models

import (
	"time"

	"gorm.io/gorm"
)

// TaskStatus represents the status of a task
type TaskStatus string

const (
	TaskTodo       TaskStatus = "todo"
	TaskInProgress TaskStatus = "in_progress"
	TaskReview     TaskStatus = "review"
	TaskCompleted  TaskStatus = "completed"
	TaskCancelled  TaskStatus = "cancelled"
)

// TaskPriority represents the priority of a task
type TaskPriority string

const (
	PriorityLow    TaskPriority = "low"
	PriorityMedium TaskPriority = "medium"
	PriorityHigh   TaskPriority = "high"
)

// Task represents a task in the system
type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null;index"`
	User        User           `json:"-" gorm:"foreignKey:UserID"`
	ClientID    uint           `json:"client_id" gorm:"index"`
	Client      Client         `json:"-" gorm:"foreignKey:ClientID"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Status      TaskStatus     `json:"status" gorm:"size:20;not null;default:'todo'"`
	Priority    TaskPriority   `json:"priority" gorm:"size:20;not null;default:'medium'"`
	DueDate     *time.Time     `json:"due_date"`
	StartDate   *time.Time     `json:"start_date"`
	EndDate     *time.Time     `json:"end_date"`
	EstimatedHours float64     `json:"estimated_hours"`
	ActualHours    float64     `json:"actual_hours"`
	HourlyRate     float64     `json:"hourly_rate"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Payments    []Payment      `json:"payments,omitempty" gorm:"foreignKey:TaskID"`
}

// BeforeCreate is a GORM hook that sets default values before creating a task
func (t *Task) BeforeCreate(tx *gorm.DB) error {
	if t.Status == "" {
		t.Status = TaskTodo
	}
	if t.Priority == "" {
		t.Priority = PriorityMedium
	}
	return nil
}

// CalculateTotal calculates the total amount for the task based on hourly rate and actual hours
func (t *Task) CalculateTotal() float64 {
	return t.HourlyRate * t.ActualHours
}
