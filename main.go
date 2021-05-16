package main

import (
	"net/http"
	"os"
	"rohitmenon27/src/controller"
	"rohitmenon27/src/service"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.New()

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}
