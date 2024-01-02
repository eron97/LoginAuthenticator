package main

import (
	"log"

	"github.com/eron97/LoginAuthenticator.git/config/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbCredentials := "admin:adminadmin@tcp(database-1.cpj0eavfzshu.us-east-1.rds.amazonaws.com:3306)/users"

	// Inicializa o banco de dados
	db := database.InitDB(dbCredentials)
	log.Println("[Conexão com database ok]")
	defer db.Close()

	r := gin.Default()

	// Armazena o ponteiro do banco de dados no contexto Gin
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.Run()
}
