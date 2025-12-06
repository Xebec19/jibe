-- name: CreateAccessToken :one
INSERT INTO access_tokens(eth_address, expires_at) 
VALUES($1, $2) RETURNING jti;

-- name: RevokeAccessToken :execrows
UPDATE access_tokens SET revoked_at = CURRENT_TIMESTAMP 
WHERE jti = $1;