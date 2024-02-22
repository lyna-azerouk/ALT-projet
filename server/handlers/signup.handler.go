package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"serveur/server/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secret string = "boufluence"

func RegistrationHandler(c *gin.Context) {
	var creds models.Credentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid request"})
		return
	}

	/*
	 * Connect to database
	 */
	db, err := ConnectDB()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed to connect to database"})
		return
	}

	hash := sha256.Sum256([]byte(creds.Password))
	_, err = db.Exec("INSERT into BL_USER (email, password, user_role) VALUES ($1, $2, $3)", creds.Email, hex.EncodeToString(hash[:]), "CLIENT")

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusConflict, gin.H{"status": "user already exist in data base"})
		return
	}
	db.Close()

	c.JSON(http.StatusOK, gin.H{"message": "ok signup"})
}

func Login_validation(c *gin.Context) {

	var code = c.Param("code_validation")
	fmt.Print(code)
	// verify if the code exist in bdd  and is correct

	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(tokenString)
	c.JSON(http.StatusCreated, gin.H{"token": tokenString})
}
