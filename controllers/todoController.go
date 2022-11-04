package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kangajos/go-todo-list.git/helpers"
	"github.com/kangajos/go-todo-list.git/request"
	"github.com/kangajos/go-todo-list.git/services"
)

type TodoController struct {
	todoService *services.TodoService
	response    *helpers.Response
}

// init
func (ac *TodoController) New(as *services.TodoService) *TodoController {
	ac.todoService = as
	ac.response = &helpers.Response{}
	return ac
}

// get all acitvity
func (ac *TodoController) FindAll(ctx *gin.Context) {
	data, _ := ac.todoService.FindAll()
	response := ac.response.Create(helpers.STATUS_SUCCESS, helpers.MESSAGE_SUCCESS, data)
	ctx.JSON(http.StatusOK, response)
}

// get todo by id
func (ac *TodoController) FindByID(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("ID"))
	data, err := ac.todoService.FindByID(ID)
	if err != nil {
		message := fmt.Sprintf("Todo with ID %d Not Found", ID)
		response := ac.response.Create(helpers.STATUS_NOT_FOUND, message, gin.H{})
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := ac.response.Create(helpers.STATUS_SUCCESS, helpers.MESSAGE_SUCCESS, data)
	ctx.JSON(http.StatusOK, response)
}

// create todo
func (ac *TodoController) Create(ctx *gin.Context) {
	data := request.TodoRequest{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, ac.response.Create("error", err.Error(), gin.H{}))
		return
	}

	create, err := ac.todoService.Create(data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ac.response.Create("error", err.Error(), gin.H{}))
		return
	}

	ctx.JSON(http.StatusCreated, ac.response.Create(helpers.STATUS_SUCCESS, helpers.MESSAGE_SUCCESS, create))

}

// create todo
func (ac *TodoController) Update(ctx *gin.Context) {
	data := request.TodoRequest{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, ac.response.Create("error", err.Error(), gin.H{}))
		return
	}

	ID, _ := strconv.Atoi(ctx.Param("ID"))
	data.ID = ID

	update, err := ac.todoService.Update(data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ac.response.Create("error", err.Error(), gin.H{}))
		return
	}

	ctx.JSON(http.StatusOK, ac.response.Create(helpers.STATUS_SUCCESS, helpers.MESSAGE_SUCCESS, update))

}

// delete todo by id
func (ac *TodoController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("ID"))

	if err := ac.todoService.Delete(ID); err != nil {
		message := fmt.Sprintf("Todo with ID %d Not Found", ID)
		response := ac.response.Create(helpers.STATUS_NOT_FOUND, message, gin.H{})
		ctx.JSON(http.StatusOK, response)
		return
	}

	ctx.JSON(http.StatusOK, ac.response.Create(helpers.STATUS_SUCCESS, helpers.MESSAGE_SUCCESS, gin.H{}))

}
