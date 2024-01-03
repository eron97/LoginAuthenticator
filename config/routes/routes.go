package routes

import (
	"github.com/eron97/LoginAuthenticator.git/config/controllers"
	"github.com/eron97/LoginAuthenticator.git/config/services"
	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine, todoService services.ListService) {
	controllers.SetTodoService(todoService)

	router.GET("/users", controllers.GetAllController)
}
