<div align="center">
   <h1 style="display: inline-block; vertical-align: middle; font-size: 32px; font-weight: bold;">
    DEUNA Challenge
  </h1>
</div>

#### This project aims to develop an online payment platform, which will be an API-based application enabling e-commerce businesses to securely and seamlessly process transactions.

## Entities Involved:
1. Customer: Individuals who make online purchases and complete payments through the platform.
2. Merchant: The seller who utilizes the payment platform to receive payments from customers.
3. Online Payment Platform: An application that validates requests, stores card information, and manages payment
   requests and responses to and from the acquiring bank.
4. Acquiring Bank: Facilitates the actual retrieval of funds from the customer's card and transfers them to the merchant.
   Additionally, it validates card information and sends payment details to the relevant processing organization.

## Requirements:
The requirements for this initial phase are as follows:
1. Payment Processing:
   The online payment platform should provide merchants with the ability to process a payment and receive either a
   successful or unsuccessful response.
2. Querying Details of Previous Payments:
   Merchants should be able to retrieve details of previously made payments using a unique payment identifier.
3. Bank Simulation:
   Utilize a bank simulator to simulate the interaction with the acquiring bank.

## Deliverables:
1. Build an API that allows merchants:
   a. To process a payment through the online payment platform.
   b. To retrieve details of a previously made payment.
   c. To process refunds for specific transactions.
2. Integrate a Bank Simulator:
   Use a bank simulator to test and simulate responses from the acquiring bank.

## Considerations:
- Execution of the Solution:
    - Provide clear instructions for setting up and running the online payment platform API.
    - Specify any dependencies or prerequisites for the solution.
- Assumptions:
  - Clarify any assumptions made during the design and implementation.
  - Areas for Improvement:
  - Identify potential areas for improvement in the online payment platform.
  - Discuss any design decisions or trade-offs made during development.
- Cloud Technologies:
  - Specify the cloud technologies used and justify the choice.

## Extra:
- Authentication and Security:
  - Implement measures for authentication and security to ensure secure transactions.
- Audit Trail:
    - Include an audit trail feature to track activities such as payment processing, queries for payment details, and
  refunds.

## Prerequisites
* Docker: Make sure you have Docker installed on your system.

## How to execute 
1. Clone this repository to your local machine.
2. Open a terminal in the project root and run.
```sh
go mod tidy 
go mod vendor
```
3. Run the following command to build and start the services defined in the ``` docker-compose.yml ``` file:
```sh
make build
```

This will build the Docker images and start the containers.

4. if you have already done the build you can use the following command to run the app
```sh
make run
```

## Stopping the Application
To stop the application and services, press ```Ctrl+C``` in the terminal where docker-compose up is running. Then, run the following command to stop and remove the containers:
```sh
make down
```

## Run Tests
If you want to run test use the following command to execute it:
```sh
make test
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

## API Endpoints

### Payment Routes

#### POST /api/v1/payment: Process a payment
* Request Body:
```sh
    {
    "Customer_id": "c1",
    "merchant_id": "m1",
    "amount": 110.50,
    "currency": "USD",
    "description": "Payment for product XYZ"
    }
```

#### GET /api/v1/payment/details/:id: Get payment details
* Parameters:
```c
    * id: Payment ID //Take some id from get payments by merchantID to test it.
```

#### GET /api/v1/merchant/:merchantId/payments: List payments for a merchant
* Parameters:
```c
    merchantId: Merchant ID //MerchantsIDs: (Merchant 1: m1) and (Merchant 2: m2) defined in init.sql migration.
```
* QueryParams:
```c
    // If you want you can use it, otherwise the api has a default limit 10 and offset 0.
    limit: 20 //OPTIONAL
    offset: 0 //OPTIONAL
```

### Refund Routes

#### POST /api/v1/refund: Process a refund
* Request Body:
```sh
    {
        "PaymentID": "string", 
        "Reason": "string"
    }
```

#### GET /api/v1/refund/details/:refundId: Get refund details
* Parameters:
```sh
    * refundId: Refund ID
```
#### Note: All routes are protected by the AuthMiddleware().
* Header: 
```c
    * Authorization: token //You can use the default token for default merchants defined in init.sql migration: 
    //For Merchant 1: token1
    //For Merchant 2: token2
```

### Assumptions
- I'm always assuming that the bank always will be able to refund a payment and to do a reversal
- I'm assuming that the bank will deposit the money into the merchant account and withdraw it from the customer account
- I'm assuming that we're communicating through a secure network + authenticated users + encryption

# Why PostgreSQL, Go, and Docker for this Online Payment Platform

## PostgreSQL
- Ensures data integrity and ACID compliance, crucial for financial transactions
- Scalable and supports advanced features like JSON data types
- Robust for handling complex financial queries and reporting

## Go (Golang)
- High performance and efficient concurrency, ideal for processing multiple transactions
- Static typing and compilation reduce runtime errors, enhancing platform stability
- Excellent standard library with built-in support for cryptography and networking
- Compiles to a single binary, simplifying deployment

## Docker
- Ensures consistency across development, testing, and production environments
- Improves security through isolation of components
- Facilitates easy scaling and deployment, crucial for a growing payment platform
- Simplifies dependency management and application portability

This combination provides a solid foundation for building a secure, scalable, and efficient online payment platform. PostgreSQL offers the reliability for financial data, Go provides the performance for transaction processing, and Docker ensures consistent and scalable deployment across different environments.

## Additional Considerations/Recomendations
#### If you want to visualize the database in visual studio code, I recommend using the extension “Database Client”.
![alt text](image.png)
*  When you install the extension, the database client icon will appear in the vs code sidebar.
    1. Create a new connection
    2. Select PostgreSQL and complete the fields (user, password, DB) with the following credentials defined in the docker compose.
        ```sh 
            - POSTGRES_USER=user
            - POSTGRES_PASSWORD=password
            - POSTGRES_DB=deuna_challenge
        ```
    3. Press connect and you are done, you should be able to see the database tables, as well as the information that is added in the initSQL, you will also be able to see each change you make when you execute a post to create a payment or to make a refund.

## If you have any questions or issues, feel free to open an issue in this repository or contact me.
