package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	roles "serveur/server/const"
	"serveur/server/const/requests"
	"serveur/server/database"
	"serveur/server/models"
	services "serveur/server/services/jwt"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

func ClientLoginHandler(c *gin.Context) {

	var creds models.ClientCredentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid request"})
		return
	}

	db, err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	hash := sha256.Sum256([]byte(creds.Password))

	row, err := db.Query(
		requests.SelectClientByEmailAndPasswordRequestTemplate,
		creds.Email, hex.EncodeToString(hash[:]))

	var email, password, role string

	if row.Next() {
		if err := row.Scan(&email, &password, &role); err != nil {
			log.Fatal(err)
		}
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid credentials"})
		return
	}

	// user auth is success, create a new token valid for 30 min
	clientClaims := buildClientCredential(creds, actualRole(role))

	signedAccessToken, err := services.NewClientAccessToken(clientClaims)
	c.JSON(http.StatusOK, gin.H{"token": signedAccessToken})
}

func actualRole(role string) string {
	println("role: " + role)
	if role == "CLIENT" {
		return roles.ClientRole
	} else {
		return roles.AdminRole
	}

}

func RestaurantLoginHandler(c *gin.Context) {
	var creds models.RestaurantCredentials

	if err := c.BindJSON(&creds); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid request"})
		return
	}

	db, err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	hash := sha256.Sum256([]byte(creds.Password))

	_, err = db.Query(
		requests.SelectRestaurantByIdAndPasswordRequestTemplate,
		creds.Id, hex.EncodeToString(hash[:]))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid credentials"})
		return
	}

	// user auth is success, create a new token valid for 30 min
	restaurantClaims := buildRestaurantCredential(creds)

	signedAccessToken, err := services.NewRestaurantAccessToken(restaurantClaims)
	c.JSON(http.StatusOK, gin.H{"token": signedAccessToken})
}

func buildClientCredential(creds models.ClientCredentials, role string) models.ClientClaims {
	userClaims := models.ClientClaims{
		Email: creds.Email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	return userClaims
}

func buildRestaurantCredential(creds models.RestaurantCredentials) models.RestaurantClaims {
	restaurantClaims := models.RestaurantClaims{
		Id: creds.Id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	return restaurantClaims
}
