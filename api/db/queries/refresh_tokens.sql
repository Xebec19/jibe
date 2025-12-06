-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens(eth_address, token_hash, expires_at, ip_address, user_agent, device_name)
VALUES($1, $2, $3, $4, $5, $6) RETURNING id;