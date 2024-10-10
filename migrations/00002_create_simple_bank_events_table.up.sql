CREATE TABLE events (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    event_type VARCHAR(255) NOT NULL,
    event_payload JSON NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_deleted_at (deleted_at)
);

INSERT INTO
    events (
        event_type,
        event_payload,
        status
    )
VALUES (
        'user_created',
        '{"user_id": "1", "username": "jdoe", "email": "jdoe@example.com"}',
        'processed'
    ),
    (
        'transaction_completed',
        '{"transaction_id": "101", "amount": 250.00, "currency": "USD", "account_id": "A12345"}',
        'processed'
    ),
    (
        'account_updated',
        '{"account_id": "A12345", "status": "active", "balance": 1250.00}',
        'pending'
    ),
    (
        'user_login',
        '{"user_id": "2", "login_time": "2024-10-10T08:00:00Z", "ip_address": "192.168.1.10"}',
        'failed'
    ),
    (
        'password_reset',
        '{"user_id": "3", "reset_time": "2024-10-10T09:00:00Z", "email": "mjones@example.com"}',
        'processed'
    );