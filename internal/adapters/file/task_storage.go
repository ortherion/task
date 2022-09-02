package file

import (
	"context"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
	"gitlab.com/g6834/team17/task-service/internal/ports"
)

var _ ports.TaskStorage = (*Database)(nil)

func (db *Database) Get(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	tasks, err := db.getTasks()
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}

	return nil, models.ErrNotFound
}

func (db *Database) List(ctx context.Context) ([]*models.Task, error) {
	return db.getTasks()
}

// В Postgre нужно использовать returning
// Exampe: insert into tasks (/.../) values (/.../) returning id;
func (db *Database) Create(ctx context.Context, task *models.Task) (uuid.UUID, error) {
	data, err := json.Marshal(task)
	if err != nil {
		return uuid.UUID{}, err
	}

	if err := db.Append(data); err != nil {
		return uuid.UUID{}, err
	}

	return uuid.NewV4(), nil
}

func (db *Database) Update(ctx context.Context, task *models.Task) error {
	tasks, err := db.getTasks()
	if err != nil {
		return err
	}

	for i, v := range tasks {
		if v.ID == task.ID {
			tasks[i] = task
			break
		}
	}

	return db.rewriteTasks(tasks)
}

func (db *Database) Delete(ctx context.Context, id uuid.UUID) error {
	tasks, err := db.getTasks()
	if err != nil {
		return err
	}

	for i, v := range tasks {
		if v.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	return db.rewriteTasks(tasks)
}

func (db *Database) getTasks() ([]*models.Task, error) {
	data := db.Read()

	var tasks []*models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db *Database) rewriteTasks(tasks []*models.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	if err = db.Truncate(); err != nil {
		return err
	}

	if err = db.Append(data); err != nil {
		return err
	}

	return nil
}

func (db Database) UpdateSign(ctx context.Context, id uuid.UUID, status models.Stage) error {
	//TODO: implement
	return nil
}
