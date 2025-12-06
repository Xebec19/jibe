-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens(eth_address, token_hash, expires_at, device_info)
VALUES($1, $2, $3, $4) RETURNING id;