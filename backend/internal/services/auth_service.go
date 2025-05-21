package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jpcode092/crm-freela/configs"
	apperrors "github.com/jpcode092/crm-freela/internal/errors"
	"github.com/jpcode092/crm-freela/internal/models"
	"github.com/jpcode092/crm-freela/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Claims representa os claims do JWT
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// AuthService define a interface do serviço de autenticação
type AuthService interface {
	Register(name, email, password string) (*models.User, error)
	Login(email, password string) (*models.User, string, error)
	RefreshToken(token string) (string, error)
	GetUserByID(id uint) (*models.User, error)
}

// authService implementa a interface AuthService
type authService struct {
	userRepo models.UserRepository
	logger   logger.Logger
	config   *configs.Config
}

// NewAuthService cria uma nova instância de AuthService
func NewAuthService(userRepo models.UserRepository, logger logger.Logger, config *configs.Config) AuthService {
	return &authService{
		userRepo: userRepo,
		logger:  logger,
		config:  config,
	}
}

// Register registra um novo usuário
func (s *authService) Register(name, email, password string) (*models.User, error) {
	// Verifica se o email já está em uso
	existingUser, err := s.userRepo.GetByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, apperrors.ErrEmailInUse
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Cria o usuário
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Status:   models.UserStatusActive,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login autentica um usuário
func (s *authService) Login(email, password string) (*models.User, string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if err == models.ErrRecordNotFound {
			return nil, "", apperrors.ErrUserNotFound
		}
		return nil, "", err
	}

	if user.Status != models.UserStatusActive {
		return nil, "", apperrors.ErrUserDeactivated
	}

	// Verifica a senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", apperrors.ErrInvalidPassword
	}

	// Gera o token JWT
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.JWT.AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}

// RefreshToken renova um token JWT
func (s *authService) RefreshToken(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperrors.ErrInvalidToken
		}
		return []byte(s.config.JWT.Secret), nil
	})

	if err != nil {
		return "", apperrors.ErrInvalidToken
	}

	if !token.Valid {
		return "", apperrors.ErrInvalidToken
	}

	user, err := s.GetUserByID(claims.UserID)
	if err != nil {
		return "", err
	}

	if user.Status != models.UserStatusActive {
		return "", apperrors.ErrUserDeactivated
	}

	// Gera um novo token
	newClaims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.JWT.AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	return newToken.SignedString([]byte(s.config.JWT.Secret))
}

// GetUserByID busca um usuário por ID
func (s *authService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if err == models.ErrRecordNotFound {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
