CREATE TABLE jwt_auth
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id    INTEGER   NOT NULL,
    token_id   TEXT      NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_jwt_auth_user_id ON jwt_auth (user_id);
CREATE INDEX idx_jwt_auth_token_id ON jwt_auth (token_id);
CREATE INDEX idx_jwt_auth_expires_at ON jwt_auth (expires_at);