package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type EventType int

const (
	Unknown EventType = iota
	Created
	Sent
	ApprovedBy
	RejectedBy
	Signed
	Deleted
	Updated
)

func (e EventType) String() string {
	switch e {
	case Unknown:
		return "unknown"
	case Created:
		return "created"
	case Sent:
		return "sent"
	case ApprovedBy:
		return "approved_by"
	case RejectedBy:
		return "rejected_by"
	case Signed:
		return "signed"
	case Deleted:
		return "deleted"
	case Updated:
		return "updated"
	default:
		return "unknown type"
	}
}

type Event struct {
	TaskID    uuid.UUID `json:"task_uuid"`
	UserUUID  uuid.UUID `json:"user_uuid"`
	EventType string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
