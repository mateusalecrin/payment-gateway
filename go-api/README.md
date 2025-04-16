# Go Gateway API

## Visão Geral

Esta aplicação é uma API Gateway desenvolvida em Go, seguindo princípios de Clean Architecture e Domain-Driven Design (DDD). Ela gerencia contas e faturas, com autenticação baseada em API Key, e utiliza PostgreSQL como banco de dados relacional.

---

## Sumário
- [Arquitetura](#arquitetura)
- [Principais Patterns e Princípios](#principais-patterns-e-princípios)
- [Módulos e Camadas](#módulos-e-camadas)
- [Endpoints](#endpoints)
- [Como rodar](#como-rodar)
- [Variáveis de ambiente](#variáveis-de-ambiente)

---

## Arquitetura

A aplicação é dividida em camadas, promovendo separação de responsabilidades e fácil manutenção:

- **Domain**: Entidades de negócio e interfaces de repositório (DDD).
- **DTO**: Objetos de transferência de dados entre camadas (DTO Pattern).
- **Repository**: Implementação da persistência de dados (Repository Pattern).
- **Service**: Lógica de negócio (Service Pattern).
- **Web**: Handlers HTTP, middlewares e configuração do servidor.
- **cmd/app**: Ponto de entrada da aplicação.

Fluxo principal:
1. Requisições chegam via HTTP (camada Web/Handlers).
2. Middlewares validam autenticação.
3. Handlers delegam para Services.
4. Services usam Repositories para persistência.
5. Dados trafegam entre camadas via DTOs.

---

## Principais Patterns e Princípios

- **Domain-Driven Design (DDD)**: Entidades ricas em regras de negócio, interfaces de repositório no domínio.
- **Repository Pattern**: Abstração da persistência, facilitando troca de banco e testes.
- **Service Pattern**: Lógica de negócio centralizada em serviços.
- **DTO Pattern**: Separação entre modelos de domínio e dados trafegados na API.
- **Injeção de Dependências**: Handlers e services recebem dependências via construtores.
- **Middlewares**: Autenticação via API Key.
- **RESTful**: Endpoints seguem boas práticas REST.

---

## Módulos e Camadas

```
internal/
  domain/        # Entidades e interfaces de domínio
  dto/           # Data Transfer Objects
  repository/    # Implementação dos repositórios
  service/       # Lógica de negócio
  web/
    handlers/    # Handlers HTTP (controllers)
    middlewares/ # Middlewares (ex: autenticação)
    server/      # Configuração do servidor HTTP
cmd/app/         # Ponto de entrada (main.go)
```

### Exemplos de entidades e interfaces (Domain)
- `Account`, `Invoice` (entidades)
- `AccountRepository`, `InvoiceRepository` (interfaces)

### DTOs
- `CreateAccountInput`, `AccountOutput`
- `CreateInvoiceInput`, `InvoiceOutput`

### Services
- `AccountService`: Criação de contas, consulta, atualização de saldo.
- `InvoiceService`: Criação e consulta de faturas.

### Repositories
- `AccountRepository`, `InvoiceRepository`: Implementam as interfaces do domínio usando PostgreSQL.

### Web
- **Handlers**: `AccountHandler`, `InvoiceHandler`
- **Middlewares**: `AuthMiddleware` (validação de API Key)
- **Server**: Configuração de rotas e inicialização do servidor HTTP (usando chi/v5)

---

## Endpoints

### Conta
- `POST /accounts` — Cria uma nova conta
- `GET /accounts` — Retorna dados da conta (requer header `X-API-Key`)

### Fatura
- `POST /invoice` — Cria uma nova fatura (requer header `X-API-Key`)
- `GET /invoice/{id}` — Busca uma fatura por ID (requer header `X-API-Key`)
- `GET /invoice` — Lista faturas da conta (requer header `X-API-Key`)

#### Exemplo de uso (test.http)
Veja o arquivo `test.http` para exemplos de requisições HTTP.

---

## Como rodar

1. **Pré-requisitos:**
   - Go 1.24+
   - Docker (para o banco de dados)
   - [golang-migrate](https://github.com/golang-migrate/migrate) (para migrações do banco de dados)
     - Instalação: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
   - Extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) (opcional, para testes de endpoints)

2. **Suba o banco de dados:**
   ```sh
docker-compose up -d
```

3. **Configure as variáveis de ambiente:**
   - Copie o arquivo `.env-example` para `.env`:
     ```sh
     cp .env-example .env
     ```
   - Edite o arquivo `.env` e preencha os valores conforme necessário.

4. **Instale as dependências:**
   ```sh
go mod tidy
```

5. **Rode a aplicação:**
   ```sh
go run cmd/app/main.go
```

A API estará disponível em `http://localhost:8080`.

---

## Variáveis de ambiente
Veja `.env` para todas as opções. Exemplo:

```
HTTP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_SSLMODE=disable
```

---

## Tecnologias
- Go
- PostgreSQL
- chi (roteador HTTP)
- Docker (banco de dados)

---

## Testes
Você pode usar o arquivo `test.http` para testar os endpoints com ferramentas como VSCode REST Client ou Insomnia.

---

## Autor
Projeto baseado no Full Cycle Imersão 22. 