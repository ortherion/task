package http

import (
	"errors"
	"github.com/go-chi/chi/v5"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/g6834/team17/task-service/internal/adapters/presenter/requests"
	"gitlab.com/g6834/team17/task-service/internal/constants"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
	"gitlab.com/g6834/team17/task-service/internal/utils"
	"net/http"
)

func (s *Server) taskHandlers() http.Handler {
	h := chi.NewMux()

	h.Use(s.ValidateAuth())

	h.Group(func(r chi.Router) {
		h.Get("/", s.ListTask)
		h.Get("/{id}", s.GetTask)
		h.Post("/", s.CreateTask)
		h.Patch("/{id}", s.UpdateTask)
		h.Delete("/{id}", s.DeleteTask)
		h.Put("/approve/{id}", s.Approve)
		h.Put("/reject/{id}", s.Reject)
	})

	return h
}

// GetTask
// @ID GetTask
// @tags task
// @Summary Get Task by ID
// @Description Get Task by ID.
// @Produce json
// @Param access_token header string true "access token"
// @Param refresh_token header string true "refresh token"
// @Param id query string true "id"
// @Success 200 {object} models.Task "ok"
// @Failure 400 {object} string "400 bad request"
// @Failure 404 {string} string "404 page not found"
// @Failure 403 {object} string "403 forbidden"
// @Failure 500 {object} string "500 internal error"
// @Router / [get]
func (s *Server) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	ID := chi.URLParam(r, "id")

	taskID, err := uuid.FromString(ID)
	if err != nil {
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	task, err := s.task.Get(ctx, taskID)
	if err != nil {
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	s.presenter.JSON(w, r, task)
}

// ListTask
// @ID ListTask
// @tags task
// @Summary Get All Tasks
// @Description Get List of all tasks
// @Produce json
// @Param access_token header string true "access token"
// @Param refresh_token header string true "refresh token"
// @Success 200 {object} []models.Task "ok"
// @Failure 400 {object} string "400 bad request"
// @Failure 404 {string} string "404 page not found"
// @Failure 403 {object} string "403 forbidden"
// @Failure 500 {object} string "500 internal error"
// @Router / [get]
func (s *Server) ListTask(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	tasks, err := s.task.List(ctx)
	if err != nil {
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	s.presenter.JSON(w, r, tasks)
}

// CreateTask
// @ID CreateTask
// @tags task
// @Summary Create a Task
// @Description Accept a Task Description, returning task's id
// @Accept json
// @Produce json
// @Param access_token header string true "access token"
// @Param refresh_token header string true "refresh token"
// @Success 200 {object} models.Task "ok"
// @Failure 400 {string} string "400 bad request"
// @Failure 404 {string} string "404 page not found"
// @Failure 403 {string} string "403 forbidden"
// @Failure 500 {string} string "500 internal error"
// @Router / [post]
func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	var req *requests.Task

	if err := utils.ReadJson(r, &req); err != nil {
		s.presenter.Error(w, r, models.ErrorBadRequest(err))
		return
	}

	user, ok := ctx.Value(constants.CTX_USER).(*models.User)
	if !ok {
		s.presenter.Error(w, r, models.ErrorInternal(models.ErrCastUser))
	}

	task := &models.Task{
		Title:       req.Title,
		Body:        req.Body,
		CreatorID:   user.ID,
		Stage:       models.Undefined,
		Signatories: nil,
	}

	for _, mail := range req.SignsMails {
		task.Signatories = append(task.Signatories, models.Signatory{
			Email:  mail,
			Status: models.Undefined,
		})
	}

	taskID, err := s.task.Create(ctx, task)
	if err != nil {
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	s.presenter.JSON(w, r, taskID)
}

// UpdateTask
// @ID UpdateTask
// @tags task
// @Summary Update a Task
// @Description Update a Task Description, returning task's id
// @Accept json
// @Produce json
// @Param access_token header string true "access token"
// @Param refresh_token header string true "refresh token"
// @Param id query string true "id"
// @Success 200 {object} models.Task "ok"
// @Failure 400 {string} string "400 bad request"
// @Failure 404 {string} string "404 page not found"
// @Failure 403 {string} string "403 forbidden"
// @Failure 500 {string} string "500 internal error"
// @Router / [patch]
func (s *Server) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	var req *requests.Task

	ID := chi.URLParam(r, "id")

	taskID, err := uuid.FromString(ID)
	if err != nil {
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	if err := utils.ReadJson(r, &req); err != nil {
		s.presenter.Error(w, r, models.ErrorBadRequest(err))
		return
	}

	task := &models.Task{
		ID:    taskID,
		Title: req.Title,
		Body:  req.Body,
	}

	if err := s.task.Update(ctx, task); err != nil {
		if errors.Is(err, models.ErrUserNotHavePermissions) {
			s.presenter.Error(w, r, models.ErrorForbidden(err))
			return
		}
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	s.presenter.JSON(w, r, task.ID)
}

// DeleteTask
// @ID DeleteTask
// @tags task
// @Summary Delete a Task
// @Description Delete a Task
// @Param access_token header string true "access token"
// @Param refresh_token header string true "refresh token"
// @Param id query string true "id"
// @Success 200 {string} string true "ok"
// @Failure 400 {string} string "400 bad request"
// @Failure 404 {string} string "404 page not found"
// @Failure 403 {string} string "403 forbidden"
// @Failure 500 {string} string "500 internal error"
// @Router / [delete]
func (s *Server) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	ID := chi.URLParam(r, "id")

	taskID, err := uuid.FromString(ID)
	if err != nil {
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	err = s.task.Delete(ctx, taskID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotHavePermissions) {
			s.presenter.Error(w, r, models.ErrorForbidden(err))
			return
		}
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	s.presenter.JSON(w, r, nil)
}

// Approve
// @ID Approve
// @tags task
// @Summary Approve the Task
// @Description Approve the Task
// @Param access_token header string true "access token"
// @Param refresh_token header string true "refresh token"
// @Param id query string true "id"
// @Success 200 {string} string true "ok"
// @Failure 400 {string} string "400 bad request"
// @Failure 404 {string} string "404 page not found"
// @Failure 403 {string} string "403 forbidden"
// @Failure 500 {string} string "500 internal error"
// @Router /approve/ [put]
func (s *Server) Approve(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	ID := chi.URLParam(r, "id")

	taskID, err := uuid.FromString(ID)
	if err != nil {
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	err = s.task.Sign(ctx, taskID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotHavePermissions) {
			s.presenter.Error(w, r, models.ErrorForbidden(err))
			return
		}
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	s.presenter.JSON(w, r, nil)
}

// Reject
// @ID Reject
// @tags task
// @Summary Reject the Task
// @Description Reject the Task
// @Param access_token header string true "access token"
// @Param refresh_token header string true "refresh token"
// @Param id query string true "id"
// @Success 200 {string} string true "ok"
// @Failure 400 {string} string "400 bad request"
// @Failure 404 {string} string "404 page not found"
// @Failure 403 {string} string "403 forbidden"
// @Failure 500 {string} string "500 internal error"
// @Router /Reject/ [put]
func (s *Server) Reject(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	ID := chi.URLParam(r, "id")

	taskID, err := uuid.FromString(ID)
	if err != nil {
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	err = s.task.Reject(ctx, taskID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotHavePermissions) {
			s.presenter.Error(w, r, models.ErrorForbidden(err))
			return
		}
		s.presenter.Error(w, r, models.ErrorInternal(err))
		return
	}

	s.presenter.JSON(w, r, nil)
}
