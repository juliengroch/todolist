package store

import (
	"context"
	"fmt"
	"strings"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/models"
	uuid "github.com/satori/go.uuid"
)

type store struct {
	db *gorm.DB
}

// Store interface to store
type Store interface {
	Migrate() error

	CreateTask(title string, description string, priority int8) (*models.Task, error)
	GetTaskByID(id string) (*models.Task, error)
	GetTasks() ([]models.Task, error)
	Save(out interface{}) error
}

// GetTaskByID get one task by id
func (s *store) GetTaskByID(id string) (*models.Task, error) {
	task := &models.Task{}
	return task, s.db.Where("id = ?", id).Find(&task).Error
}

func (s *store) GetTasks() ([]models.Task, error) {
	tasks := []models.Task{}
	return tasks, s.db.Find(&tasks).Error
}

// CreateTask make a new task
func (s *store) CreateTask(title string, description string, priority int8) (*models.Task, error) {
	task := &models.Task{
		ID:          strings.Replace(uuid.NewV4().String(), "-", "", -1),
		Title:       title,
		Description: description,
		Priority:    priority,
		Created:     time.Now(),
		Modified:    time.Now(),
	}

	return task, s.db.Create(task).Error
}

func (s *store) Save(out interface{}) error {
	return s.db.Save(out).Error
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
	return FromContext(ctx).Migrate()
}

func (s *store) Migrate() error {
	s.db.AutoMigrate(&models.Task{})

	return nil
}
