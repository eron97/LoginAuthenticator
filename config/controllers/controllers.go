package controllers

import (
	"github.com/eron97/LoginAuthenticator.git/config/models"
	"github.com/eron97/LoginAuthenticator.git/config/services"
)

type UserController struct {
	userService *services.UserService
}

// NewUserController cria uma nova instância de UserController
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// GetUsersHandler é um manipulador que chama a função GetUsers do UserService
func (controller *UserController) GetUsersHandler() ([]models.User, error) {
	// Chama a função GetUsers do UserService
	users, err := controller.userService.GetUsers()
	if err != nil {
		// Handle error (log, return a specific error, etc.)
		return nil, err
	}

	// Pode realizar outras operações ou manipulações aqui, se necessário

	return users, nil
}
