//go:generate mockgen -source redis_repository.go -destination mock/redis_repository.go -package mock
package session

import (
	"context"

	"github.com/dinorain/checkoutaja/internal/models"
)

// Session repository
type SessRepository interface {
	CreateSession(ctx context.Context, session *models.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}
