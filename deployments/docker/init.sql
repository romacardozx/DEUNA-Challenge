-- Crear tabla customers primero
CREATE TABLE IF NOT EXISTS customers (
    id VARCHAR(36) PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20),
    address TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_customers_email ON customers(email);

-- Luego crear tabla merchants
CREATE TABLE IF NOT EXISTS merchants (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20),
    address TEXT,
    business_type VARCHAR(100),
    tax_id VARCHAR(50),
    bank_account_info TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_merchants_email ON merchants(email);

-- Ahora podemos crear la tabla payments
CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(36) PRIMARY KEY,
    customer_id VARCHAR(36) NOT NULL,
    merchant_id VARCHAR(36) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL,
    transaction_id VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (merchant_id) REFERENCES merchants(id)
);

CREATE INDEX idx_payments_customer ON payments(customer_id);
CREATE INDEX idx_payments_merchant ON payments(merchant_id);
CREATE INDEX idx_payments_status ON payments(status);

-- Finalmente, crear la tabla refunds
CREATE TABLE IF NOT EXISTS refunds (
    id VARCHAR(36) PRIMARY KEY,
    payment_id VARCHAR(36) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    reason TEXT,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (payment_id) REFERENCES payments(id)
);

CREATE INDEX idx_refunds_payment ON refunds(payment_id);
CREATE INDEX idx_refunds_status ON refunds(status);


-- Insert initial merchants
INSERT INTO merchants (id, name, email, phone, business_type, tax_id) VALUES
('m1', 'Merchant 1', 'merchant1@example.com', '1234567890', 'Retail', 'TAX123'),
('m2', 'Merchant 2', 'merchant2@example.com', '0987654321', 'Online', 'TAX456');

-- Insert initial customers
INSERT INTO customers (id, first_name, last_name, email, phone) VALUES
('c1', 'John', 'Doe', 'john@example.com', '1112223333'),
('c2', 'Jane', 'Smith', 'jane@example.com', '4445556666');

-- Insert merchant tokens (for basic auth)
CREATE TABLE IF NOT EXISTS merchant_tokens (
    merchant_id VARCHAR(36) PRIMARY KEY,
    token VARCHAR(100) NOT NULL UNIQUE,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id)
);

INSERT INTO merchant_tokens (merchant_id, token) VALUES
('m1', 'token1'),
('m2', 'token2');

-- Crear tabla de auditoría
CREATE TABLE IF NOT EXISTS audit_logs (
    id SERIAL PRIMARY KEY,
    payment_id VARCHAR(36),
    amount DECIMAL(10, 2),
    merchant_id VARCHAR(36),
    customer_id VARCHAR(36),
    currency VARCHAR(3),
    transaction_id VARCHAR(100),
    status VARCHAR(20),
    message TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Crear índices para la tabla de auditoría
CREATE INDEX idx_audit_logs_payment_id ON audit_logs(payment_id);
CREATE INDEX idx_audit_logs_merchant_id ON audit_logs(merchant_id);
CREATE INDEX idx_audit_logs_customer_id ON audit_logs(customer_id);
CREATE INDEX idx_audit_logs_status ON audit_logs(status);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);

-- Crear función para insertar logs de auditoría
CREATE OR REPLACE FUNCTION insert_audit_log(
    p_payment_id VARCHAR(36),
    p_amount DECIMAL(10, 2),
    p_merchant_id VARCHAR(36),
    p_customer_id VARCHAR(36),
    p_currency VARCHAR(3),
    p_transaction_id VARCHAR(100),
    p_status VARCHAR(20),
    p_message TEXT
) RETURNS VOID AS $$
BEGIN
    INSERT INTO audit_logs (
        payment_id, amount, merchant_id, customer_id, currency, 
        transaction_id, status, message
    ) VALUES (
        p_payment_id, p_amount, p_merchant_id, p_customer_id, p_currency, 
        p_transaction_id, p_status, p_message
    );
END;
$$ LANGUAGE plpgsql;