CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY,
    session_order SERIAL,
    ip_address VARCHAR(100) NOT NULL,
    user_id UUID NOT NULL,
    fcm_token VARCHAR(250),
    platform_name VARCHAR(150) NOT NULL,
    platform_type VARCHAR(10) CHECK (platform_type IN ('mobile')) NOT NULL,
    login_at TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
    created_at TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);