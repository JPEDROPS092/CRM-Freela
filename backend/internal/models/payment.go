package models

import (
	"time"

	"gorm.io/gorm"
)

// PaymentStatus represents the status of a payment
type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentPaid      PaymentStatus = "paid"
	PaymentOverdue   PaymentStatus = "overdue"
	PaymentCancelled PaymentStatus = "cancelled"
)

// PaymentMethod represents the method of payment
type PaymentMethod string

const (
	MethodBankTransfer PaymentMethod = "bank_transfer"
	MethodCreditCard   PaymentMethod = "credit_card"
	MethodPayPal       PaymentMethod = "paypal"
	MethodCash         PaymentMethod = "cash"
	MethodOther        PaymentMethod = "other"
)

// Payment represents a payment in the system
type Payment struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	UserID        uint           `json:"user_id" gorm:"not null;index"`
	User          User           `json:"-" gorm:"foreignKey:UserID"`
	ClientID      uint           `json:"client_id" gorm:"index"`
	Client        Client         `json:"-" gorm:"foreignKey:ClientID"`
	TaskID        *uint          `json:"task_id" gorm:"index"`
	Task          *Task          `json:"-" gorm:"foreignKey:TaskID"`
	Amount        float64        `json:"amount" gorm:"not null"`
	Currency      string         `json:"currency" gorm:"size:3;not null;default:'USD'"`
	Status        PaymentStatus  `json:"status" gorm:"size:20;not null;default:'pending'"`
	Method        PaymentMethod  `json:"method" gorm:"size:20"`
	Description   string         `json:"description" gorm:"type:text"`
	InvoiceNumber string         `json:"invoice_number" gorm:"size:50"`
	DueDate       time.Time      `json:"due_date"`
	PaidDate      *time.Time     `json:"paid_date"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

// BeforeCreate is a GORM hook that sets default values before creating a payment
func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	if p.Status == "" {
		p.Status = PaymentPending
	}
	if p.Currency == "" {
		p.Currency = "USD"
	}
	return nil
}

// MarkAsPaid marks the payment as paid
func (p *Payment) MarkAsPaid() {
	p.Status = PaymentPaid
	now := time.Now()
	p.PaidDate = &now
}

// CheckOverdue checks if the payment is overdue and updates the status if necessary
func (p *Payment) CheckOverdue() bool {
	if p.Status == PaymentPending && time.Now().After(p.DueDate) {
		p.Status = PaymentOverdue
		return true
	}
	return false
}
