package ports

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
)

type TaskStorage interface {
	Get(ctx context.Context, id uuid.UUID) (*models.Task, error)
	List(ctx context.Context) ([]*models.Task, error)
	Create(ctx context.Context, task *models.Task) (uuid.UUID, error)
	Update(ctx context.Context, task *models.Task) error
	UpdateSign(ctx context.Context, id uuid.UUID, status models.Stage) error
	Delete(ctx context.Context, id uuid.UUID) error
}
