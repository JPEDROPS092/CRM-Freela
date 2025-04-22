package configs

import (
	"fmt"
	"time"

	"github.com/jpcode092/crm-freela/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Database é uma estrutura que encapsula a conexão com o banco de dados
type Database struct {
	DB     *gorm.DB
	Logger logger.Logger
}

// NewDatabase cria uma nova instância de conexão com o banco de dados
func NewDatabase(config *Config, logger logger.Logger) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.Name,
		config.DB.SSLMode,
	)

	// Configuração do GORM
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		PrepareStmt: true, // Prepara as consultas SQL para melhor performance
	}

	// Conecta ao banco de dados
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		logger.Error(fmt.Sprintf("Falha ao conectar ao banco de dados: %v", err))
		return nil, err
	}

	// Configura o pool de conexões
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error(fmt.Sprintf("Falha ao configurar o pool de conexões: %v", err))
		return nil, err
	}

	// Define o número máximo de conexões abertas
	sqlDB.SetMaxOpenConns(100)
	// Define o número máximo de conexões inativas
	sqlDB.SetMaxIdleConns(10)
	// Define o tempo máximo de vida de uma conexão
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.Info("Conexão com o banco de dados estabelecida com sucesso")

	return &Database{
		DB:     db,
		Logger: logger,
	}, nil
}

// Close fecha a conexão com o banco de dados
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		d.Logger.Error(fmt.Sprintf("Erro ao obter a conexão SQL: %v", err))
		return err
	}
	
	if err := sqlDB.Close(); err != nil {
		d.Logger.Error(fmt.Sprintf("Erro ao fechar a conexão com o banco de dados: %v", err))
		return err
	}
	
	d.Logger.Info("Conexão com o banco de dados fechada com sucesso")
	return nil
}

// Ping verifica se a conexão com o banco de dados está ativa
func (d *Database) Ping() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		d.Logger.Error(fmt.Sprintf("Erro ao obter a conexão SQL: %v", err))
		return err
	}
	
	if err := sqlDB.Ping(); err != nil {
		d.Logger.Error(fmt.Sprintf("Erro ao pingar o banco de dados: %v", err))
		return err
	}
	
	d.Logger.Info("Conexão com o banco de dados está ativa")
	return nil
}
