package test

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
	"time"
)

var (
	taskID = uuid.NewV4()
	task   = models.Task{
		ID:        taskID,
		Title:     "test",
		Body:      "test",
		CreatorID: uuid.NewV4(),
		Stage:     models.Undefined,
		Signatories: []models.Signatory{{
			ID:     uuid.NewV4(),
			TaskID: taskID,
			Email:  "",
			Status: models.Undefined,
		}, {
			ID:     taskID,
			TaskID: uuid.NewV4(),
			Email:  "",
			Status: models.Undefined,
		}},
	}
)

func (suite *TestContainersSuite) TestTaskRepoCreateSuccess() {
	newTask := &models.Task{
		Title:       "test123",
		Body:        "test",
		CreatorID:   uuid.NewV4(),
		Stage:       models.Undefined,
		Signatories: task.Signatories,
	}
	newTask.CreatedDate = time.Now()

	id, err := suite.taskStorage.Create(context.Background(), newTask)
	suite.NoError(err)
	suite.NotNil(id)

	tasks, err := suite.taskStorage.List(context.Background())

	suite.NoError(err)
	suite.NotNil(tasks, "tasks must be not nil")
	suite.Condition(func() bool {
		return len(tasks) > 0
	})
}

func (suite *TestContainersSuite) TestTaskRepoCreateDuplicateName() {
	id, err := suite.taskStorage.Create(context.Background(), &task)

	suite.NotNil(err)
	suite.NotNil(id)
}

func (suite *TestContainersSuite) TestTaskRepoGetAllSuccess() {
	tasks, err := suite.taskStorage.List(context.Background())

	suite.NoError(err)
	suite.NotNil(tasks, "tasks must be not nil")
	suite.Condition(func() bool {
		return len(tasks) > 0
	})
}

func (suite *TestContainersSuite) TestTaskrRepoGetByIDSuccess() {
	taskDB, err := suite.taskStorage.Get(context.Background(), taskID)

	suite.NoError(err)
	suite.NotNil(taskDB, "task must be not nil")
	suite.Equal(taskDB, task)
}

func (suite *TestContainersSuite) TestTaskRepoGetSuccess() {
	suite.Suite.T().SkipNow()
	tasks, err := suite.taskStorage.List(context.Background())

	suite.NoError(err)
	suite.NotNil(tasks, "tasks must be not nil")
	suite.Condition(func() bool {
		return len(tasks) > 0
	})

	suite.Require().Contains(tasks, taskID)

	taskDB, err := suite.taskStorage.Get(context.Background(), taskID)

	suite.NoError(err)
	suite.NotNil(task, "task must be not nil")
	suite.Equal(task, taskDB)
}

func (suite *TestContainersSuite) TestTaskRepoGetInvalidId() {
	u, err := suite.taskStorage.Get(context.Background(), uuid.FromStringOrNil("lkfjsdlajfds"))

	suite.Nil(u)
	suite.NotNil(err)
}

func (suite *TestContainersSuite) TestTaskRepoUpdateSuccess() {
	newTask := &models.Task{
		ID:    task.ID,
		Title: "new",
		Body:  "new",
	}
	newTask.UpdatedDate = time.Now()

	err := suite.taskStorage.Update(context.Background(), newTask)
	suite.Nil(err)

	taskDB, err := suite.taskStorage.Get(context.Background(), taskID)

	suite.Equal(task.CreatedDate, taskDB.CreatedDate)
	suite.Equal(task.CreatorID, taskDB.CreatorID)
	suite.Equal(task.Signatories, taskDB.Signatories)
	suite.Equal(task.Stage, taskDB.Stage)

	suite.NotEqualValues(task.Title, taskDB.Title)
	suite.NotEqualValues(task.Body, taskDB.Body)
}
