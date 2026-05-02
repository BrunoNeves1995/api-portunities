# API Portunities - Documentação

API REST para gerenciamento de oportunidades de trabalho (vagas de emprego).

## Estrutura do Projeto

```
api-portunities-aula3/
├── main.go              # Arquivo principal da aplicação
├── makefile             # Comandos de build, testes e documentação
├── go.mod               # Dependências do módulo Go
│
├── config/              # Configuração da aplicação
│   ├── config.go        # Configurações gerais
│   ├── logger.go        # Logger estruturado
│   └── sqlite.go        # Configuração do banco de dados SQLite
│
├── db/                  # Migrações e scripts do banco de dados
│
├── schemas/             # Modelos de dados
│   └── opening.go       # Struct Opening (Vaga de Emprego)
│
├── handler/             # Handlers HTTP (Controllers)
│   ├── handler.go       # Inicialização dos handlers
│   ├── createOpening.go # POST - Criar nova vaga
│   ├── getOpening.go    # GET - Obter vaga por ID
│   ├── deleteOpening.go # DELETE - Deletar vaga
│   ├── updateOpening.go # PUT - Atualizar vaga
│   ├── listOpening.go   # GET - Listar todas as vagas
│   ├── request.go       # Structs de requisição
│   ├── response.go      # Structs de resposta
│
├── router/              # Configuração de rotas
│   ├── router.go        # Inicialização do Gin router
│   └── routes.go        # Definição das rotas da API
│
└── docs/                # Documentação Swagger
    ├── docs.go          # Definições geradas pelo Swagger
    ├── swagger.json     # Especificação OpenAPI em JSON
    └── swagger.yaml     # Especificação OpenAPI em YAML
```

## Descrição das Pastas

### 📁 `config/`
Gerencia toda a configuração da aplicação:
- **config.go**: Inicializa o banco de dados
- **logger.go**: Logger estruturado com níveis de severidade
- **sqlite.go**: Configura conexão com SQLite

### 📁 `db/`
Armazena scripts de migração e inicialização do banco de dados (se necessário).

### 📁 `schemas/`
Define os modelos de dados da aplicação:
- **opening.go**: Struct `Opening` com campos: Role, Company, Location, Remote, Link, Salary

### 📁 `handler/`
Controllers HTTP que processam as requisições:
- **createOpening.go**: POST `/api/v1/opening` - Criar nova vaga
- **getOpening.go**: GET `/api/v1/opening?id=X` - Obter vaga específica
- **listOpening.go**: GET `/api/v1/openings` - Listar todas as vagas
- **updateOpening.go**: PUT `/api/v1/opening?id=X` - Atualizar vaga
- **deleteOpening.go**: DELETE `/api/v1/opening?id=X` - Deletar vaga
- **request.go**: Structs de validação de entrada
- **response.go**: Structs de resposta padronizadas

### 📁 `router/`
Configura as rotas da API:
- **router.go**: Inicializa o engine Gin
- **routes.go**: Define endpoints e monta o Swagger

### 📁 `docs/`
Gerada automaticamente pelo Swagger (swag):
- **docs.go**: Definições em Go
- **swagger.json**: Especificação em JSON
- **swagger.yaml**: Especificação em YAML

## Endpoints da API

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| POST | `/api/v1/opening` | Criar nova vaga |
| GET | `/api/v1/opening?id=X` | Obter vaga por ID |
| GET | `/api/v1/openings` | Listar todas as vagas |
| PUT | `/api/v1/opening?id=X` | Atualizar vaga |
| DELETE | `/api/v1/opening?id=X` | Deletar vaga |
| GET | `/swagger/*any` | UI Swagger (documentação interativa) |

## Executar a Aplicação

```bash
# Instalar dependências
go mod download

# Rodar com documentação Swagger
make run-with-docs

# Ou simplesmente
make

# Gerar apenas documentação
make docs

# Build
make build

# Testes
make test

# Lint
make lint

# Formatar código
make fmt

# Limpar
make clean
```

## Acessar a API

- **Swagger UI**: http://localhost:8000/swagger/index.html
- **API Base**: http://localhost:8000/api/v1

## Exemplo de Requisição

```bash
# Criar vaga
curl -X POST http://localhost:8000/api/v1/opening \
  -H "Content-Type: application/json" \
  -d '{
    "role": "Senior Developer",
    "company": "Tech Company",
    "location": "São Paulo, Brazil",
    "remote": true,
    "link": "https://company.com/careers",
    "salary": 15000
  }'
```

## Tecnologias

- **Go 1.20+**
- **Gin** - Framework Web
- **GORM** - ORM para Go
- **SQLite** - Banco de dados
- **Swagger/OpenAPI** - Documentação da API
