CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- access_tokens table :- it keeps jit of generated jwt tokens
CREATE TABLE IF NOT EXISTS access_tokens(
    jti uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    eth_address varchar(42) not null,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP NOT NULL
);

-- refresh_tokens table :- it keeps refresh tokens which can be used to generate access token
CREATE TABLE IF NOT EXISTS refresh_tokens(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    eth_address varchar(42) not null,
    token_hash VARCHAR(255) NOT NULL,  -- store hashed, not plain
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    ip_address VARCHAR(45),
    user_agent TEXT,
    device_name VARCHAR(255),  -- e.g., "Chrome on Windows"
    family_id UUID                      -- optional: for rotation detection
);