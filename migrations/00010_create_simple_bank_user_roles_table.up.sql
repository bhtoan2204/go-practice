CREATE TABLE user_roles (
    user_id CHAR(36) NOT NULL,
    role_id INT NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE
);

INSERT INTO
    user_roles (user_id, role_id)
VALUES (
        '11111111-1111-1111-1111-111111111111',
        1
    ), -- John Doe as Admin
    (
        '22222222-2222-2222-2222-222222222222',
        2
    ), -- Alice Smith as Customer
    (
        '33333333-3333-3333-3333-333333333333',
        3
    ), -- Robert Johnson as Teller
    (
        '44444444-4444-4444-4444-444444444444',
        4
    ), -- Michael Brown as Manager
    (
        '55555555-5555-5555-5555-555555555555',
        5
    );
-- Carol White as Auditor