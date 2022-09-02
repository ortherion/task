package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"gitlab.com/g6834/team17/task-service/internal/constants"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
	"time"
)

func (p *Producer) SendEvent(ctx context.Context, task *models.Task, event models.EventType) error {
	user, ok := ctx.Value(constants.CTX_USER).(*models.User)
	if !ok {
		return models.ErrCastUser
	}

	eventMsg := models.Event{
		TaskID:    task.ID,
		EventType: event.String(),
	}

	switch event {
	case models.Created:
		eventMsg.UserUUID = task.CreatorID
		eventMsg.Timestamp = task.CreatedDate
	case models.ApprovedBy:
		eventMsg.UserUUID = user.ID
		eventMsg.Timestamp = task.UpdatedDate
	case models.RejectedBy:
		eventMsg.UserUUID = user.ID
		eventMsg.Timestamp = task.UpdatedDate
	case models.Signed:
		eventMsg.UserUUID = user.ID
		eventMsg.Timestamp = task.UpdatedDate
	case models.Sent:
		eventMsg.UserUUID = task.CreatorID
		eventMsg.Timestamp = time.Now()
	default:
		eventMsg.EventType = models.Unknown.String()
		eventMsg.UserUUID = task.CreatorID
		eventMsg.Timestamp = time.Now()
	}

	data, err := json.Marshal(&eventMsg)
	if err != nil {
		return err
	}

	if err := p.WriteMessages(ctx, kafka.Message{
		Key:   task.ID.Bytes(),
		Value: data,
		Time:  time.Now(),
	},
	); err != nil {
		return err
	}
	return nil
}
