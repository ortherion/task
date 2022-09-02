package models

import uuid "github.com/satori/go.uuid"

type Signatory struct {
	ID     uuid.UUID `db:"id" json:"id"`
	TaskID uuid.UUID `db:"task_id" json:"task_id"`
	Email  string    `db:"email" json:"email"`
	Status Stage     `db:"status_task" json:"status"`

	//Signing []string `db:"email"`
	//Signed  []string
}
