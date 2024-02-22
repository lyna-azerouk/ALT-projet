package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"serveur/server/models"
	services "serveur/server/services/jwt"
	"time"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid credentials"})
		return
	}

	// user auth is success, create a new token valid for 30 min
	userClaims := models.UserClaims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	signedAccessToken, err := services.NewAccessToken(userClaims)
	c.JSON(http.StatusOK, gin.H{"token": signedAccessToken})
}
