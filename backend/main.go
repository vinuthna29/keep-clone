package main

import (
	"keep-clone-backend/api"
	"keep-clone-backend/db"
	"log"

	// "net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	router := gin.Default()

	// router.Static("/", "./frontend")
	router.Use(cors.Default())

	router.POST("/notes", api.CreateNote)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
