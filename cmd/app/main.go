package main

import (
	"log"

	"github.com/eron97/LoginAuthenticator.git/config/database"
	"github.com/eron97/LoginAuthenticator.git/config/services"

	"github.com/gin-gonic/gin"
)

// No arquivo main.go

func main() {
	dbCredentials := "admin:adminadmin@tcp(database-1.cpj0eavfzshu.us-east-1.rds.amazonaws.com:3306)/users"

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// Passa as credenciais diretamente para a função InitDB
		db := database.InitDB(dbCredentials)
		c.Set("db", db)
		defer db.Close()
		c.Next()
	})

	log.Println("[Servidor iniciado]")

	r.GET("/users", func(c *gin.Context) {
		// Chama a função QueryAllUsers diretamente aqui
		users, err := services.QueryAllUsers(c)
		if err != nil {
			// Lida com o erro, se necessário
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Retorna os usuários
		c.JSON(200, users)
	})

	r.Run()
}
