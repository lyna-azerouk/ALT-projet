package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"serveur/server/models"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "" //replace with your user_name
	Password = ""
	Dbname   = "data_bouffuence"
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)

	fmt.Println("Log: Info BDD : " + psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	//hash := sha256.Sum256([]byte(creds.Password))
	//_, err = db.Exec("INSERT into authentication (email,password) VALUES ($1, $2)", creds.Email, hex.EncodeToString(hash[:]))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

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
