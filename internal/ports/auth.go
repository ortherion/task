package ports

import (
	"context"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
)

type Auth interface {
	Info(ctx context.Context, token *models.TokenPair) (*models.User, error)
	Validate(ctx context.Context, token *models.TokenPair) (*models.TokenPair, error)
	ParseToken(ctx context.Context, tokenString string) (*models.User, bool, error)
}
