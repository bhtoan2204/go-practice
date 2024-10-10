CREATE TABLE notifications (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

INSERT INTO
    notifications (
        user_id,
        title,
        message,
        is_read
    )
VALUES (
        '11111111-1111-1111-1111-111111111111',
        'Welcome!',
        'Thanks for joining Simple Bank!',
        FALSE
    ),
    (
        '22222222-2222-2222-2222-222222222222',
        'Transaction Alert',
        'You withdrew $200 from your savings account.',
        FALSE
    ),
    (
        '33333333-3333-3333-3333-333333333333',
        'Account Suspended',
        'Your account has been suspended.',
        TRUE
    ),
    (
        '44444444-4444-4444-4444-444444444444',
        'New Login Detected',
        'Your account was accessed from a new device.',
        FALSE
    ),
    (
        '55555555-5555-5555-5555-555555555555',
        'Password Reset',
        'Your password was successfully reset.',
        TRUE
    );