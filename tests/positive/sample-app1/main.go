package sample_app1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.New()
	server.GET("/health", IsHealthy)
}

func IsHealthy(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusOK, nil)
}
