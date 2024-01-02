// No package database

package database

import (
	"database/sql"
	"log"

	"github.com/eron97/LoginAuthenticator.git/config/models"
	_ "github.com/go-sql-driver/mysql"
)

// DBService define as operações necessárias no banco de dados.
type DBService interface {
	QueryAllUsers() ([]models.User, error)
}

// DatabaseServiceImpl é a implementação concreta de DBService.
type DatabaseServiceImpl struct {
	DB *sql.DB
}

func (s *DatabaseServiceImpl) QueryAllUsers() ([]models.User, error) {
	rows, err := s.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.BirthDate, &user.PhoneNumber)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func InitDB(dbCredentials string) *sql.DB {
	var err error
	db, err := sql.Open("mysql", dbCredentials)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[InitDB: Conexão bem-sucedida ao AWS RDS]")

	return db
}
