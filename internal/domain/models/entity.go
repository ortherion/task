package models

import "time"

type entity struct {
	CreatedDate time.Time `db:"created_at" json:"created_date"`
	UpdatedDate time.Time `db:"updated_at" json:"updated_date"`
	DeletedDate time.Time `db:"deleted_at" json:"deleted_date"`
}

type deletable struct {
	IsDeleted bool `db:"is_deleted"`
}
