# Micro Ecommerce

Micro Ecommerce is a research-based learning project for building a simple ecommerce app using a microservices architecture.

## Project Goals

- Learn and practice microservices architecture fundamentals (service separation, communication, deployment)
- Understand how a request flows through a distributed system (frontend → broker → services)
- Practice designing and implementing an authentication/authorization flow across services
- Learn service-to-service communication patterns (REST, message queues, etc.)
- Practice centralized logging and observability across multiple services
- Get hands-on experience with independent deployment and scaling of services
- Build a working, end-to-end reference project to revisit and extend as learning progresses

## Non-Goals (for now)

- Production-grade performance, security hardening, or scalability
- Full ecommerce feature set (payments, shipping, inventory management, etc.)
- Polished UI/UX — the frontend exists to demonstrate the flow, not to be a finished product

## Architecture Overview

```
┌──────────┐      ┌─────────┐      ┌──────────────────────────┐
│ Frontend │ ───> │ Broker  │ ───> │  Auth / Product / ...    │
└──────────┘      └─────────┘      └──────────────────────────┘
                        │
                        ▼
                    ┌────────┐
                    │  Log   │  <── errors & warnings from all services
                    └────────┘
```

## Services

| Service   | Description                                                        | Status      |
|-----------|---------------------------------------------------------------------|-------------|
| Frontend  | Website for users to access the app                                 | Planned     |
| Broker    | Handles incoming requests and routes them to the correct service    | Planned     |
| Auth      | Authentication service (login, register, tokens)                    | Planned     |
| Product   | Product service (listing, details, catalog)                         | Planned     |
| Log       | Stores error/warning logs from all services                         | Planned     |

## Tech Stack

> _To be decided — fill in as you choose languages/frameworks for each service._

- Frontend:
- Broker:
- Auth:
- Product:
- Log:
- Database(s):
- Messaging / Communication:
- Containerization / Orchestration:

## Project Structure

```
micro-ecommerce/
├── frontend/
├── broker/
├── auth/
├── product/
├── log/
├── docker-compose.yml
└── README.md
```

## Getting Started

> _To be filled in once the first services are scaffolded._

```bash
# clone
git clone <repo-url>
cd micro-ecommerce

# run (example, adjust once docker-compose is set up)
docker-compose up --build
```

## Learning Log

Use this section to track what you learn/decide as the project evolves.

- [ ] Define API contracts between broker and services
- [ ] Decide on auth strategy (JWT, session, OAuth?)
- [ ] Decide on inter-service communication (HTTP, gRPC, message queue?)
- [ ] Set up centralized logging format
- [ ] Add docker-compose for local dev

## License

This project is for personal learning purposes.