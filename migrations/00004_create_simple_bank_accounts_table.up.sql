CREATE TABLE accounts (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    account_number VARCHAR(255) NOT NULL,
    account_type VARCHAR(50) NOT NULL,
    balance BIGINT NOT NULL,
    currency VARCHAR(10) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

INSERT INTO
    accounts (
        user_id,
        account_number,
        account_type,
        balance,
        currency,
        status
    )
VALUES (
        '11111111-1111-1111-1111-111111111111',
        '1234567890',
        'checking',
        500000,
        'USD',
        'active'
    ),
    (
        '11111111-1111-1111-1111-111111111111',
        '1234567891',
        'savings',
        1000000,
        'USD',
        'active'
    ),
    (
        '22222222-2222-2222-2222-222222222222',
        '9876543210',
        'savings',
        1500000,
        'USD',
        'active'
    ),
    (
        '33333333-3333-3333-3333-333333333333',
        '4567891230',
        'checking',
        250000,
        'USD',
        'suspended'
    ),
    (
        '44444444-4444-4444-4444-444444444444',
        '5678901234',
        'checking',
        800000,
        'EUR',
        'inactive'
    ),
    (
        '55555555-5555-5555-5555-555555555555',
        '6789012345',
        'business',
        2000000,
        'USD',
        'active'
    );