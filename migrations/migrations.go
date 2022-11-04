package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Activities struct {
	ID        int
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
}

type Todos struct {
	ID              int
	ActivityGroupID int
	Title           string
	IsActive        bool
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *gorm.DeletedAt
}
