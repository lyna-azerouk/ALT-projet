package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"serveur/server/const/requests"
	"serveur/server/database"
	"serveur/server/models"

	"github.com/gin-gonic/gin"
)

func RegistrationHandler(c *gin.Context) {
	var creds models.ClientCredentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid request"})
		return
	}

	db, err := database.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed to connect to database"})
		return
	}

	hash := sha256.Sum256([]byte(creds.Password))
	query := requests.SelectClientByEmailAndPasswordRequestTemplate
	_, err = db.Exec(query, creds.Email, hex.EncodeToString(hash[:]), "CLIENT")

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"status": "user already exist in data base"})
		return
	}
	db.Close()

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully, you can now login"})
}
