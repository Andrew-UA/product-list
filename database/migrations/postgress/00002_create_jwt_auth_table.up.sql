CREATE TABLE jwt_auth
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT       NOT NULL,
    token_id   VARCHAR(512) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP    NOT NULL,

    CONSTRAINT fk_jwt_auth_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX idx_jwt_auth_user_id ON jwt_auth (user_id);
CREATE INDEX idx_jwt_auth_token_id ON jwt_auth (token_id);
CREATE INDEX idx_jwt_auth_expires_at ON jwt_auth (expires_at);
