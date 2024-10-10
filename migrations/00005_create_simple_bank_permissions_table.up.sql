CREATE TABLE permissions (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    permission_name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_deleted_at (deleted_at)
);

INSERT INTO
    permissions (permission_name, description)
VALUES (
        'view_account',
        'Allows viewing of account details'
    ),
    (
        'edit_account',
        'Allows editing of account details'
    ),
    (
        'delete_account',
        'Allows deletion of accounts'
    ),
    (
        'create_transaction',
        'Allows creating new transactions'
    ),
    (
        'view_transaction',
        'Allows viewing of transaction history'
    ),
    (
        'manage_users',
        'Allows management of user accounts'
    ),
    (
        'view_audit_logs',
        'Allows viewing of audit logs for compliance'
    );