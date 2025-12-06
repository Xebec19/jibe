-- name: CreateAccessToken :one
INSERT INTO access_tokens(eth_address, expires_at) 
VALUES($1, $2) RETURNING jti;
