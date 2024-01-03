package services

import (
	"github.com/eron97/LoginAuthenticator.git/config/database"
	"github.com/eron97/LoginAuthenticator.git/config/models"
)

type ListService interface {
	AllUsers() ([]models.User, error)
}

type MyListServices struct {
	DBRepository database.DBRepository
}

func NewListService(DBRepository database.DBRepository) *MyListServices {
	return &MyListServices{
		DBRepository: DBRepository,
	}
}

func (db *MyListServices) AllUsers() ([]models.User, error) {
	return db.DBRepository.ReadAll()
	// restante do processamento da lógica de negócios
	// output
}
