CREATE TABLE users (
    id CHAR(36) NOT NULL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    status ENUM(
        'active',
        'inactive',
        'suspended'
    ) NOT NULL DEFAULT 'active',
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_deleted_at (deleted_at)
);

INSERT INTO
    users (
        id,
        username,
        password,
        email,
        first_name,
        last_name,
        status,
        is_active
    )
VALUES (
        '11111111-1111-1111-1111-111111111111',
        'jdoe',
        'password123',
        'jdoe@example.com',
        'John',
        'Doe',
        'active',
        TRUE
    ),
    (
        '22222222-2222-2222-2222-222222222222',
        'asmith',
        'securepass',
        'asmith@example.com',
        'Alice',
        'Smith',
        'active',
        TRUE
    ),
    (
        '33333333-3333-3333-3333-333333333333',
        'rjohnson',
        'passw0rd',
        'rjohnson@example.com',
        'Robert',
        'Johnson',
        'suspended',
        FALSE
    ),
    (
        '44444444-4444-4444-4444-444444444444',
        'mbrown',
        'mypassword',
        'mbrown@example.com',
        'Michael',
        'Brown',
        'inactive',
        FALSE
    ),
    (
        '55555555-5555-5555-5555-555555555555',
        'cwhite',
        'secure123',
        'cwhite@example.com',
        'Carol',
        'White',
        'active',
        TRUE
    );