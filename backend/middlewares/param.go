package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-compose-rest/constant"
)

func GetParamId(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "could not parse id"})
		return
	}
	context.Set(constant.TaskId, id)
	context.Next()
}
