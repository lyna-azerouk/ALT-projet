package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const Log = "LOG : "

const (
	host     = "localhost"
	port     = 5432
	user     = "" //replace with your user_name
	password = "52fdc5a882ad0cc490297a43dce208cc36639f0c5224fc47bc849a978bd16d98"
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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
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

	// hash := sha256.Sum256([]byte(creds.Password))
	// _, err = db.Exec("INSERT into authentication (email,password) VALUES ($1, $2)", creds.Email, hex.EncodeToString(hash[:]))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("LOG : BDD Error " + err.Error())
		return
	}

	response := Response{
		Success: 200,
		Message: "Successfully  signUp",
	}

	jsonBytes, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

	fmt.Println(Log + "signUp Success")

}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.ListenAndServe(":8080", nil)
}
