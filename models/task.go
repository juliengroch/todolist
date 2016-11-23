package models

import "time"

// Task model.
type Task struct {
	ID          string `sql:"type:varchar(255)"`
	Title       string `sql:"type:varchar(30)"`
	Description string `sql:"type:varchar(255)"`
	Priority    int8
	Created     time.Time `sql:"not null"`
	Modified    time.Time `sql:"not null"`
}

// TableName returns the table name.
func (Task) TableName() string {
	return "task"
}

// GetID returns instance ID.
func (t Task) GetID() string {
	return t.ID
}
