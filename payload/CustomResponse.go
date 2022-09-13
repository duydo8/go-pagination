package payload

import (
	"github.com/gin-gonic/gin"
	"go/types"
)

func CustomResponse(ctx *gin.Context, code int, message string, data types.Object) {
	ctx.JSON(code, gin.H{
		"data":    data,
		"message": message,
		"status":  code,
	})
}
