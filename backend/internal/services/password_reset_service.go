package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/jpcode092/crm-freela/internal/models"
	"github.com/jpcode092/crm-freela/internal/repository"
	"github.com/jpcode092/crm-freela/pkg/email"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

var (
	ErrInvalidToken = errors.New("token inválido ou expirado")
	ErrUserNotFound = errors.New("usuário não encontrado")
)

// PasswordResetService define a interface para o serviço de recuperação de senha
type PasswordResetService interface {
	RequestReset(email string) error
	ValidateToken(token string) (*models.User, error)
	ResetPassword(token, newPassword string) error
}

type passwordResetService struct {
	userRepo     repository.UserRepository
	emailService email.EmailService
	logger       logger.Logger
}

// NewPasswordResetService cria uma nova instância de PasswordResetService
func NewPasswordResetService(userRepo repository.UserRepository, emailService email.EmailService, logger logger.Logger) PasswordResetService {
	return &passwordResetService{
		userRepo:     userRepo,
		emailService: emailService,
		logger:      logger,
	}
}

// generateToken gera um token aleatório para recuperação de senha
func (s *passwordResetService) generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RequestReset inicia o processo de recuperação de senha
func (s *passwordResetService) RequestReset(email string) error {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return ErrUserNotFound
	}

	token, err := s.generateToken()
	if err != nil {
		return err
	}

	// Salva o token e a data de expiração no usuário
	user.ResetToken = &token
	user.ResetTokenExpires = time.Now().Add(1 * time.Hour)
	
	err = s.userRepo.Update(user)
	if err != nil {
		return err
	}

	// Envia o email com o token
	return s.emailService.SendPasswordReset(user.Email, token)
}

// ValidateToken verifica se o token é válido e não expirou
func (s *passwordResetService) ValidateToken(token string) (*models.User, error) {
	user, err := s.userRepo.GetByResetToken(token)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if user.ResetTokenExpires.Before(time.Now()) {
		return nil, ErrInvalidToken
	}

	return user, nil
}

// ResetPassword redefine a senha do usuário
func (s *passwordResetService) ResetPassword(token, newPassword string) error {
	user, err := s.ValidateToken(token)
	if err != nil {
		return err
	}

	// Atualiza a senha e limpa o token
	user.Password = newPassword
	user.ResetToken = nil
	user.ResetTokenExpires = time.Time{}

	return s.userRepo.Update(user)
}
