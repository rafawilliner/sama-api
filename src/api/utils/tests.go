package utils

import "github.com/gin-gonic/gin"

func GetTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	return router
}