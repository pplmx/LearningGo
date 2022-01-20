package demo

import "github.com/google/uuid"

type Task struct {
	Id          uuid.UUID
	CreatedTime int64
	UpdatedTime int64
	DeletedTime int64
}
