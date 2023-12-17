package helpers

import (
	"dbo-technical-test/params"
	"math"

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

func CalculatePagination(totalCount int64, currentPage, pageSize, currentCount int) params.PaginationResponse {
	if currentPage == 0 {
		currentPage = 1
	}

	nextPage := currentPage + 1

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	firstPage := 1
	lastPage := totalPages

	if currentPage < firstPage {
		currentPage = firstPage
	}

	if currentPage > lastPage {
		currentPage = lastPage
	}

	if nextPage < firstPage {
		nextPage = firstPage
	}

	if nextPage > lastPage {
		nextPage = lastPage
	}

	return params.PaginationResponse{
		CurrentPage:  currentPage,
		PageSize:     pageSize,
		TotalCount:   totalCount,
		TotalPages:   totalPages,
		FirstPage:    firstPage,
		NextPage:     nextPage,
		LastPage:     lastPage,
		CurrentCount: currentCount,
	}
}
