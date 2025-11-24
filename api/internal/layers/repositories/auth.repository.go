package repositories

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/Xebec19/jibe/api/internal/db"
	"github.com/Xebec19/jibe/api/pkg/logger"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthRepository interface {
	CreateNonce(addr string) (string, error)
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
