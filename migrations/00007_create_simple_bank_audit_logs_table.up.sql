CREATE TABLE audit_logs (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255),
    action VARCHAR(255) NOT NULL,
    details TEXT,
    ip_address VARCHAR(45),
    device_token VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL
);

INSERT INTO
    audit_logs (
        user_id,
        action,
        details,
        ip_address,
        device_token
    )
VALUES (
        '11111111-1111-1111-1111-111111111111',
        'login',
        'User logged in',
        '192.168.1.1',
        'token1'
    ),
    (
        '22222222-2222-2222-2222-222222222222',
        'withdrawal',
        'User completed a withdrawal',
        '192.168.1.2',
        'token2'
    ),
    (
        '33333333-3333-3333-3333-333333333333',
        'account_suspended',
        'Account suspended due to inactivity',
        '192.168.1.3',
        'token3'
    ),
    (
        '44444444-4444-4444-4444-444444444444',
        'login',
        'User logged in from new device',
        '192.168.1.4',
        'token4'
    ),
    (
        '55555555-5555-5555-5555-555555555555',
        'password_reset',
        'User reset password',
        '192.168.1.5',
        'token5'
    );