package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"serveur/server/const/affluences"
	"serveur/server/const/requests"
	"serveur/server/database"
	"serveur/server/models"
	"serveur/server/services"

	"github.com/gin-gonic/gin"
)

func ClientRegistrationHandler(c *gin.Context) {
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
	query := requests.InsertNewClientRequestTemplate
	_, err = db.Exec(query, creds.Email, hex.EncodeToString(hash[:]), "CLIENT", creds.FirstName, creds.LastName)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"status": "user already exist in data base"})
		return
	}
	db.Close()

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully, you can now login"})
}

func RestaurantRegistrationHandler(c *gin.Context) {
	var creds models.RestaurantCredentials
	var affluence string = affluences.MEDIUM_AFFLUENCE // default value
	var averageOrderDuration int = 30

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
	query := requests.InsertNewRestaurantRequestTemplate
	_, err = db.Exec(query, creds.Id, creds.Name, hex.EncodeToString(hash[:]), affluence, averageOrderDuration)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"status": "Restaurant already exist in data base"})
		return
	}

	// initialize affluence
	err = services.ResetAffluence(creds.Id)
	if err != nil {
		log.Println("Error:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed to initialize affluence"})
		return
	}

	db.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Restaurant created successfully, you can now login"})
}
