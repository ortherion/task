package ports

import (
	"context"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
)

type AnalyticsService interface {
	SendEvent(ctx context.Context, task *models.Task, event models.EventType) error
}
