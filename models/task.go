package models

import "time"

// Task model.
type Task struct {
	ID          string    `sql:"type:varchar(255)"`
	Title       string    `sql:"type:varchar(30)"`
	Description string    `sql:"type:varchar(255)"`
	Priority    int8      `sql:"type:integer"`
	UserID      string    `sql:"type:varchar(255)"`
	Created     time.Time `sql:"not null"`
	Modified    time.Time `sql:"not null"`
	User        User
}

// TableName returns the table name.
func (Task) TableName() string {
	return "task"
}

// GetID returns instance ID.
func (t Task) GetID() string {
	return t.ID
}
