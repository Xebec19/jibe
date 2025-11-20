-- SIWE_Nonce Table
-- It keeps track of nonce generated for Wallet signup
CREATE TABLE siwe_nonces (
    value VARCHAR(64) PRIMARY KEY,
    address VARCHAR(42),
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_expires_at (expires_at)
)