package repositories

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Xebec19/jibe/api/internal/db"
	"github.com/Xebec19/jibe/api/pkg/logger"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthRepository interface {
	CreateNonce(addr string) (string, error)

	// CheckNonce check if the given nonce exists in db and is valid. After validating,
	// it saves the address of the user who used it
	CheckNonce(nonce, addr string) (bool, error)
}

func NewAuthRepository(ctx context.Context, logger *logger.Logger, q *db.Queries) AuthRepository {

	return &authRepository{
		ctx:    ctx,
		logger: *logger,
		q:      q,
	}
}

type authRepository struct {
	ctx    context.Context
	logger logger.Logger
	q      *db.Queries
}

func (repo *authRepository) CreateNonce(addr string) (string, error) {

	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		repo.logger.Error("nonce generation failed", "error", err)
		return "", err
	}

	nonce := hex.EncodeToString(bytes)

	expireAt := time.Now().Add(10 * time.Minute)

	params := db.CreateNonceParams{
		Value:      nonce,
		EthAddress: pgtype.Text{String: addr, Valid: true},
		ExpiresAt: pgtype.Timestamp{
			Time:  expireAt,
			Valid: true,
		},
	}

	return repo.q.CreateNonce(repo.ctx, params)
}

func (repo *authRepository) CheckNonce(nonce, addr string) (bool, error) {

	arg := db.ConsumeNonceParams{
		EthAddress: pgtype.Text{Valid: true, String: addr},
		Value:      nonce,
	}

	rows, err := repo.q.ConsumeNonce(repo.ctx, arg)

	if err != nil {
		return false, fmt.Errorf("nonce consumption failed %w", err)
	}

	if rows == 0 {
		return false, fmt.Errorf("nonce not found")
	}

	return true, nil

}
