package controllers

import (
	"net/http"

	"github.com/eron97/LoginAuthenticator.git/config/services"
	"github.com/gin-gonic/gin"
)

var listServices services.ListService

func SetTodoService(services services.ListService) {
	listServices = services
}

func GetAllController(c *gin.Context) {
	getAll, err := listServices.AllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": getAll})
}
