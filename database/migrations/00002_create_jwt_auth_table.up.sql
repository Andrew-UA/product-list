CREATE TABLE jwt_auth
(
    id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id    BIGINT UNSIGNED NOT NULL,
    token      VARCHAR(512) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP    NOT NULL,

    INDEX      idx_jwt_auth_user_id (user_id),
    INDEX      idx_jwt_auth_expires_at (expires_at),

    CONSTRAINT fk_jwt_auth_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);