package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"serveur/server/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	host     = "bouffluence-4322.g95.gcp-us-west2.cockroachlabs.cloud"
	port     = 26257
	user     = "bouffluence"
	password = "gTsPKkviQpqV3wl6JYeiOw"
	dbname   = "bouffluence"
)

var secret string = "boufluence"

func LoginHandler(c *gin.Context) {
	var creds models.Credentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid request"})
		return
	}

	/*
	 * Connect to database
	 */
	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=require", user, password, host, port, dbname)

	fmt.Println("Log: Info BDD : " + psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	hash := sha256.Sum256([]byte(creds.Password))
	_, err = db.Exec("INSERT into \"USER\" (email, password, user_role) VALUES ($1, $2, $3)", creds.Email, hex.EncodeToString(hash[:]), "CLIENT")

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusConflict, gin.H{"status": "user already exist in data base"})
		return
	}
	db.Close()

	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(tokenString)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
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
