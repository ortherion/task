package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
	"time"
)

func (p *Producer) WriteAll(ctx context.Context) error {
	tasks, err := p.taskService.List(ctx)
	if err != nil {
		return err
	}

	messages := make([]kafka.Message, 0, len(tasks))

	for _, task := range tasks {
		data, err := json.Marshal(task)
		if err != nil {
			return err
		}

		message := kafka.Message{
			Key:     task.ID.Bytes(),
			Value:   data,
			Headers: nil,
			Time:    time.Now(),
		}

		messages = append(messages, message)

		if err := p.SendEvent(ctx, task, models.Sent); err != nil {
			return err
		}
	}

	if err := p.WriteMessages(ctx, messages...); err != nil {
		return err
	}

	return nil
}

func (p *Producer) Write(ctx context.Context, task *models.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:     task.ID.Bytes(),
		Value:   data,
		Headers: nil,
		Time:    time.Now(),
	}

	if err := p.WriteMessages(ctx, message); err != nil {
		return err
	}

	return nil
}
