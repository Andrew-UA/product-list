CREATE TABLE users (
                       id            BIGSERIAL PRIMARY KEY,
                       first_name    VARCHAR(255) NOT NULL,
                       second_name   VARCHAR(255) NOT NULL,
                       email         VARCHAR(255) UNIQUE NOT NULL,
                       nickname      VARCHAR(255),
                       role          VARCHAR(255) NOT NULL DEFAULT 'default',
                       password_hash VARCHAR(255),
                       created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at    TIMESTAMP,

                       CONSTRAINT uq_users_email UNIQUE (email)
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_nickname ON users(nickname);
