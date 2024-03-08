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

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid credentials"})
		return
	}

	var clientClaims models.ClientClaims
	var role string
	for row.Next() {
		err := row.Scan(&clientClaims.Id, &clientClaims.Email, &role)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid credentials"})
			return
		}
	}

	// user auth is success, create a new token valid for 30 min
	clientClaims = buildClientCredential(creds, actualRole(role), clientClaims.Id)

	signedAccessToken, err := services.NewClientAccessToken(clientClaims)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedAccessToken, "id": clientClaims.Id})
}

func actualRole(role string) string {
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
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedAccessToken, "id": creds.Id})
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

func buildClientCredential(creds models.ClientCredentials, role string, id uint64) models.ClientClaims {
	userClaims := models.ClientClaims{
		Id:    id,
		Email: creds.Email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	return userClaims
}
