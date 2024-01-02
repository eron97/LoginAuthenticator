// No package services

package services

import (
	"database/sql"
	"log"

	"github.com/eron97/LoginAuthenticator.git/config/database"
	"github.com/eron97/LoginAuthenticator.git/config/models"
	"github.com/gin-gonic/gin"
)

// QueryAllUsers chama a função correspondente do DatabaseService para obter todos os usuários.
func QueryAllUsers(c *gin.Context) ([]models.User, error) {
	// Obtém o ponteiro do banco de dados do contexto Gin
	db, exists := c.Get("db")
	if !exists {
		log.Fatal("[GetDB: Ponteiro não está no contexto Gin]")
	}

	// Cria uma instância do DatabaseServiceImpl
	dbService := &database.DatabaseServiceImpl{DB: db.(*sql.DB)}

	// Chama a função QueryAllUsers do pacote database
	return dbService.QueryAllUsers()
}
