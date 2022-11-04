package services

import (
	"errors"
	"time"

	"github.com/kangajos/go-todo-list.git/models"
	"github.com/kangajos/go-todo-list.git/request"
)

type TodoServiceImpl interface {
	FindAll() ([]models.Todo, error)
	FindByID(ID int) (*models.Todo, error)
	Create(request.TodoRequest) (*models.Todo, error)
	Update(models.Todo) (*TodoService, error)
	Delete(ID int) error
}

type TodoService struct {
	todoRepository *models.TodoRepository
}

// init
func (ts *TodoService) New(tr *models.TodoRepository) *TodoService {
	ts.todoRepository = tr
	return ts
}

// get all todo
func (ts TodoService) FindAll() ([]models.Todo, error) {
	return ts.todoRepository.FindAll()
}

// get all todo
func (ts TodoService) FindByID(ID int) (*models.Todo, error) {
	data, _ := ts.todoRepository.FindByID(ID)
	if data.ID == 0 {
		return nil, errors.New("Activitiy ID #{ID} Not Found")
	}
	return data, nil
}

// create todo
func (ts TodoService) Create(todoRequest request.TodoRequest) (*models.Todo, error) {
	todo := models.Todo{}
	todo.ActivityGroupID = todoRequest.ActivityGroupID
	todo.Title = todoRequest.Title
	todo.Priority = todoRequest.Priority
	todo.IsActive = todoRequest.IsActive
	return ts.todoRepository.Create(todo)
}

// update todo
func (ts TodoService) Update(todoRequest request.TodoRequest) (*models.Todo, error) {
	todo, err := ts.todoRepository.FindByID(todoRequest.ID)
	if err != nil {
		return nil, err
	}
	todo.ActivityGroupID = todoRequest.ActivityGroupID
	todo.Title = todoRequest.Title
	todo.Priority = todoRequest.Priority
	todo.IsActive = todoRequest.IsActive
	todo.UpdatedAt = time.Now()
	return ts.todoRepository.Update(todo)
}

// delete todo
func (ts TodoService) Delete(ID int) error {
	data, _ := ts.todoRepository.FindByID(ID)
	if data.ID == 0 {
		return errors.New("Activitiy ID #{ID} Not Found")
	}

	ts.todoRepository.Delete(ID)
	return nil
}
