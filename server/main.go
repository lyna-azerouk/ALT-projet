package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

const Log = "LOG : "

var secret string = "boufluence"

const (
	host     = "localhost"
	port     = 5432
	user     = "" //replace with your user_name
	password = ""
	dbname   = "data_base_test"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Response struct {
	Success int    `json:"success"`
	Message string `json:"message,omitempty"`
}

func loginHandler(c *gin.Context) {
	var creds Credentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid request"})
		return
	}

	/*
	 * Connect to database
	 */
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println(Log + "Info BDD : " + psqlInfo)

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
	c.JSON(http.StatusCreated, gin.H{"status": "Succesfull", "token": tokenString})
}

func main() {
	router := gin.Default()
	router.Static("/", "./client")

	router.POST("/login", loginHandler)
	router.Run(":8080")

}
