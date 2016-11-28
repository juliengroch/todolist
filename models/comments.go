package models

import "time"

// Comment model
type Comment struct {
	ID       string    `sql:"type:varchar(255)"`
	Message  string    `sql:"type:varchar(255)"`
	UserID   string    `sql:"type:varchar(255)"`
	TaskID   string    `sql:"type:varchar(255)"`
	Created  time.Time `sql:"not null"`
	Modified time.Time `sql:"not null"`
	User     User
}

// TableName returns the table name.
func (Comment) TableName() string {
	return "comment"
}

// GetID returns instance ID.
func (c Comment) GetID() string {
	return c.ID
}
