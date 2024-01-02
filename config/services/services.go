package services

import (
	"github.com/eron97/LoginAuthenticator.git/config/models"

	"github.com/eron97/LoginAuthenticator.git/config/database"
)

// GetUsers retorna todos os usuários e aplica regras de negócios se necessário
func GetUsers(userRepository database.UserRepository) ([]models.User, error) {
	// Obter usuários do repositório
	users, err := userRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	// Aplicar regras de negócios, se necessário
	// Exemplo: Filtrar usuários, aplicar transformações, etc.

	return users, nil
}
