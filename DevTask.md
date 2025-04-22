# Histórias de Usuário - Mini CRM de Freelancers

## Autenticação e Conta de Usuário

### US-01: Registro de Usuário

**Como** um freelancer,

**Quero** poder me registrar no sistema,

**Para que** eu possa criar minha conta e começar a usar o CRM.

**Critérios de Aceitação:**

* Formulário com campos: nome, e-mail, senha e confirmação de senha
* Validação de e-mail único no sistema
* Validação de força de senha
* Confirmação de e-mail por link enviado
* Redirecionamento para o dashboard após registro bem-sucedido

### US-02: Login de Usuário

**Como** um freelancer registrado,

**Quero** fazer login no sistema,

**Para que** eu possa acessar meus dados e gerenciar meus clientes e projetos.

**Critérios de Aceitação:**

* Formulário com campos de e-mail e senha
* Validação de credenciais corretas
* Opção "Lembrar-me"
* Opção "Esqueci minha senha"
* Redirecionamento para o dashboard após login bem-sucedido

### US-03: Recuperação de Senha

**Como** um freelancer que esqueceu a senha,

**Quero** poder redefini-la,

**Para que** eu possa recuperar o acesso à minha conta.

**Critérios de Aceitação:**

* Formulário para informar e-mail
* E-mail enviado com link de redefinição de senha
* Link válido por 24 horas
* Formulário para criar nova senha
* Feedback de sucesso após redefinição

### US-04: Edição de Perfil

**Como** um freelancer logado,

**Quero** poder editar meu perfil,

**Para que** eu possa manter minhas informações atualizadas.

**Critérios de Aceitação:**

* Formulário com campos editáveis: nome, e-mail, telefone, foto, etc.
* Opção para alterar senha
* Validação de dados
* Feedback de sucesso após salvar alterações

### US-05: Gerenciamento de Plano

**Como** um freelancer,

**Quero** poder visualizar e alterar meu plano de assinatura,

**Para que** eu possa escolher o nível de serviço adequado às minhas necessidades.

**Critérios de Aceitação:**

* Visualização do plano atual (gratuito ou pago)
* Informações sobre limite de clientes e recursos disponíveis
* Opção para upgrade para plano pago
* Processo de pagamento seguro
* Histórico de pagamentos e faturas

## Gerenciamento de Clientes

### US-06: Cadastro de Cliente

**Como** um freelancer,

**Quero** cadastrar um novo cliente no sistema,

**Para que** eu possa organizar os trabalhos relacionados a ele.

**Critérios de Aceitação:**

* Formulário com campos: nome, e-mail, telefone, empresa, endereço, etc.
* Validação de dados
* Upload de logo/foto do cliente (opcional)
* Indicação se o cliente está ativo ou inativo
* Aviso quando atingir o limite de 3 clientes ativos no plano gratuito

### US-07: Listagem de Clientes

**Como** um freelancer,

**Quero** visualizar a lista de todos os meus clientes,

**Para que** eu possa ter uma visão geral e acessar rapidamente suas informações.

**Critérios de Aceitação:**

* Lista de clientes com informações básicas (nome, empresa, status)
* Ordenação por nome, data de adição, status
* Filtros por status (ativo/inativo)
* Busca por nome ou empresa
* Indicação visual de clientes com projetos em andamento

### US-08: Visualização de Detalhes do Cliente

**Como** um freelancer,

**Quero** visualizar todos os detalhes de um cliente específico,

**Para que** eu possa ter acesso completo às suas informações e projetos relacionados.

**Critérios de Aceitação:**

* Exibição de todas as informações do cliente
* Lista de projetos/tarefas associados ao cliente
* Histórico de pagamentos
* Notas e observações
* Opções para editar ou arquivar o cliente

### US-09: Edição de Cliente

**Como** um freelancer,

**Quero** editar as informações de um cliente existente,

**Para que** eu possa manter os dados atualizados.

**Critérios de Aceitação:**

* Formulário preenchido com dados atuais
* Todos os campos editáveis
* Validação de dados
* Feedback de sucesso após salvar alterações

### US-10: Arquivamento/Desativação de Cliente

**Como** um freelancer,

**Quero** arquivar ou desativar um cliente,

**Para que** eu possa manter minha lista organizada e dentro do limite do plano gratuito.

**Critérios de Aceitação:**

* Opção para arquivar/desativar cliente
* Confirmação antes de arquivar
* Cliente arquivado não conta no limite de clientes ativos
* Possibilidade de reativar cliente arquivado
* Todos os dados do cliente preservados

## Gerenciamento de Tarefas/Projetos

### US-11: Criação de Tarefa

**Como** um freelancer,

**Quero** criar uma nova tarefa associada a um cliente,

**Para que** eu possa acompanhar os trabalhos que preciso realizar.

**Critérios de Aceitação:**

* Formulário com campos: título, descrição, cliente associado, prazo, prioridade, valor
* Seleção de status inicial (ex: orçamento, em andamento, revisão, etc.)
* Campo para estimativa de horas
* Anexo de arquivos relevantes
* Tarefa criada visível no dashboard e na lista de tarefas

### US-12: Listagem de Tarefas

**Como** um freelancer,

**Quero** visualizar todas as minhas tarefas,

**Para que** eu possa ter uma visão geral do meu trabalho.

**Critérios de Aceitação:**

* Lista de tarefas com informações essenciais
* Filtros por status, cliente, prazo
* Ordenação por prazo, prioridade, cliente
* Indicadores visuais de tarefas atrasadas ou próximas do prazo
* Visualização em formato de lista ou kanban

### US-13: Visualização de Detalhes da Tarefa

**Como** um freelancer,

**Quero** visualizar todos os detalhes de uma tarefa específica,

**Para que** eu possa ter acesso completo às informações relevantes.

**Critérios de Aceitação:**

* Exibição de todas as informações da tarefa
* Histórico de alterações de status
* Arquivos anexados
* Notas e comentários
* Informações de pagamento relacionado

### US-14: Atualização de Status da Tarefa

**Como** um freelancer,

**Quero** atualizar o status das minhas tarefas,

**Para que** eu possa acompanhar o progresso do trabalho.

**Critérios de Aceitação:**

* Interface intuitiva para mudança de status (drag-and-drop no kanban)
* Lista de status pré-definidos (orçamento, aprovado, em andamento, revisão, finalizado)
* Atualização automática de data de cada mudança de status
* Notificação opcional para o cliente sobre mudança de status (premium)

### US-15: Edição de Tarefa

**Como** um freelancer,

**Quero** editar as informações de uma tarefa existente,

**Para que** eu possa manter os dados atualizados.

**Critérios de Aceitação:**

* Formulário preenchido com dados atuais
* Todos os campos editáveis
* Histórico de alterações (premium)
* Validação de dados
* Feedback de sucesso após salvar alterações

### US-16: Conclusão de Tarefa

**Como** um freelancer,

**Quero** marcar uma tarefa como concluída,

**Para que** eu possa acompanhar meu trabalho finalizado.

**Critérios de Aceitação:**

* Opção para marcar como concluído
* Preenchimento opcional de informações de conclusão (data, comentários)
* Notificação opcional para o cliente (premium)
* Tarefa movida para lista de concluídos
* Atualização do status financeiro relacionado à tarefa

## Gerenciamento de Pagamentos

### US-17: Registro de Pagamento

**Como** um freelancer,

**Quero** registrar pagamentos recebidos dos clientes,

**Para que** eu possa controlar minhas finanças.

**Critérios de Aceitação:**

* Formulário com campos: valor, data, método de pagamento, cliente, tarefa relacionada
* Opção para pagamento parcial ou total
* Upload de comprovante (opcional)
* Atualização automática do status financeiro da tarefa relacionada
* Registro no histórico financeiro

### US-18: Listagem de Pagamentos

**Como** um freelancer,

**Quero** visualizar todos os pagamentos recebidos,

**Para que** eu possa ter um controle financeiro do meu trabalho.

**Critérios de Aceitação:**

* Lista de pagamentos com informações essenciais
* Filtros por cliente, período, status
* Totalizadores por período
* Indicação visual de pagamentos pendentes
* Exportação da lista (premium)

### US-19: Geração de Relatório Financeiro

**Como** um usuário premium,

**Quero** gerar relatórios financeiros,

**Para que** eu possa analisar meus ganhos e planejar meu negócio.

**Critérios de Aceitação:**

* Seleção de período para o relatório
* Filtros por cliente, tipo de projeto
* Gráficos e visualizações das informações
* Opções de exportação (PDF, CSV)
* Cálculos de médias, totais e tendências

### US-20: Registro de Despesa

**Como** um usuário premium,

**Quero** registrar despesas relacionadas a projetos,

**Para que** eu possa calcular o lucro real.

**Critérios de Aceitação:**

* Formulário com campos: valor, data, descrição, categoria, projeto relacionado
* Upload de comprovante (opcional)
* Despesa incluída nos cálculos de relatórios
* Visualização de despesas por projeto

## Dashboard e Visualizações

### US-21: Visualização do Dashboard

**Como** um freelancer,

**Quero** acessar um dashboard com informações relevantes,

**Para que** eu possa ter uma visão geral do meu trabalho e finanças.

**Critérios de Aceitação:**

* Resumo de tarefas por status
* Lista de tarefas com prazo próximo ou atrasadas
* Resumo financeiro do mês atual
* Atalhos para ações comuns
* Gráficos de desempenho (premium)

### US-22: Calendário de Tarefas

**Como** um freelancer,

**Quero** visualizar minhas tarefas em um calendário,

**Para que** eu possa planejar melhor meu tempo.

**Critérios de Aceitação:**

* Visualização em formato de calendário mensal/semanal
* Tarefas exibidas nas datas de prazo
* Diferentes cores por cliente ou status
* Possibilidade de arrastar e soltar para ajustar prazos
* Visualização de disponibilidade

### US-23: Configuração de Lembretes

**Como** um usuário premium,

**Quero** configurar lembretes para prazos e pagamentos,

**Para que** eu não perca datas importantes.

**Critérios de Aceitação:**

* Configuração de lembretes por e-mail
* Definição de antecedência para notificações
* Tipos de lembrete: prazos, pagamentos pendentes, follow-up
* Personalização de mensagens
* Histórico de lembretes enviados

## Integrações

### US-24: Integração com Google Calendar

**Como** um usuário premium,

**Quero** sincronizar minhas tarefas com o Google Calendar,

**Para que** eu possa ter todos os compromissos em um só lugar.

**Critérios de Aceitação:**

* Autorização da conta Google
* Seleção de calendário para sincronização
* Opções de sincronização bidirecional ou unidirecional
* Configuração de quais tarefas sincronizar (por status, cliente)
* Atualização automática quando houver mudanças

### US-25: Integração com Notion

**Como** um usuário premium,

**Quero** sincronizar minhas tarefas com o Notion,

**Para que** eu possa usar minhas ferramentas preferidas junto com o CRM.

**Critérios de Aceitação:**

* Autorização da conta Notion
* Seleção de database/página para sincronização
* Mapeamento de campos entre os sistemas
* Configuração de frequência de sincronização
* Logs de sincronização

### US-26: Exportação de Dados

**Como** um usuário premium,

**Quero** exportar meus dados em diferentes formatos,

**Para que** eu possa usar em outras ferramentas ou manter backups.

**Critérios de Aceitação:**

* Exportação em formatos CSV, Excel e PDF
* Seleção de dados a exportar (clientes, tarefas, pagamentos)
* Filtros para exportação parcial
* Download direto ou envio por e-mail
* Agendamento de exportações periódicas

## Gerenciamento de Plano e Pagamentos do Sistema

### US-27: Upgrade para Plano Premium

**Como** um usuário do plano gratuito,

**Quero** fazer upgrade para o plano premium,

**Para que** eu possa acessar todos os recursos avançados.

**Critérios de Aceitação:**

* Visualização de comparação entre planos
* Seleção de periodicidade (mensal/anual)
* Processo de pagamento seguro
* Ativação imediata dos recursos premium
* E-mail de confirmação com detalhes da assinatura

### US-28: Gerenciamento de Assinatura

**Como** um usuário premium,

**Quero** gerenciar minha assinatura,

**Para que** eu possa alterar forma de pagamento ou cancelar quando necessário.

**Critérios de Aceitação:**

* Visualização do status da assinatura atual
* Histórico de pagamentos
* Opção para atualizar método de pagamento
* Opção para cancelar assinatura
* Informações sobre próxima cobrança


# Mini CRM de Freelancers - Planejamento Técnico para MVP

Vou criar um planejamento técnico completo para um MVP de CRM para freelancers usando Go no backend e Nuxt no frontend, conforme solicitado. Vou apresentar a estrutura de pastas, arquivos e o backlog técnico do projeto sem incluir código.

## Visão Geral da Arquitetura

O sistema será construído com:

* **Backend** : Go (Golang) com Gin framework
* **Frontend** : Nuxt.js 3 (Vue.js)
* **Banco de Dados** : PostgreSQL
* **Autenticação** : JWT
* **Integrações** : Google Calendar API, Notion API

## Estrutura de Pastas do Projeto

### Backend (Go)

```
backend/
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada da aplicação
├── configs/                 # Configurações da aplicação
│   ├── config.go
│   └── env.go
├── internal/
│   ├── api/                 # Handlers da API
│   │   ├── auth.go
│   │   ├── clients.go
│   │   ├── payments.go
│   │   ├── tasks.go
│   │   └── user.go
│   ├── middleware/          # Middlewares
│   │   ├── auth.go
│   │   └── cors.go
│   ├── models/              # Modelos de dados
│   │   ├── client.go
│   │   ├── payment.go
│   │   ├── task.go
│   │   └── user.go
│   ├── repository/          # Camada de acesso a dados
│   │   ├── client_repo.go
│   │   ├── payment_repo.go
│   │   ├── task_repo.go
│   │   └── user_repo.go
│   └── services/            # Lógica de negócios
│       ├── auth_service.go
│       ├── client_service.go
│       ├── integration/
│       │   ├── gcalendar.go
│       │   └── notion.go
│       ├── payment_service.go
│       └── task_service.go
├── migrations/              # Migrações do banco de dados
│   ├── 000001_create_users_table.up.sql
│   ├── 000001_create_users_table.down.sql
│   ├── 000002_create_clients_table.up.sql
│   └── ...
├── pkg/                     # Código compartilhável
│   ├── logger/
│   │   └── logger.go
│   └── util/
│       └── datetime.go
├── .env.example             # Exemplo de arquivo de variáveis de ambiente
├── Dockerfile               # Configuração para Docker
├── go.mod                   # Dependências do Go
└── go.sum
```

### Frontend (Nuxt.js)

```
frontend/
├── assets/                  # Arquivos estáticos (CSS, imagens)
│   ├── css/
│   │   └── main.css
│   └── images/
├── components/              # Componentes Vue reutilizáveis
│   ├── auth/
│   │   ├── LoginForm.vue
│   │   └── RegisterForm.vue
│   ├── clients/
│   │   ├── ClientCard.vue
│   │   ├── ClientForm.vue
│   │   └── ClientList.vue
│   ├── dashboard/
│   │   ├── DashboardStats.vue
│   │   └── RecentActivity.vue
│   ├── payments/
│   │   ├── PaymentForm.vue
│   │   └── PaymentList.vue
│   ├── tasks/
│   │   ├── TaskForm.vue
│   │   └── TaskList.vue
│   └── ui/
│       ├── Button.vue
│       ├── Card.vue
│       └── Modal.vue
├── layouts/                 # Layouts da aplicação
│   ├── default.vue
│   └── auth.vue
├── middleware/              # Middlewares do Nuxt
│   └── auth.ts
├── pages/                   # Páginas da aplicação
│   ├── index.vue
│   ├── auth/
│   │   ├── login.vue
│   │   └── register.vue
│   ├── clients/
│   │   ├── index.vue
│   │   ├── [id].vue
│   │   └── new.vue
│   ├── payments/
│   │   ├── index.vue
│   │   └── [id].vue
│   └── tasks/
│       ├── index.vue
│       └── [id].vue
├── plugins/                 # Plugins do Nuxt
│   └── api.ts
├── public/                  # Arquivos públicos
│   ├── favicon.ico
│   └── robots.txt
├── server/                  # API Server do Nuxt
│   └── api/
├── store/                   # Estado da aplicação (Pinia)
│   ├── auth.ts
│   ├── client.ts
│   ├── payment.ts
│   └── task.ts
├── types/                   # Tipos TypeScript
│   ├── client.ts
│   ├── payment.ts
│   └── task.ts
├── utils/                   # Funções utilitárias
│   ├── date.ts
│   └── validation.ts
├── .env.example             # Exemplo de variáveis de ambiente
├── .gitignore
├── app.vue                  # Componente raiz
├── nuxt.config.ts           # Configuração do Nuxt
├── package.json             # Dependências do projeto
├── tailwind.config.js       # Configuração do Tailwind CSS
└── tsconfig.json            # Configuração do TypeScript
```

## Banco de Dados - Modelo

### Tabelas Principais:

1. **users**
2. **clients**
3. **tasks**
4. **payments**
5. **integrations**

## Backlog Técnico para o MVP

### Sprint 1: Setup e Infraestrutura Básica

#### Backend:

1. Configurar projeto Go com estrutura limpa
2. Configurar banco de dados PostgreSQL e migrações
3. Implementar sistema de autenticação JWT
4. Criar modelo de usuário e endpoints de cadastro/login

#### Frontend:

1. Inicializar projeto Nuxt.js 3
2. Configurar Tailwind CSS para estilização
3. Criar páginas de login e registro
4. Implementar gerenciamento de estados com Pinia
5. Criar layout básico da aplicação

### Sprint 2: Gerenciamento de Clientes

#### Backend:

1. Criar modelo de Cliente
2. Implementar CRUD de Clientes
3. Adicionar validações
4. Implementar limite de 3 clientes ativos para contas gratuitas

#### Frontend:

1. Criar páginas de listagem, visualização e criação de clientes
2. Implementar formulários de cliente com validação
3. Criar componentes para exibição de detalhes do cliente
4. Adicionar indicador de limite para usuários gratuitos

### Sprint 3: Gerenciamento de Tarefas

#### Backend:

1. Criar modelo de Tarefas
2. Implementar CRUD de Tarefas
3. Adicionar relacionamento com Clientes
4. Implementar filtros por status, cliente e data

#### Frontend:

1. Criar páginas de listagem e gerenciamento de tarefas
2. Implementar filtros na interface
3. Criar componentes para exibição de status da tarefa
4. Adicionar drag-and-drop para atualização de status

### Sprint 4: Gerenciamento de Pagamentos

#### Backend:

1. Criar modelo de Pagamentos
2. Implementar CRUD de Pagamentos
3. Relacionar pagamentos com clientes e tarefas
4. Implementar cálculos de valores e status de pagamento

#### Frontend:

1. Criar páginas de listagem e visualização de pagamentos
2. Implementar formulários de registros de pagamentos
3. Criar componentes para exibição de status de pagamento
4. Implementar dashboards simples com informações financeiras

### Sprint 5: Recursos Premium e Planos

#### Backend:

1. Implementar sistema de planos (gratuito/pago)
2. Criar funcionalidade de exportação de dados (CSV/PDF)
3. Implementar sistema de lembretes
4. Criar endpoints para relatórios

#### Frontend:

1. Criar página de planos e assinatura
2. Implementar funcionalidade de exportação na interface
3. Criar componentes para configuração de lembretes
4. Adicionar visualização de relatórios

### Sprint 6: Integrações

#### Backend:

1. Implementar integração com Google Calendar API
2. Implementar integração com Notion API
3. Criar endpoints para sincronização de dados
4. Implementar armazenamento de dados de integração

#### Frontend:

1. Criar páginas de configuração de integrações
2. Implementar fluxo de autorização OAuth
3. Criar componentes para visualizar status de sincronização
4. Adicionar configurações de mapeamento de dados

### Sprint 7: Testes e Polimento

#### Backend:

1. Implementar testes unitários
2. Adicionar testes de integração
3. Otimizar consultas ao banco de dados
4. Configurar logging e monitoramento

#### Frontend:

1. Implementar testes de componentes
2. Otimizar performance e carregamento
3. Polir interface para responsividade
4. Implementar feedback de usuário e tratamento de erros

### Sprint 8: Implantação e Lançamento

1. Configurar ambiente de produção
2. Implementar CI/CD
3. Criar scripts de backup
4. Configurar monitoramento e alertas
5. Preparar documentação para usuários finais

## Requisitos Técnicos

### Backend:

* Go 1.21+
* Gin Web Framework
* GORM para ORM
* JWT para autenticação
* PostgreSQL 15+
* Docker para containerização

### Frontend:

* Nuxt.js 3
* Vue.js 3
* TypeScript
* Tailwind CSS
* Pinia para gerenciamento de estado
* Vitest para testes

## Considerações de Segurança

1. Implementar autenticação JWT com refresh tokens
2. Sanitizar todas as entradas do usuário
3. Usar HTTPS para todas as comunicações
4. Implementar rate limiting
5. Armazenar senhas com hash seguro (bcrypt)
6. Seguir princípios OWASP para proteção contra vulnerabilidades comuns

## Considerações de Escalabilidade

1. Implementar cache para melhorar performance
2. Projetar API com paginação
3. Utilizar filas para processamento assíncrono (como lembretes e relatórios)
4. Preparar para possível implementação de microserviços no futuro

Este planejamento técnico fornece uma estrutura sólida para o desenvolvimento do MVP do CRM para freelancers, com uma arquitetura moderna usando Go e Nuxt.js, seguindo boas práticas de desenvolvimento e com um caminho claro para a implementação das funcionalidades principais.
