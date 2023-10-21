package model

import (
	"time"
)

type Status string

const (
	StatusTodo       Status = "TODO"
	StatusInProgress Status = "IN_PROGRESS"
	StatusDone       Status = "DONE"
)

type Todo struct {
	ID          string
	Title       string
	Description string
	Status      Status
	TodoDate    time.Time
	CreatedAt   time.Time
}

func (t Todo) TableName() {

}

type Assssss struct {
}
