package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jpcode092/crm-freela/pkg/logger"
)

// LoggerMiddleware registra informações sobre as requisições HTTP
func LoggerMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Tempo de início da requisição
		startTime := time.Now()

		// Processa a requisição
		c.Next()

		// Calcula o tempo de resposta
		latency := time.Since(startTime)

		// Obtém informações da requisição
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		clientIP := c.ClientIP()

		// Converte o status para string
		statusStr := ""
		switch {
		case status >= 200 && status < 300:
			statusStr = "\033[32m" + string(rune(status)) + "\033[0m" // Verde para sucesso
		case status >= 300 && status < 400:
			statusStr = "\033[36m" + string(rune(status)) + "\033[0m" // Ciano para redirecionamento
		case status >= 400 && status < 500:
			statusStr = "\033[33m" + string(rune(status)) + "\033[0m" // Amarelo para erro do cliente
		default:
			statusStr = "\033[31m" + string(rune(status)) + "\033[0m" // Vermelho para erro do servidor
		}

		// Registra a requisição
		log.RequestInfo(method, path, clientIP, statusStr, latency)
	}
}
