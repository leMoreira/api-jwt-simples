package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/leMoreira/api_simples/cmd/api/controllers"
	"github.com/leMoreira/api_simples/cmd/api/middlewares"
)

func Routes() {
	g := gin.Default()

	// Authenticate
	g.POST("/auth/signup", controllers.Createuser)
	g.POST("/auth/login", controllers.Login)
	g.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)

	// Rotas simples crud
	g.GET("/", controllers.Inicio)
	g.GET("/tasks", controllers.TasksShow)
	g.POST("/task", controllers.InsertTask)
	g.GET("/tasks/:id", controllers.SelectTaskId)
	g.PUT("/task/:id", controllers.UpdateTask)
	g.DELETE("/task/:id", controllers.DeleteTask)

	//Tipos de Consultas com Title e Description
	g.GET("/task/:title", controllers.SelectTaskTitle)

	g.Run(":" + os.Getenv("PORT"))
}
