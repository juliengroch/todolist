package store

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"

	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/loggers"
	"github.com/juliengroch/todolist/models"
)

func getLogger() loggers.Logger {
	return loggers.GetLogger("store.go")
}

type store struct {
	db *gorm.DB
}

// Store interface to store
type Store interface {
	Migrate() error

	CreateTask(title string, description string, priority int8) (*models.Task, error)
	GetTaskByID(id string) (*models.Task, error)
	FindTasks() ([]models.Task, error)
	Save(out interface{}) error

	GetUserByKey(key string) (*models.User, error)
}

// GetTaskByID get one task by id
func (s *store) GetTaskByID(id string) (*models.Task, error) {
	task := &models.Task{}
	return task, s.db.Where("id = ?", id).Find(&task).Error
}

func (s *store) FindTasks() ([]models.Task, error) {
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

// GetUserByKey get user by username
func (s *store) GetUserByKey(key string) (*models.User, error) {
	user := &models.User{}
	return user, s.db.Where("api_key = ?", key).Find(&user).Error
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
	s.db.AutoMigrate(&models.Task{}, &models.User{})
	getLogger().Info("Migrate on BDD ok")
	return nil
}
