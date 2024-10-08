# DEUNA-Challenge

#### Online Payment Platform
* This project aims to develop an online payment platform, which will be an API-based application enabling e-commerce businesses to securely and seamlessly process transactions.

## Project Structure
```sh
DEUNA-CHALLENGE/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── core/
│   │   ├── models/
│   │   │   ├── payment.go
│   │   │   ├── merchant.go
│   │   │   ├── bank.go
│   │   │   ├── refund.go
│   │   │   └── customer.go
│   │   ├── repositories/
│   │   │   ├── payment_repo.go
│   │   │   └── merchant_repo.go
│   │   └── services/
│   │       ├── payment_service.go
│   │       ├── refund_service.go
│   │       ├── bank_simulator.go
│   │       └── bank_simulator_test.go
│   ├── handlers/
│   │   ├── v1/
│   │   │   ├── payment_handler.go
│   │   │   └── refund_handler.go
│   │   ├── middleware/
│   │   │   ├── auth.go
│   │   │   └── logging.go
│   │   └── routes.go
│   └── database/
│       └── db.go
├── pkg/
│   ├── logger/
│   │   └── logger.go
│   └── errors/
│       └── errors.go
├── config/
│   └── config.go
├── deployments/
│   └── docker/
│       ├── init.sql
│       ├── Dockerfile
│       └── docker-compose.yml
├── .dockerignore
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```