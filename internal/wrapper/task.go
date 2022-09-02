package wrapper

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
	"gitlab.com/g6834/team17/task-service/internal/ports"
)

type TaskWrapper struct {
	taskS ports.Task
	mp    ports.MessageProducer[models.Task]
}

func NewTaskWrapper[T any](taskS ports.Task, mp ports.MessageProducer[models.Task]) *TaskWrapper {
	return &TaskWrapper{
		taskS: taskS,
		mp:    mp,
	}
}

func (t *TaskWrapper) Sign(ctx context.Context, id uuid.UUID) error {
	err := t.taskS.Sign(ctx, id)
	if err != nil {
		return err
	}
	if err = t.sendMessage(ctx, id); err != nil {
		return err
	}

	return nil
}

func (t *TaskWrapper) Reject(ctx context.Context, id uuid.UUID) error {
	err := t.taskS.Reject(ctx, id)
	if err != nil {
		return err
	}

	if err = t.sendMessage(ctx, id); err != nil {
		return err
	}

	return nil
}

func (t *TaskWrapper) Get(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	return t.taskS.Get(ctx, id)
}

func (t *TaskWrapper) List(ctx context.Context) ([]*models.Task, error) {
	return t.taskS.List(ctx)
}

func (t *TaskWrapper) Create(ctx context.Context, task *models.Task) (uuid.UUID, error) {
	id, err := t.taskS.Create(ctx, task)
	if err = t.sendMessage(ctx, id); err != nil {
		return id, err
	}

	return id, nil
}

func (t *TaskWrapper) Update(ctx context.Context, task *models.Task) error {
	err := t.taskS.Update(ctx, task)
	if err != nil {
		return err
	}
	if err = t.sendMessage(ctx, task.ID); err != nil {
		return err
	}

	return nil
}

func (t *TaskWrapper) Delete(ctx context.Context, id uuid.UUID) error {
	err := t.taskS.Delete(ctx, id)
	if err != nil {
		return err
	}
	if err = t.sendMessage(ctx, id); err != nil {
		return err
	}

	return nil
}

func (t *TaskWrapper) sendMessage(ctx context.Context, id uuid.UUID) error {
	data, err := t.taskS.Get(ctx, id)
	if err != nil {
		return err
	}

	if err := t.mp.Send(id.String(), data); err != nil {
		return err
	}

	return nil
}
