# DEUNA-Challenge

#### Online Payment Platform
* This project aims to develop an online payment platform, which will be an API-based application enabling e-commerce businesses to securely and seamlessly process transactions.

## Initial proposal for project structure
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
│   └── database/
│       ├── migrations/
│       │   ├── 001_create_payments_table.sql
│       │   └── 002_create_merchants_table.sql
│       └── db.go                   # Configuración y conexión a la base de datos
├── pkg/
│   ├── logger/
│   │   └── logger.go               # Implementación de logging personalizado
│   └── errors/
│       └── errors.go               # Manejo de errores personalizado
├── config/
│   └── config.go                   # Configuraciones de la aplicación
├── .air.toml                       # Configuración de Air para recarga en caliente
├── go.mod
├── go.sum
└── README.md
```