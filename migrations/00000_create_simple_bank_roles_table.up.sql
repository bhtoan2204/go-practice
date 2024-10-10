CREATE TABLE roles (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_deleted_at (deleted_at)
);

INSERT INTO
    roles (name, description)
VALUES (
        'Admin',
        'Administrator with full access to manage the application'
    ),
    (
        'Customer',
        'Standard user with access to their own bank accounts and transactions'
    ),
    (
        'Teller',
        'Bank employee who can assist customers with transactions'
    ),
    (
        'Manager',
        'Bank manager with oversight of teller activities'
    ),
    (
        'Auditor',
        'User with read-only access for auditing purposes'
    );