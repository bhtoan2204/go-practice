CREATE TABLE transactions (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    account_id BIGINT UNSIGNED NOT NULL,
    transaction_type ENUM(
        'deposit',
        'withdrawal',
        'transfer'
    ) NOT NULL,
    amount DECIMAL(18, 2) NOT NULL,
    currency VARCHAR(10) NOT NULL,
    reference_id VARCHAR(255),
    description TEXT,
    status ENUM(
        'pending',
        'success',
        'failed'
    ) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (account_id) REFERENCES accounts (id) ON DELETE CASCADE
);

INSERT INTO
    transactions (
        account_id,
        transaction_type,
        amount,
        currency,
        reference_id,
        description,
        status
    )
VALUES (
        1,
        'deposit',
        500.00,
        'USD',
        'TX123',
        'Initial deposit',
        'success'
    ),
    (
        1,
        'withdrawal',
        100.00,
        'USD',
        'TX124',
        'ATM withdrawal',
        'success'
    ),
    (
        2,
        'deposit',
        200.00,
        'USD',
        'TX125',
        'Savings deposit',
        'success'
    ),
    (
        3,
        'transfer',
        300.00,
        'USD',
        'TX126',
        'Transfer to external account',
        'pending'
    ),
    (
        4,
        'deposit',
        700.00,
        'EUR',
        'TX127',
        'Account opening deposit',
        'success'
    ),
    (
        5,
        'withdrawal',
        400.00,
        'USD',
        'TX128',
        'Online purchase',
        'failed'
    ),
    (
        1,
        'withdrawal',
        50.00,
        'USD',
        'TX129',
        'Gas station purchase',
        'success'
    );