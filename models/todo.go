package models

import (
	"time"

	"gorm.io/gorm"
)

type TodoImpl interface {
	FindAll() ([]Todo, error)
	FindByID(ID int) (*Todo, error)
	Create(Todo) (*Todo, error)
	Update(*Todo) (*Todo, error)
	Delete(ID int) error
}

type Todo struct {
	ID              int             `json:"id"`
	ActivityGroupID int             `json:"activity_group_id"`
	Title           string          `json:"title"`
	IsActive        bool            `json:"is_active"`
	Priority        string          `json:"priority"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       *gorm.DeletedAt `json:"deleted_at"`
}

type TodoRepository struct {
	db gorm.DB
}

func (tr *TodoRepository) New(db gorm.DB) *TodoRepository {
	tr.db = db
	return tr
}

// get all todos
func (tr *TodoRepository) FindAll() ([]Todo, error) {

	todos := []Todo{}
	if err := tr.db.Model(&Todo{}).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// get todo by ID
func (tr *TodoRepository) FindByID(ID int) (*Todo, error) {

	todo := Todo{}
	if err := tr.db.Model(&Todo{}).Where("id = ?", ID).Find(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// create todo
func (tr *TodoRepository) Create(todo Todo) (*Todo, error) {

	if err := tr.db.Create(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// update todo by ID
func (tr *TodoRepository) Update(todo *Todo) (*Todo, error) {

	if err := tr.db.Save(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

// delete todo by ID
func (tr *TodoRepository) Delete(ID int) error {

	if err := tr.db.Where("id = ?", ID).Delete(&Todo{}).Error; err != nil {
		return err
	}
	return nil
}
