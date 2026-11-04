# Micro E-Commerce

A microservices-based e-commerce backend written in Go. Services run in Docker and communicate internally over Go's `net/rpc`.

## Architecture

```
                 ┌──────────────┐   net/rpc (TCP :1234)   ┌───────────────┐
 Client ──HTTP──▶│ api-gateway  │────────────────────────▶│  auth-service │
                 │    :8080     │                         │     :1234     │
                 └──────────────┘                         └───────┬───────┘
                                                                  │ GORM
                                                          ┌───────▼───────┐
                                                          │  MySQL 8.0    │
                                                          │     :3306     │
                                                          └───────────────┘
```

| Service        | Description                                              | Port (container) | Published to host |
|----------------|----------------------------------------------------------|------------------|-------------------|
| `api-gateway`  | Public HTTP entry point. Forwards requests via RPC.      | 8080             | Yes (`8080`)      |
| `auth-service` | User registration and login, exposed as an RPC service.  | 1234             | No (internal)     |
| `mysql`        | Database used by `auth-service`.                         | 3306             | Yes (`3306`)      |

## Tech Stack

- Go 1.26
- `net/rpc` (gob encoding, not gRPC)
- GORM with the MySQL driver
- bcrypt for password hashing
- Docker and Docker Compose

## Project Structure

```
micro-ecommerce/
├── api-gateway/
│   ├── cmd/server/main.go
│   └── Dockerfile
├── auth-service/
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── config/        # DB connection (reads env vars)
│   │   ├── model/         # GORM models
│   │   ├── repository/    # Data access layer
│   │   └── service/       # RPC methods
│   └── Dockerfile
├── docker-compose.yml
├── .env
└── README.md
```

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.26+ (only needed to run services outside Docker)

### Configuration

Create a `.env` file in the project root:

```env
MYSQL_ROOT_PASSWORD=rootpassword
MYSQL_DATABASE=micro_ecommerce
MYSQL_USER=user
MYSQL_PASSWORD=password
```

`docker-compose.yml` maps these to the variables each service reads:

| Variable      | Used by        | Description                          |
|---------------|----------------|--------------------------------------|
| `DB_HOST`     | `auth-service` | Database host (`mysql` in compose)   |
| `DB_PORT`     | `auth-service` | Database port (`3306`)               |
| `DB_USER`     | `auth-service` | Database user                        |
| `DB_PASSWORD` | `auth-service` | Database password                    |
| `DB_NAME`     | `auth-service` | Database name                        |
| `PORT`        | `auth-service` | RPC listen port (default `1234`)     |

> Never commit `.env` to version control. Make sure it is listed in `.gitignore`.

### Run with Docker Compose

```bash
docker compose up --build
```

Check that everything is up:

```bash
docker compose ps
docker compose logs auth-service
```

A healthy `auth-service` prints:

```
Database connection established
Database connection established and model auto-migrated
Auth RPC server listening on :1234
```

It is normal to see a few `connection refused` retries first while MySQL is starting.

Stop everything:

```bash
docker compose down        # keep data
docker compose down -v     # also delete the MySQL volume
```

### Run a service locally (without Docker)

Start only MySQL in Docker, then run the service from your machine. `DB_HOST` must be `localhost` here, not `mysql`.

```bash
docker compose up -d mysql

cd auth-service
export DB_HOST=localhost DB_PORT=3306 DB_USER=user DB_PASSWORD=password DB_NAME=micro_ecommerce
go run ./cmd/server
```

Go does not read `.env` files automatically, so the variables must be exported (or loaded with a tool such as `godotenv`).

## Auth Service RPC API

`auth-service` registers itself as `AuthService` and listens on TCP port `1234`.

| RPC method             | Args                              | Reply                               |
|------------------------|-----------------------------------|-------------------------------------|
| `AuthService.Register` | `RegisterArgs{Email, Password, Name}` | `RegisterReply{Success, Message, User}` |
| `AuthService.Login`    | `LoginArgs{Email, Password}`      | `LoginReply{Success, Message, User}`    |

Business errors (duplicate email, wrong password) are returned in `reply.Success` and `reply.Message`. The returned `error` is reserved for transport-level failures.

Passwords are hashed with bcrypt before storage and are never returned in replies.

### Calling from another Go service

```go
client, err := rpc.Dial("tcp", "auth-service:1234")
if err != nil {
	log.Fatal(err)
}
defer client.Close()

var reply service.RegisterReply
err = client.Call("AuthService.Register", &service.RegisterArgs{
	Email:    "user@example.com",
	Password: "secret",
	Name:     "Eric",
}, &reply)
```

The method name format is `"<RegisteredName>.<Method>"`. Both services must share the same args/reply struct definitions (for example, in a shared module), since gob encodes by field name and type.

## Troubleshooting

**`Cant read dsn` / `database configuration is not set properly`**
One or more `DB_*` environment variables are empty. Running `go run` from your terminal does not load `.env` or the compose `environment:` block. Export the variables or run through Docker Compose.

**`go.mod requires go >= X (running go Y)` during Docker build**
The Go version in the Dockerfile base image is older than the one in `go.mod`. Update `FROM golang:<version>-alpine` to match.

**`connection refused` to MySQL on startup**
MySQL takes several seconds to become ready. `auth-service` retries the connection automatically. For stricter ordering, add a MySQL healthcheck and use `depends_on` with `condition: service_healthy`.

**Cannot reach `auth-service` from the host**
This is intentional. It is only reachable inside the compose network at `auth-service:1234`. Add `ports: ["1234:1234"]` temporarily if you need to test it directly.

## Roadmap

- [ ] JWT issuing and validation in the gateway
- [ ] Gateway HTTP routes for register and login
- [ ] Product service
- [ ] Order service
- [ ] Shared module for RPC types
- [ ] MySQL healthcheck in compose
- [ ] TLS or authentication between services

## License

Add your license here.