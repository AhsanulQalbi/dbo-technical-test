package router

import (
	"dbo-technical-test/middlewares"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RouterConfig(db *gorm.DB) *gin.Engine {
	route := gin.Default()
	logger := log.New()
	route.Use(middlewares.ErrorHandler(logger))

	return route
}
