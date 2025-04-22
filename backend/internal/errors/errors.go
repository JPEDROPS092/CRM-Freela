package errors

import "errors"

// Erros comuns da aplicação
var (
	ErrUserNotFound     = errors.New("usuário não encontrado")
	ErrEmailInUse       = errors.New("email já está em uso")
	ErrInvalidPassword  = errors.New("senha inválida")
	ErrUserDeactivated  = errors.New("usuário desativado")
	ErrInvalidToken     = errors.New("token inválido")
	ErrTokenExpired     = errors.New("token expirado")
	ErrInvalidCredentials = errors.New("credenciais inválidas")
)
