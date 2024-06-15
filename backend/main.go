package main

import (
	"fmt"
	"keep-clone-backend/api"
	"keep-clone-backend/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("In main")

	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	router := gin.Default()
	router.POST("/notes", api.CreateNote)
	router.Run(":8080")

}
