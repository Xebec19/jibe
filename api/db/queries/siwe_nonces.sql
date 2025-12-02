-- name: CreateNonce :one
INSERT INTO siwe_nonces(value, eth_address, expires_at)
values($1, $2, $3) RETURNING value;

-- name: ConsumeNonce :execrows
UPDATE siwe_nonces SET used = TRUE, eth_address=$1
WHERE value = $2 and expires_at > CURRENT_TIMESTAMP;
