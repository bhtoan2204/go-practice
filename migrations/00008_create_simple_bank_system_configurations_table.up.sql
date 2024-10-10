CREATE TABLE system_configurations (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    config_key VARCHAR(255) NOT NULL UNIQUE,
    config_value TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_deleted_at (deleted_at)
);

INSERT INTO
    system_configurations (
        config_key,
        config_value,
        description
    )
VALUES (
        'max_login_attempts',
        '5',
        'Max login attempts before lockout'
    ),
    (
        'session_timeout',
        '30',
        'Session timeout in minutes'
    ),
    (
        'currency',
        'USD',
        'Default transaction currency'
    ),
    (
        'maintenance_mode',
        'false',
        'Flag indicating if the application is in maintenance mode'
    ),
    (
        'support_email',
        'support@simplebank.com',
        'Email address for customer support'
    ),
    (
        'min_password_length',
        '8',
        'Minimum password length for user accounts'
    );