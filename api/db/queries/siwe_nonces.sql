-- name: CreateNonce :exec
INSERT INTO siwe_nonces(value, eth_address, expires_at) 
values($1, $2, $3);