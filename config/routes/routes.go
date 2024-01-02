package routes

import (
	"github.com/eron97/LoginAuthenticator.git/config/controllers"
	"github.com/eron97/LoginAuthenticator.git/config/services"
	"github.com/gin-gonic/gin"
)

// SetupTaskRoutes configura as rotas para o UserController
func SetupTaskRoutes(router *gin.Engine, userService *services.UserService) {
	userController := controllers.NewUserController(userService)

	router.GET("/AllUsers", func(c *gin.Context) {
		// Chama o manipulador GetUsersHandler do UserController
		users, err := userController.GetUsersHandler()
		if err != nil {
			// Handle error (log, return a specific error, etc.)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		// Retorna os usuários como resposta
		c.JSON(200, users)
	})
}
