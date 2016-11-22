package models

import "time"

// Task model.
type Task struct {
	ID            int64
	Provider      string    `sql:"type:varchar(30)"`
	PaymentMethod string    `sql:"type:varchar(255)"`
	Created       time.Time `sql:"not null"`
	Modified      time.Time `sql:"not null"`
}

// TableName returns the table name.
func (Task) TableName() string {
	return "task"
}

// GetID returns instance ID.
func (t Task) GetID() int64 {
	return t.ID
}
