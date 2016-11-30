package store

import (
	"context"

	"github.com/juliengroch/todolist/models"
)

// GetCommentByID get one comment by id
func GetCommentByID(ctx context.Context, id string, userID string) (*models.Comment, error) {
	comment := &models.Comment{}

	return comment, Query(ctx).
		Preload("User").
		Joins("LEFT JOIN task ON comment.task_id = task.id").
		Where("comment.id = ? AND (comment.user_id = ? OR task.user_id = ?)", id, userID, userID).
		Find(comment).Error
}

// FindComments find all taks for one user
func FindComments(ctx context.Context, userID string) ([]models.Comment, error) {
	comments := []models.Comment{}
	return comments, Query(ctx).Preload("User").Where("user_id = ?", userID).Find(&comments).Error
}
