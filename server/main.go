package main

import (
    "encoding/json"
    "net/http"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
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

    // verfy if the user exist in BDD  
    //  !!! jwt pour l'autehtification !!
    //
    response := map[string]string{"message": "Login successful"}
    json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/login", loginHandler)
    http.Handle("/", http.FileServer(http.Dir("./client")))
    http.ListenAndServe(":8080", nil)
}
