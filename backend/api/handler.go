package api

import (
	"keep-clone-backend/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateNote(c *gin.Context) {
	var note Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := db.CreateNote(note.Title, note.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	note.Id = id
	c.JSON(http.StatusCreated, note)
}
