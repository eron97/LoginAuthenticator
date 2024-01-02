package database

import (
	"database/sql"
	"log"

	"github.com/eron97/LoginAuthenticator.git/config/models"
	_ "github.com/go-sql-driver/mysql"
)

// UserRepository define a interface para operações de banco de dados associadas à entidade User
type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUserByID(userID int) (models.User, error)
	CreateUser(user models.User) (int, error)
	UpdateUser(user models.User) error
	DeleteUser(userID int) error
}

// MySQLUserRepository é uma implementação de UserRepository para MySQL
type MySQLUserRepository struct {
	db *sql.DB
}

// NewMySQLUserRepository cria uma nova instância de MySQLUserRepository
func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

// GetUsers retorna todos os usuários do banco de dados
func (r *MySQLUserRepository) GetUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
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

// InitDB inicia a conexão com o banco de dados MySQL na AWS RDS
func InitDB(dbCredentials string) (*sql.DB, error) {
	var err error
	db, err := sql.Open("mysql", dbCredentials)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Verifica a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("[InitDB: Conexão bem-sucedida ao AWS RDS]")

	return db, nil
}
