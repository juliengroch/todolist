package store

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/loggers"
	"github.com/juliengroch/todolist/models"
)

type store struct {
	db *gorm.DB
}

// Store interface to store
type Store interface {
	Migrate(ctx context.Context) error

	GetTaskByID(id string, userID string) (*models.Task, error)
	FindTasks(userID string) ([]models.Task, error)
	Create(out interface{}) error
	Save(out interface{}) error

	GetUserByKey(key string) (*models.User, error)
}

// GetTaskByID get one task by id
func (s *store) GetTaskByID(id string, userID string) (*models.Task, error) {
	task := &models.Task{}
	return task, s.db.Preload("User").Where("id = ? AND user_id = ?", id, userID).Find(&task).Error
}

func (s *store) FindTasks(userID string) ([]models.Task, error) {
	tasks := []models.Task{}
	return tasks, s.db.Preload("User").Where("user_id = ?", userID).Find(&tasks).Error
}

// GetUserByKey get user by username
func (s *store) GetUserByKey(key string) (*models.User, error) {
	user := &models.User{}
	return user, s.db.Where("api_key = ?", key).Find(&user).Error
}

func (s *store) Save(out interface{}) error {
	return s.db.Save(out).Error
}

func (s *store) Create(out interface{}) error {
	return s.db.Create(out).Error
}

// New init store
func New(cfg config.Database) (*store, error) {
	// open connection
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", cfg.Host, cfg.User, cfg.Name, cfg.Password)

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	return &store{db}, err
}

// Migrate migrates the schema
func Migrate(ctx context.Context) error {
	return FromContext(ctx).Migrate(ctx)
}

func (s *store) Migrate(ctx context.Context) error {
	s.db.AutoMigrate(&models.Task{}, &models.User{})
	loggers.FromContext(ctx).Info("Migrate on BDD ok")
	return nil
}
