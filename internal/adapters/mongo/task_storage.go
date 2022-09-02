package mongo

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
	"gitlab.com/g6834/team17/task-service/internal/ports"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

var _ ports.TaskStorage = (*Database)(nil)

const (
	DB_COLLECTION = "tasks"
)

func (db *Database) Get(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	query := db.DB.Collection(DB_COLLECTION).FindOne(ctx, bson.M{"id": id})

	var task models.Task
	err := query.Decode(task)

	return &task, err
}

func (db *Database) List(ctx context.Context) ([]*models.Task, error) {
	query, err := db.DB.Collection(DB_COLLECTION).Find(ctx, bson.D{})
	defer query.Close(ctx)

	if err != nil {
		return nil, err
	}

	tasks := make([]*models.Task, 0)
	for query.Next(ctx) {
		var task models.Task
		err := query.Decode(task)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (db *Database) Create(ctx context.Context, task *models.Task) (uuid.UUID, error) {
	task.ID = uuid.NewV4()
	_, err := db.DB.Collection(DB_COLLECTION).InsertOne(ctx, task)
	if err != nil {
		return uuid.UUID{}, err
	}

	return task.ID, nil
}

func (db *Database) Update(ctx context.Context, task *models.Task) error {
	task.UpdatedDate = time.Now()

	dataReq := bson.M{
		"$set": bson.M{
			"signatories": task.Signatories,
			"update_date": task.UpdatedDate,
			"stage":       task.Stage,
		},
	}

	_, err := db.DB.Collection(DB_COLLECTION).UpdateOne(ctx, bson.M{"id": task.ID}, dataReq)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Delete(ctx context.Context, id uuid.UUID) error {
	dataReq := bson.M{
		"$set": bson.M{
			"is_deleted": true,
		},
	}

	_, err := db.DB.Collection(DB_COLLECTION).UpdateOne(ctx, bson.M{"id": id}, dataReq)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) UpdateSign(ctx context.Context, id uuid.UUID, status models.Stage) error {
	//TODO: implement
	return nil
}
