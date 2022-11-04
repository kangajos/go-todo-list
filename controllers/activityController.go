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

type ActivityController struct {
	activityService *services.ActivityService
	response        *helpers.Response
}

// init
func (ac *ActivityController) New(as *services.ActivityService) *ActivityController {
	ac.activityService = as
	ac.response = &helpers.Response{}
	return ac
}

// get all acitvity
func (ac *ActivityController) FindAll(ctx *gin.Context) {
	data, _ := ac.activityService.FindAll()
	response := ac.response.Create(helpers.STATUS_SUCCESS, helpers.MESSAGE_SUCCESS, data)
	ctx.JSON(http.StatusOK, response)
}

// get activity by id
func (ac *ActivityController) FindByID(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("ID"))
	data, err := ac.activityService.FindByID(ID)
	if err != nil {
		message := fmt.Sprintf("Activity with ID %d Not Found", ID)
		response := ac.response.Create(helpers.STATUS_NOT_FOUND, message, gin.H{})
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := ac.response.Create(helpers.STATUS_SUCCESS, helpers.MESSAGE_SUCCESS, data)
	ctx.JSON(http.StatusOK, response)
}

// create activity
func (ac *ActivityController) Create(ctx *gin.Context) {
	data := request.ActivityRequest{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, ac.response.Create("error", err.Error(), gin.H{}))
		return
	}

	create, err := ac.activityService.Create(data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ac.response.Create("error", err.Error(), gin.H{}))
		return
	}

	ctx.JSON(http.StatusCreated, ac.response.Create(helpers.STATUS_SUCCESS, helpers.MESSAGE_SUCCESS, create))

}

// create activity
func (ac *ActivityController) Update(ctx *gin.Context) {
	data := request.ActivityRequest{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, ac.response.Create("error", err.Error(), gin.H{}))
		return
	}

	ID, _ := strconv.Atoi(ctx.Param("ID"))
	data.ID = ID

	update, err := ac.activityService.Update(data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ac.response.Create("error", err.Error(), gin.H{}))
		return
	}

	ctx.JSON(http.StatusOK, ac.response.Create(helpers.STATUS_SUCCESS, helpers.MESSAGE_SUCCESS, update))

}

// delete activity by id
func (ac *ActivityController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("ID"))

	if err := ac.activityService.Delete(ID); err != nil {
		message := fmt.Sprintf("Activity with ID %d Not Found", ID)
		response := ac.response.Create(helpers.STATUS_NOT_FOUND, message, gin.H{})
		ctx.JSON(http.StatusOK, response)
		return
	}

	ctx.JSON(http.StatusOK, ac.response.Create(helpers.STATUS_SUCCESS, helpers.MESSAGE_SUCCESS, gin.H{}))

}
