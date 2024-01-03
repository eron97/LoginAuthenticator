package database

import (
	"database/sql"
	"log"

	"github.com/eron97/LoginAuthenticator.git/config/models"
	_ "github.com/go-sql-driver/mysql"
)

type DBRepository interface {
	ReadAll() ([]models.User, error)
}

type Database struct {
	db *sql.DB
}

func (con *Database) ReadAll() ([]models.User, error) {
	rows, err := con.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.BirthDate, &user.PhoneNumber)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func NewDatabase(credentials string) (*Database, error) {
	db, err := sql.Open("mysql", credentials)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("[InitDB] Conexão bem-sucedida com AWS RDS")
	return &Database{db}, nil
}

func (con *Database) CloseDatabase() {
	if con.db != nil {
		con.db.Close()
	}
}
