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

func LoginHandler(c *gin.Context) {

	var creds models.Credentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid request"})
		return
	}

	db, err := ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	hash := sha256.Sum256([]byte(creds.Password))
	fmt.Println(hex.EncodeToString(hash[:]))
	// user role is an enum in the type USER_ROLE
	rows, err := db.Query("SELECT * FROM BL_USER WHERE email = $1 AND password = $2", creds.Email, hex.EncodeToString(hash[:]))

	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid credentials"})
		return
	}
	// create token
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "token": tokenString})
}
