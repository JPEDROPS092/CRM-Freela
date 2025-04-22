# Mini CRM de Freelancers

Um sistema CRM completo para freelancers gerenciarem clientes, tarefas e pagamentos.

## Visão Geral da Arquitetura

O sistema é construído com:

- **Backend**: Go (Golang) com Gin framework
- **Frontend**: Nuxt.js 3 (Vue.js)
- **Banco de Dados**: PostgreSQL
- **Autenticação**: JWT

## Estrutura do Projeto

O projeto está dividido em duas partes principais:

- `/backend` - API REST em Go
- `/frontend` - Interface de usuário em Nuxt.js

## Funcionalidades Principais

- Gerenciamento de clientes
- Controle de tarefas e projetos
- Registro de pagamentos
- Dashboard com métricas
- Planos gratuito e premium

## Requisitos de Desenvolvimento

### Backend
- Go 1.21+
- PostgreSQL 15+
- Docker (opcional)

### Frontend
- Node.js 18+
- Nuxt.js 3
- Vue.js 3
- Tailwind CSS

## Instalação e Execução

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
- Configure as variáveis de ambiente no arquivo `.env`

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
- Configure a URL da API em `VITE_API_URL`

4. Inicie o servidor de desenvolvimento:
```bash
npm run dev
```

O frontend estará disponível em `http://localhost:3000`

## Endpoints da API

### Autenticação
- `POST /api/v1/auth/register` - Registro de usuário
- `POST /api/v1/auth/login` - Login
- `POST /api/v1/auth/refresh` - Renovar token
- `POST /api/v1/auth/forgot-password` - Solicitar recuperação de senha
- `POST /api/v1/auth/reset-password` - Redefinir senha

### Clientes
- `GET /api/v1/clients` - Listar clientes
- `POST /api/v1/clients` - Criar cliente
- `GET /api/v1/clients/:id` - Buscar cliente
- `PUT /api/v1/clients/:id` - Atualizar cliente
- `DELETE /api/v1/clients/:id` - Remover cliente

### Tarefas
- `GET /api/v1/tasks` - Listar tarefas
- `POST /api/v1/tasks` - Criar tarefa
- `GET /api/v1/tasks/:id` - Buscar tarefa
- `PUT /api/v1/tasks/:id` - Atualizar tarefa
- `DELETE /api/v1/tasks/:id` - Remover tarefa

### Pagamentos
- `GET /api/v1/payments` - Listar pagamentos
- `POST /api/v1/payments` - Criar pagamento
- `GET /api/v1/payments/:id` - Buscar pagamento
- `PUT /api/v1/payments/:id` - Atualizar pagamento
- `DELETE /api/v1/payments/:id` - Remover pagamento

## Licença

MIT
