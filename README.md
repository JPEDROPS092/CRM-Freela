# Mini CRM de Freelancers

<div align="center">
  <img src="frontend/public/logo.png" alt="CRM Freelancer Logo" width="200"/>
  <p><strong>Um sistema completo para freelancers gerenciarem clientes, tarefas e pagamentos</strong></p>
</div>

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8.svg)](https://go.dev/)
[![Nuxt Version](https://img.shields.io/badge/Nuxt-3.x-00DC82.svg)](https://nuxt.com/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Índice

- [Visão Geral](#visão-geral)
- [Funcionalidades](#funcionalidades)
- [Arquitetura](#arquitetura)
- [Requisitos de Sistema](#requisitos-de-sistema)
- [Instalação e Configuração](#instalação-e-configuração)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Documentação da API](#documentação-da-api)
- [Desenvolvimento](#desenvolvimento)
- [Contribuição](#contribuição)
- [Licença](#licença)

## Visão Geral

O **Mini CRM de Freelancers** é uma solução completa para profissionais autônomos gerenciarem seus clientes, projetos, tarefas e pagamentos em um único lugar. O sistema oferece uma interface intuitiva e responsiva, permitindo o acesso de qualquer dispositivo.

### Demonstração

- **Backend API**: http://localhost:8080
- **Frontend**: http://localhost:3000
- **Documentação da API (Swagger)**: http://localhost:8080/swagger/index.html

## Funcionalidades

### Gerenciamento de Clientes
- Cadastro completo de clientes com informações detalhadas
- Visualização de histórico de projetos e pagamentos por cliente
- Categorização e status de clientes (ativo, inativo, arquivado)
- Limite de 3 clientes ativos no plano gratuito

### Controle de Tarefas e Projetos
- Criação de tarefas com prazos, prioridades e valores
- Fluxo de trabalho personalizável (orçamento, aprovado, em andamento, etc.)
- Visualização em formato Kanban ou lista
- Anexo de arquivos e documentos

### Registro de Pagamentos
- Controle de pagamentos recebidos e pendentes
- Geração de relatórios financeiros (plano premium)
- Registro de despesas relacionadas a projetos (plano premium)
- Exportação de dados financeiros

### Dashboard e Relatórios
- Visão geral de tarefas, prazos e pagamentos
- Métricas de desempenho e produtividade
- Calendário de compromissos e prazos
- Gráficos e visualizações personalizáveis (plano premium)

### Planos e Assinaturas
- Plano gratuito com recursos básicos
- Plano premium com recursos avançados
- Gerenciamento de assinatura e pagamentos

## Arquitetura

O sistema é construído com uma arquitetura moderna e escalável:

- **Backend**: Go (Golang) com Gin framework para uma API REST eficiente e de alta performance
- **Frontend**: Nuxt.js 3 (Vue.js) para uma interface de usuário moderna e responsiva
- **Banco de Dados**: PostgreSQL para armazenamento de dados robusto e relacional
- **Autenticação**: JWT (JSON Web Tokens) para autenticação segura e stateless
- **Documentação da API**: Swagger/OpenAPI para documentação interativa

## Requisitos de Sistema

### Backend
- Go 1.21+
- PostgreSQL 15+
- Docker (opcional para containerização)

### Frontend
- Node.js 18+
- NPM 9+ ou Yarn 1.22+

## Instalação e Configuração

### Backend

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/CRM-Freela.git
cd CRM-Freela/backend
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure o banco de dados:
- Crie um banco PostgreSQL
- Copie o arquivo `.env.example` para `.env`
- Configure as variáveis de ambiente no arquivo `.env`:
  ```
  DB_HOST=localhost
  DB_PORT=5432
  DB_USER=postgres
  DB_PASSWORD=sua_senha
  DB_NAME=crm_freela
  JWT_SECRET=seu_jwt_secret
  PORT=8080
  ```

4. Execute as migrações:
```bash
go run cmd/migrate/main.go
```

5. Inicie o servidor:
```bash
go run cmd/api/main.go
```

O servidor estará rodando em `http://localhost:8080`

### Frontend

1. Entre na pasta do frontend:
```bash
cd ../frontend
```

2. Instale as dependências:
```bash
npm install
```

3. Configure as variáveis de ambiente:
- Copie o arquivo `.env.example` para `.env`
- Configure as variáveis de ambiente:
  ```
  API_BASE=http://localhost:8080/api
  ```

4. Inicie o servidor de desenvolvimento:
```bash
npm run dev
```

O frontend estará disponível em `http://localhost:3000`

## Estrutura do Projeto

### Backend (Go)

```
backend/
├── cmd/                      # Pontos de entrada da aplicação
│   ├── api/                  # Servidor da API
│   └── migrate/              # Ferramenta de migração
├── configs/                  # Configurações da aplicação
├── internal/                 # Código interno da aplicação
│   ├── api/                  # Handlers da API
│   ├── middleware/           # Middlewares
│   ├── models/               # Modelos de dados
│   ├── repository/           # Camada de acesso a dados
│   └── services/             # Lógica de negócios
├── migrations/               # Migrações do banco de dados
├── pkg/                      # Código compartilhável
│   ├── logger/
│   └── util/
├── .env.example              # Exemplo de variáveis de ambiente
├── Dockerfile                # Configuração para Docker
├── go.mod                    # Dependências do Go
└── go.sum
```

### Frontend (Nuxt.js)

```
frontend/
├── assets/                   # Arquivos estáticos (CSS, imagens)
├── components/               # Componentes Vue reutilizáveis
├── layouts/                  # Layouts da aplicação
├── middleware/               # Middlewares do Nuxt
├── pages/                    # Páginas da aplicação
├── plugins/                  # Plugins do Nuxt
├── public/                   # Arquivos públicos
├── server/                   # API Server do Nuxt
├── store/                    # Estado da aplicação (Pinia)
├── types/                    # Tipos TypeScript
├── utils/                    # Funções utilitárias
├── .env.example              # Exemplo de variáveis de ambiente
├── nuxt.config.ts            # Configuração do Nuxt
└── package.json              # Dependências do projeto
```

## Documentação da API

A documentação completa da API está disponível através do Swagger UI em:

```
http://localhost:8080/swagger/index.html
```

### Endpoints Principais

#### Autenticação
- `POST /api/auth/register` - Registro de usuário
- `POST /api/auth/login` - Login
- `POST /api/auth/refresh` - Renovar token
- `GET /api/user/profile` - Obter perfil do usuário

#### Clientes
- `GET /api/clients` - Listar clientes
- `POST /api/clients` - Criar cliente
- `GET /api/clients/:id` - Buscar cliente
- `PUT /api/clients/:id` - Atualizar cliente
- `DELETE /api/clients/:id` - Remover cliente

#### Tarefas
- `GET /api/tasks` - Listar tarefas
- `POST /api/tasks` - Criar tarefa
- `GET /api/tasks/:id` - Buscar tarefa
- `PUT /api/tasks/:id` - Atualizar tarefa
- `DELETE /api/tasks/:id` - Remover tarefa
- `PUT /api/tasks/:id/status` - Atualizar status da tarefa

#### Pagamentos
- `GET /api/payments` - Listar pagamentos
- `POST /api/payments` - Criar pagamento
- `GET /api/payments/:id` - Buscar pagamento
- `PUT /api/payments/:id` - Atualizar pagamento
- `DELETE /api/payments/:id` - Remover pagamento

### Exemplos de Requisições

#### Login

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "usuario@exemplo.com",
    "password": "senha123"
  }'
```

#### Criar Cliente

```bash
curl -X POST http://localhost:8080/api/clients \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer SEU_TOKEN_JWT' \
  -d '{
    "name": "Nome do Cliente",
    "email": "cliente@exemplo.com",
    "phone": "(11) 98765-4321",
    "company": "Empresa XYZ",
    "status": "active"
  }'
```

## Desenvolvimento

### Backend

#### Executar Testes
```bash
cd backend
go test ./...
```

#### Gerar Documentação Swagger
```bash
cd backend
swag init -g cmd/api/main.go
```

### Frontend

#### Executar Testes
```bash
cd frontend
npm run test
```

#### Build para Produção
```bash
cd frontend
npm run build
```

## Contribuição

Contribuições são bem-vindas! Para contribuir:

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Faça commit das suas alterações (`git commit -m 'Adiciona nova feature'`)
4. Faça push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.
