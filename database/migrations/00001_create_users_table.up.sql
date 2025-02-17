CREATE TABLE users
(
    id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    first_name    VARCHAR(255)        NOT NULL,
    second_name   VARCHAR(255)        NOT NULL,
    email         VARCHAR(255) UNIQUE NOT NULL,
    nickname      VARCHAR(255) NULL,
    role          VARCHAR(255)        NOT NULL DEFAULT 'default',
    password_hash VARCHAR(255) NULL,
    created_at    TIMESTAMP                    DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP                    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP NULL,
    INDEX         idx_users_email (email),
    INDEX         idx_users_nickname (nickname)
);