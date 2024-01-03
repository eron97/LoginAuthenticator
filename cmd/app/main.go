package main

import (
	"log"
	"os"

	"github.com/eron97/LoginAuthenticator.git/config/database"
	"github.com/eron97/LoginAuthenticator.git/config/routes"
	"github.com/eron97/LoginAuthenticator.git/config/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente:", err)
	}

	dbCredentials := os.Getenv("DB_CREDENTIALS")

	con, err := database.NewDatabase(dbCredentials)
	if err != nil {
		log.Fatal("Erro ao inicializar o banco de dados:", err)
	}

	defer con.CloseDatabase()

	todoService := services.NewListService(con)

	r := gin.Default()
	routes.SetupTaskRoutes(r, todoService)
	r.Run()

}
