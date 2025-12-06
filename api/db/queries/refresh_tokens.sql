-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens(eth_address, token_hash, expires_at, ip_address, user_agent, device_name)
VALUES($1, $2, $3, $4, $5, $6) RETURNING id;

-- name: RevokeRefreshToken :execrows
UPDATE refresh_tokens SET revoked_at = CURRENT_TIMESTAMP 
WHERE token_hash = $1;