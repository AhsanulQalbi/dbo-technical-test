package helpers

import (
	"dbo-technical-test/params"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleErrorController(c *gin.Context, statusCode int, errorMsg string) {
	log.Errorln(ERROR_STRING, errorMsg)

	c.AbortWithStatusJSON(statusCode, params.Response{
		Status:  statusCode,
		Payload: errorCustomMessage(statusCode, errorMsg),
	})
}

func HandleErrorService(statusCode int, errMsg string) *params.Response {
	log.Errorln(ERROR_STRING, errMsg)

	return &params.Response{
		Status: statusCode,
		Payload: gin.H{
			"error": errorCustomMessage(statusCode, errMsg),
		},
	}
}

func errorCustomMessage(statusCode int, errMsg string) string {
	if statusCode >= 500 {
		return APPLICATION_ERROR
	}
	return errMsg
}
