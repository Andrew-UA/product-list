CREATE TABLE users
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name    TEXT NOT NULL,
    second_name   TEXT NOT NULL,
    email         TEXT NOT NULL UNIQUE,
    nickname      TEXT,
    role          TEXT NOT NULL DEFAULT 'default',
    password_hash TEXT,
    created_at    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP
);

CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_nickname ON users (nickname);