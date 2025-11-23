-- name: CreateNonce :one
INSERT INTO siwe_nonces(value, eth_address, expires_at) 
values($1, $2, $3) RETURNING value;