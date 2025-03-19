package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-compose-rest/constant"
	"github.com/go-compose-rest/models"
)

func getTasks(context *gin.Context) {
	tasks, error := models.GetTasks()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not query task"})
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func createTask(context *gin.Context) {
	var task models.Task
	error := context.ShouldBindBodyWithJSON(&task)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	error = task.Save()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create task"})
		return
	}
	context.JSON(http.StatusCreated, task)
}

func getTaskById(context *gin.Context) {
	taskId := context.GetInt64(constant.TaskId)
	task, error := models.GetTaskById(taskId)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch task, please try again later"})
		return
	}
	context.JSON(http.StatusOK, task)
}

func updateTaskById(context *gin.Context) {
	taskId := context.GetInt64(constant.TaskId)

	task, error := models.GetTaskById(taskId)
	if error != nil || task == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch task, please try again later"})
		return
	}
	var updateTask models.Task
	error = context.ShouldBindBodyWithJSON(&updateTask)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	updateTask.ID = taskId
	error = updateTask.UpdateTask()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
}

func deleteTask(context *gin.Context) {
	taskId := context.GetInt64(constant.TaskId)
	task, error := models.GetTaskById(taskId)
	if error != nil || task == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch task, please try again later"})
		return
	}

	error = task.DeleteTask()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
