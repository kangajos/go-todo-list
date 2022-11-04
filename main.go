package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kangajos/go-todo-list.git/config"
	"github.com/kangajos/go-todo-list.git/routes"
)

func main() {
	db := config.DBConnect()
	r := gin.Default()
	routes.RouteApi(r, *db)
	r.Run(":3030")
}
