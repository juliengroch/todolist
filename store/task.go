package store

import (
	"context"

	"github.com/juliengroch/todolist/models"
)

// GetTaskByID get one task by id
func GetTaskByID(ctx context.Context, id string, userID string) (*models.Task, error) {
	task := &models.Task{}
	return task, Query(ctx).
		Preload("User").Preload("Comments").Preload("Comments.User").
		Where("id = ? AND user_id = ?", id, userID).
		Find(task).Error
}

// FindTasks find all taks for one user
func FindTasks(ctx context.Context, userID string) ([]models.Task, error) {
	tasks := []models.Task{}
	return tasks, Query(ctx).
		Preload("User").Preload("Comments").Preload("Comments.User").
		Where("user_id = ?", userID).
		Find(&tasks).Error
}
