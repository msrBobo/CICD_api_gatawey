package v1

import (
	"dennic-api-gateway/api/models/model_common"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HandleError(c *gin.Context, err error, l *zap.Logger, statusCode int, msg string) bool {
	if err == nil {
		return false
	}
	c.JSON(statusCode,
		&model_common.ResponseError{
			Code:    http.StatusText(statusCode),
			Message: msg,
			Data:    err.Error(),
		})
	l.Log(1, err.Error())
	return true
}
