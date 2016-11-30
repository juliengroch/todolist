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

// ----------------------------------------------------------------------------
// Store
// ----------------------------------------------------------------------------

type store struct {
	db *gorm.DB
}

// Store interface to store
type Store interface {
	DB() *gorm.DB

	Create(out interface{}) error
	Save(out interface{}) error
	Migrate(ctx context.Context) error
	ResetDB(ctx context.Context) error
	Close(ctx context.Context) error
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

func (s *store) DB() *gorm.DB {
	return s.db
}

func (s *store) Save(out interface{}) error {
	return s.db.Save(out).Error
}

func (s *store) Create(out interface{}) error {
	return s.db.Create(out).Error
}

// Close connection
func (s *store) Close(ctx context.Context) error {
	loggers.FromContext(ctx).Info("Close BDD connection")
	return s.db.Close()
}

// ResetDB drop all tables
func (s *store) ResetDB(ctx context.Context) error {
	loggers.FromContext(ctx).Info("All tables from bdd have been dropped")
	return s.db.DropTable(&models.Task{}, &models.User{}, &models.Comment{}).Error
}

func (s *store) Migrate(ctx context.Context) error {
	s.db.AutoMigrate(&models.Task{}, &models.User{}, &models.Comment{})
	loggers.FromContext(ctx).Info("Migrate on BDD ok")
	return nil
}

// ----------------------------------------------------------------------------
// Query helpers
// ----------------------------------------------------------------------------

// Query a value in the store
func Query(ctx context.Context) *gorm.DB {
	return FromContext(ctx).DB()
}

// Create a new resource
func Create(ctx context.Context, out interface{}) error {
	return FromContext(ctx).Create(out)
}

// Save a resource
func Save(ctx context.Context, out interface{}) error {
	return FromContext(ctx).Save(out)
}

// Migrate migrates the schema
func Migrate(ctx context.Context) error {
	return FromContext(ctx).Migrate(ctx)
}
