package managers

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/payloads"
)

// CreateTasks create a task
func CreateTasks(ctx context.Context, payload *payloads.Task) (models.Task, error) {
	var err error
	modelTask := models.Task{
		ID:          1,
		Title:       payload.Title,
		Description: payload.Description,
	}

	// open connection
	db, err := gorm.Open("sqlite3", "todolist.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// // Migrate the schema
	// db.AutoMigrate(&modelTask{})

	// // Create
	// db.Create(&modelTask{})

	return modelTask, err
}
