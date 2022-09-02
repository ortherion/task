package task

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/g6834/team17/task-service/internal/constants"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
	"gitlab.com/g6834/team17/task-service/internal/ports"
	"gitlab.com/g6834/team17/task-service/internal/utils"
	"time"
)

type Service struct {
	db     ports.TaskStorage
	auth   ports.Auth
	sender ports.AnalyticsService
}

func New(db ports.TaskStorage, auth ports.Auth, sender ports.AnalyticsService) *Service {
	return &Service{
		db:     db,
		auth:   auth,
		sender: sender,
	}
}

func (s *Service) Sign(ctx context.Context, id uuid.UUID) error {
	ctx, span := utils.StartSpan(ctx)
	defer span.End()

	return s.updateTaskStatus(ctx, id, models.Accept)
}

func (s *Service) Reject(ctx context.Context, id uuid.UUID) error {
	ctx, span := utils.StartSpan(ctx)
	defer span.End()

	return s.updateTaskStatus(ctx, id, models.Reject)
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	ctx, span := utils.StartSpan(ctx)
	defer span.End()

	return s.db.Get(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]*models.Task, error) {
	ctx, span := utils.StartSpan(ctx)
	defer span.End()

	return s.db.List(ctx)
}

func (s *Service) Create(ctx context.Context, task *models.Task) (uuid.UUID, error) {
	ctx, span := utils.StartSpan(ctx)
	defer span.End()

	task.CreatedDate = time.Now()

	taskID, err := s.db.Create(ctx, task)
	if err != nil {
		return taskID, err
	}

	if err := s.sender.SendEvent(ctx, task, models.Created); err != nil {
		return uuid.UUID{}, err
	}

	return taskID, nil
}

func (s *Service) Update(ctx context.Context, task *models.Task) error {
	ctx, span := utils.StartSpan(ctx)
	defer span.End()

	user, ok := ctx.Value(constants.CTX_USER).(*models.User)
	if !ok {
		return models.ErrCastUser
	}

	oldTask, err := s.db.Get(ctx, task.ID)
	if err != nil {
		return err
	}

	if user.ID != oldTask.ID {
		return models.ErrUserNotHavePermissions
	}

	task.UpdatedDate = time.Now()

	if err := s.sender.SendEvent(ctx, task, models.Updated); err != nil {
		return err
	}

	return s.db.Update(ctx, task)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span := utils.StartSpan(ctx)
	defer span.End()

	user, ok := ctx.Value(constants.CTX_USER).(*models.User)
	if !ok {
		return models.ErrCastUser
	}

	task, err := s.db.Get(ctx, id)
	if err != nil {
		return err
	}

	if task.ID != user.ID {
		return models.ErrUserNotHavePermissions
	}

	task.DeletedDate = time.Now()
	task.IsDeleted = true

	if err := s.sender.SendEvent(ctx, task, models.Deleted); err != nil {
		return err
	}

	return s.db.Delete(ctx, id)
}

func (s *Service) updateTaskStatus(ctx context.Context, id uuid.UUID, status models.Stage) error {
	user, ok := ctx.Value(constants.CTX_USER).(*models.User)
	if !ok {
		return models.ErrCastUser
	}

	task, err := s.db.Get(ctx, id)
	if err != nil {
		return err
	}

	for _, signatory := range task.Signatories {
		if signatory.TaskID != user.ID {
			return models.ErrUserNotSignatory
		}
	}

	switch status {
	case models.Reject:
		task.Stage = models.Reject
		task.UpdatedDate = time.Now()

		if err := s.db.Update(ctx, task); err != nil {
			return err
		}

		if err := s.sender.SendEvent(ctx, task, models.RejectedBy); err != nil {
			return err
		}

		for _, signatory := range task.Signatories {

			if signatory.ID == user.ID {
				signatory.Status = models.Reject

				if err := s.db.UpdateSign(ctx, signatory.ID, status); err != nil {
					return err
				}
			}
		}

	case models.Accept:
		for _, v := range task.Signatories {
			if v.ID == user.ID {
				if v.Status == models.Accept {
					return models.ErrUserHasAlreadySigned
				}

				v.Status = models.Accept
				task.UpdatedDate = time.Now()

				err := s.db.UpdateSign(ctx, v.ID, v.Status)
				if err != nil {
					return err
				}

				err = s.db.Update(ctx, task)
				if err != nil {
					return err
				}

				if err := s.sender.SendEvent(ctx, task, models.ApprovedBy); err != nil {
					return err
				}
			}

			if v.Status != models.Accept {
				task.Stage = models.InProcess
				return nil
			}

			task.Stage = models.Accept

			if err := s.sender.SendEvent(ctx, task, models.Signed); err != nil {
				return err
			}

			if err := s.db.Update(ctx, task); err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("only accept or reject status can be use")
	}

	return nil
}
