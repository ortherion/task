package ports

import (
	"context"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
)

type MessageService interface {
	Write(ctx context.Context, task *models.Task) error
	WriteAll(ctx context.Context) error
}
