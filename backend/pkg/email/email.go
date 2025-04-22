package email

import (
	"fmt"
	"net/smtp"
)

// EmailService define a interface para envio de emails
type EmailService interface {
	SendPasswordReset(to, token string) error
}

type emailService struct {
	from     string
	password string
	host     string
	port     string
}

// NewEmailService cria uma nova instância de EmailService
func NewEmailService(from, password, host, port string) EmailService {
	return &emailService{
		from:     from,
		password: password,
		host:     host,
		port:     port,
	}
}

// SendPasswordReset envia um email com o token de recuperação de senha
func (s *emailService) SendPasswordReset(to, token string) error {
	subject := "Recuperação de Senha - CRM Freela"
	body := fmt.Sprintf(`
		<h2>Recuperação de Senha</h2>
		<p>Você solicitou a recuperação de senha. Use o link abaixo para redefinir sua senha:</p>
		<p><a href="http://localhost:3000/auth/reset-password?token=%s">Redefinir Senha</a></p>
		<p>Se você não solicitou a recuperação de senha, ignore este email.</p>
		<p>O link é válido por 1 hora.</p>
	`, token)

	msg := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body)

	auth := smtp.PlainAuth("", s.from, s.password, s.host)
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	return smtp.SendMail(addr, auth, s.from, []string{to}, []byte(msg))
}
