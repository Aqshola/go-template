package common_parsing

import "github.com/gin-gonic/gin"

func JSONResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}
