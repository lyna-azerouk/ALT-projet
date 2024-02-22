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
	host     = "frost-hippo-13790.8nj.gcp-europe-west1.cockroachlabs.cloud"
	port     = 26257
	user     = "bouffluence"
	password = "NKi9yHEPNbAY-_MrwE8IRw"
	dbname   = "defaultdb"
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
	_, err = db.Exec("INSERT into BL_USER (email, password, user_role) VALUES ($1, $2, $3)", creds.Email, hex.EncodeToString(hash[:]), "CLIENT")

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
