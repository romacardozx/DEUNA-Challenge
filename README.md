# DEUNA-Challenge

#### Online Payment Platform
* This project aims to develop an online payment platform, which will be an API-based application enabling e-commerce businesses to securely and seamlessly process transactions.

## New proposal for project structure with Cloud
```sh
DEUNA-CHALLENGE/
├── cmd/
│   └── api/
│       └── main.go                 # Punto de entrada principal de la aplicación
├── internal/
│   ├── app/
│   │   └── app.go                  # Configuración y inicialización de la aplicación
│   ├── core/
│   │   ├── models/
│   │   │   ├── payment.go          # Definición de la entidad Payment
│   │   │   ├── merchant.go         # Definición de la entidad Merchant
│   │   │   └── customer.go         # Definición de la entidad Customer
│   │   ├── repositories/
│   │   │   ├── payment_repo.go     # Interfaz del repositorio de Payment
│   │   │   └── merchant_repo.go    # Interfaz del repositorio de Merchant
│   │   ├── services/
│   │   │   ├── payment_service.go  # Lógica de negocio para pagos
│   │   │   ├── refund_service.go   # Lógica de negocio para reembolsos
│   │   │   └── bank_simulator.go   # Simulador del banco
│   │   └── queries/
│   │       ├── payment_queries.go  # Consultas SQL complejas para pagos
│   │       └── merchant_queries.go # Consultas SQL complejas para comerciantes
│   ├── handlers/
│   │   ├── v1/
│   │   │   ├── payment_handler.go  # Manejador HTTP para pagos
│   │   │   └── refund_handler.go   # Manejador HTTP para reembolsos
│   │   ├── middleware/
│   │   │   ├── auth.go             # Middleware de autenticación
│   │   │   └── logging.go          # Middleware de logging
│   │   └── routes.go               # Definición de rutas de la API
│   ├── database/
│   │   ├── migrations/
│   │   │   ├── 001_create_payments_table.sql
│   │   │   └── 002_create_merchants_table.sql
│   │   └── db.go                   # Configuración y conexión a la base de datos
│   └── cloud/
│       ├── aws/
│       │   ├── dynamodb/
│       │   │   └── repository.go   # Implementación de repositorio con DynamoDB
│       │   ├── lambda/
│       │   │   └── handler.go      # Adaptador para AWS Lambda
│       │   └── sqs/
│       │       └── queue.go        # Implementación de cola con SQS
│       └── interfaces/
│           ├── database.go         # Interfaces para abstracción de base de datos
│           └── queue.go            # Interfaces para abstracción de cola
├── pkg/
│   ├── logger/
│   │   └── logger.go               # Implementación de logging personalizado
│   └── errors/
│       └── errors.go               # Manejo de errores personalizado
├── config/
│   └── config.go                   # Configuraciones de la aplicación
├── deployments/
│   ├── docker/
│   │   ├── Dockerfile
│   │   └── docker-compose.yml
│   └── aws/
│       ├── cloudformation/
│       │   └── stack.yaml          # Template de CloudFormation para infraestructura AWS
│       └── terraform/
│           └── main.tf             # Configuración de Terraform (alternativa a CloudFormation)
├── .dockerignore
├── go.mod
├── go.sum
└── README.md
```