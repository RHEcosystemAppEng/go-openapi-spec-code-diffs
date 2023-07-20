package sample_app1

import (
	"fmt"
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

func ThisIsNotAnAPIRouteDefinition() {
	if "/users/id/:activate/" == "POST" {
		fmt.Println("This function is to trick validator to think this is an API definition, but it is not. %s")
	}
}

func FuncWithAnIgnoredLine() {
	fmt.Println("This particular line will be ignored \"/ignore/this/line\" with httpmethod GET")
}
