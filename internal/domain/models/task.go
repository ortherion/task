package models

import uuid "github.com/satori/go.uuid"

// Task swagger: model Task
type Task struct {
	ID          uuid.UUID   `db:"id" json:"id"`
	Title       string      `db:"title" json:"title"`
	Body        string      `db:"body" json:"body"`
	CreatorID   uuid.UUID   `db:"creator_id" json:"creator_id"`
	Stage       Stage       `db:"status_task" json:"stage"`
	Signatories []Signatory `json:"signatories"`
	entity
	deletable
}
