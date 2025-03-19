package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-compose-rest/middlewares"
)

func RegisterServer(server *gin.Engine) {
	taskRoute := server.Group("/api/tasks")
	taskRoute.GET("/", getTasks)
	taskRoute.POST("/", createTask)
	taskRoute.GET("/:id", middlewares.GetParamId, getTaskById)
	taskRoute.PUT("/:id", middlewares.GetParamId, updateTaskById)
	taskRoute.DELETE("/:id", middlewares.GetParamId, deleteTask)
}
