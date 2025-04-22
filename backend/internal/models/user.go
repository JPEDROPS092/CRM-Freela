package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// PlanType represents the subscription plan type
type PlanType string

const (
	FreePlan  PlanType = "free"
	BasicPlan PlanType = "basic"
	ProPlan   PlanType = "pro"
)

// UserStatus represents the user status
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusBlocked  UserStatus = "blocked"
)

// UserRole represents the user role
type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

// User represents a user in the system
type User struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	Name             string         `json:"name" gorm:"size:100;not null"`
	Email            string         `json:"email" gorm:"size:100;not null;uniqueIndex"`
	Password         string         `json:"-" gorm:"size:100;not null"`
	Role             UserRole       `json:"role" gorm:"size:20;not null;default:'user'"`
	Plan             PlanType       `json:"plan" gorm:"size:20;not null;default:'free'"`
	Status           UserStatus     `json:"status" gorm:"size:20;not null;default:'active'"`
	ResetToken       *string        `json:"-" gorm:"size:100"`
	ResetTokenExpires time.Time     `json:"-"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
}

// BeforeSave is a GORM hook that hashes the password before saving
func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password != "" && len(u.Password) < 60 { // Verifica se a senha não está hasheada
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword checks if the provided password matches the stored hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// IsAllowedToCreateClient checks if the user can create a new client based on their plan
func (u *User) IsAllowedToCreateClient(currentClientCount int) bool {
	switch u.Plan {
	case FreePlan:
		return currentClientCount < 5
	case BasicPlan:
		return currentClientCount < 20
	case ProPlan:
		return true
	default:
		return false
	}
}

// IsAllowedToCreateTask checks if the user can create a new task based on their plan
func (u *User) IsAllowedToCreateTask(currentTaskCount int) bool {
	switch u.Plan {
	case FreePlan:
		return currentTaskCount < 10
	case BasicPlan:
		return currentTaskCount < 50
	case ProPlan:
		return true
	default:
		return false
	}
}
