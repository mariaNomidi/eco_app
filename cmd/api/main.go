package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	db "github.com/mariaNomidi/eco-app/internal/db"
	"github.com/mariaNomidi/eco-app/internal/handler"
	"github.com/mariaNomidi/eco-app/internal/repository"
	"github.com/mariaNomidi/eco-app/internal/service"
)

func main() {
	//setup DB
	err := godotenv.Load()
	if err != nil {
		log.Println("could not load .env")
	}

	dbURL := os.Getenv("DB_URL")

	database, err := db.NewPostgresDB(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	dbPool := database.Pool
	//setup DB

	//setup layers
	repo := repository.NewIncidentTypeRepo(dbPool)
	service := service.NewIncidentTypeService(repo)
	handler := handler.NewIncidentTypeHandler(service)
	//setup layers

	r := gin.Default()

	//map routes
	r.GET("/incident-types", handler.GetAll)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}
