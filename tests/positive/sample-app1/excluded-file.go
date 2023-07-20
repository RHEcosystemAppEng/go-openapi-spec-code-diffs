package sample_app1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusOK, nil)
}

func RegisterRoutes() {
	server := gin.New()
	server.GET("/excluded-hello", HelloWorld)
}
