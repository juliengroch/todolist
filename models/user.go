package models

// User model.
type User struct {
	ID       string
	Username string `sql:"not null; unique; type:varchar(150)"`
	APIKey   string `sql:"not null; type:varchar(255)"`
}

// TableName returns the table name.
func (User) TableName() string {
	return "user"
}

// GetID returns instance ID.
func (u User) GetID() string {
	return u.ID
}
