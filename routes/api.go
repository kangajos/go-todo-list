package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kangajos/go-todo-list.git/controllers"
	"github.com/kangajos/go-todo-list.git/models"
	"github.com/kangajos/go-todo-list.git/services"
	"gorm.io/gorm"
)

func RouteApi(r *gin.Engine, db gorm.DB) *gin.Engine {

	activityRepo := models.ActivityRepository{}
	activityRepo = *activityRepo.New(db)

	activitySrv := services.ActivityService{}
	activitySrv = *activitySrv.New(&activityRepo)

	activityCtr := controllers.ActivityController{}
	activityCtr = *activityCtr.New(&activitySrv)

	activityGroup := r.Group("/activity-groups")
	{
		activityGroup.GET("/", activityCtr.FindAll)
		activityGroup.GET("/:ID", activityCtr.FindByID)
		activityGroup.POST("/", activityCtr.Create)
		activityGroup.PATCH("/:ID", activityCtr.Update)
		activityGroup.DELETE("/:ID", activityCtr.Delete)
	}

	todoRepo := models.TodoRepository{}
	todoRepo = *todoRepo.New(db)

	todoSrv := services.TodoService{}
	todoSrv = *todoSrv.New(&todoRepo)

	todoCtr := controllers.TodoController{}
	todoCtr = *todoCtr.New(&todoSrv)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to API Todo List :)"})
	})

	todoGroup := r.Group("/todo-items")
	{
		todoGroup.GET("/", todoCtr.FindAll)
		todoGroup.GET("/:ID", todoCtr.FindByID)
		todoGroup.POST("/", todoCtr.Create)
		todoGroup.PATCH("/:ID", todoCtr.Update)
		todoGroup.DELETE("/:ID", todoCtr.Delete)
	}
	return r
}
