// services/user_service.go
package services

import (
	"github.com/eron97/LoginAuthenticator.git/config/database"
	"github.com/eron97/LoginAuthenticator.git/config/models"
)

type UserService struct {
	userRepository database.UserRepository
}

// NewUserService cria uma nova instância de UserService
func NewUserService(userRepository database.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// GetUsers retorna todos os usuários e aplica regras de negócios se necessário
func (userService *UserService) GetUsers() ([]models.User, error) {
	// Obter usuários do repositório
	users, err := userService.userRepository.DBGetUsers()
	if err != nil {
		return nil, err
	}

	// Aplicar regras de negócios, se necessário
	// Exemplo: Filtrar usuários, aplicar transformações, etc.

	return users, nil
}
