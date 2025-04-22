package models

import (
	"time"

	"gorm.io/gorm"
)

// ClientStatus represents the status of a client
type ClientStatus string

const (
	ClientActive   ClientStatus = "active"
	ClientInactive ClientStatus = "inactive"
	ClientArchived ClientStatus = "archived"
)

// Client represents a client in the system
type Client struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null;index"`
	User        User           `json:"-" gorm:"foreignKey:UserID"`
	Name        string         `json:"name" gorm:"size:100;not null"`
	Email       string         `json:"email" gorm:"size:100"`
	Phone       string         `json:"phone" gorm:"size:20"`
	Company     string         `json:"company" gorm:"size:100"`
	Address     string         `json:"address" gorm:"size:200"`
	Notes       string         `json:"notes" gorm:"type:text"`
	Status      ClientStatus   `json:"status" gorm:"size:20;not null;default:'active'"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Tasks       []Task         `json:"tasks,omitempty" gorm:"foreignKey:ClientID"`
	Payments    []Payment      `json:"payments,omitempty" gorm:"foreignKey:ClientID"`
}

// BeforeCreate is a GORM hook that sets default values before creating a client
func (c *Client) BeforeCreate(tx *gorm.DB) error {
	if c.Status == "" {
		c.Status = ClientActive
	}
	return nil
}
