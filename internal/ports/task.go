package ports

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
)

type Task interface {
	Sign(ctx context.Context, id uuid.UUID) error
	Reject(ctx context.Context, id uuid.UUID) error

	Get(ctx context.Context, id uuid.UUID) (*models.Task, error)
	List(ctx context.Context) ([]*models.Task, error)
	Create(ctx context.Context, task *models.Task) (uuid.UUID, error)
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
}
