<div align="center">
   <h1 style="display: inline-block; vertical-align: middle; font-size: 32px; font-weight: bold;">
    DEUNA Challenge
  </h1>
</div>

#### This project aims to develop an online payment platform, which will be an API-based application enabling e-commerce businesses to securely and seamlessly process transactions.

## Prerequisites
* Docker: Make sure you have Docker installed on your system.

## How to execute 
1. Clone this repository to your local machine.
2. Open a terminal in the project root.
3. Run the following command to build and start the services defined in the ``` docker-compose.yml ``` file:
```sh
make build
```
This will build the Docker images and start the containers.

## Stopping the Application
To stop the application and services, press ```Ctrl+C``` in the terminal where docker-compose up is running. Then, run the following command to stop and remove the containers:
```sh
make down
```

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