package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leMoreira/api_simples/cmd/api/bd"
	"github.com/leMoreira/api_simples/cmd/api/models"
)

func Inicio(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func TasksShow(ctx *gin.Context) {
	var tasksShow []models.Task

	bd.DB.Find(&tasksShow)

	ctx.JSON(200, gin.H{
		"dados": tasksShow,
	})

}

func InsertTask(ctx *gin.Context) {
	var task models.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	bd.DB.Create(&task)
	ctx.JSON(http.StatusOK, task)

}

func SelectTaskId(ctx *gin.Context) {
	var task models.Task

	id := ctx.Params.ByName("id")
	bd.DB.First(&task, id)

	if task.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Tarefa não encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, task)

}

func UpdateTask(ctx *gin.Context) {
	var task models.Task

	id := ctx.Params.ByName("id")

	bd.DB.First(&task, id)

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})

		return
	}

	bd.DB.Model(&task).UpdateColumns(task)

	ctx.JSON(http.StatusOK, task)

}

func DeleteTask(ctx *gin.Context) {
	var task models.Task

	id := ctx.Params.ByName("id")

	bd.DB.Delete(&task, id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Tarefa deletada com sucesso!",
	})

}

// Tipos de Consultas por title ou description

func SelectTaskTitle(ctx *gin.Context) {

	title := ctx.Param("title")

	var task = models.Task{Title: title}

	bd.DB.First(&task)

	ctx.JSON(http.StatusOK, task)

}
