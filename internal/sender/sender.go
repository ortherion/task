package sender

import (
	"context"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
	"gitlab.com/g6834/team17/task-service/internal/ports"
)

//TODO: Add Wrapper

type Sender struct {
	aS ports.AnalyticsService
	mS ports.MessageService
}

func NewSender(service ports.AnalyticsService, messageService ports.MessageService) *Sender {
	return &Sender{
		aS: service,
		mS: messageService,
	}
}

func (s Sender) Send(ctx context.Context, task *models.Task, event models.EventType) error {
	if err := s.aS.SendEvent(ctx, task, event); err != nil {
		return err
	}
	if err := s.mS.Write(ctx, task); err != nil {
		return err
	}
	return nil
}

type AnalyticsSender struct {
	aS ports.AnalyticsService
}

func NewAnalyticsSender(service ports.AnalyticsService) *AnalyticsSender {
	return &AnalyticsSender{
		aS: service,
	}
}

func (s AnalyticsSender) SendEvent(ctx context.Context, task *models.Task, event models.EventType) error {
	return s.aS.SendEvent(ctx, task, event)
}

// MessageSender /* Better for goroutine in main */
type MessageSender struct {
	mS ports.MessageService
}

func NewMessageSender(service ports.MessageService) *MessageSender {
	return &MessageSender{
		mS: service,
	}
}

func (s MessageSender) WriteAll(ctx context.Context) error {
	return s.mS.WriteAll(ctx)
}

func (s MessageSender) Write(ctx context.Context, task *models.Task) error {
	return s.mS.Write(ctx, task)
}
