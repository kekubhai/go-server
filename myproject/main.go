package main

import (
	"fmt"
	"myproject/utils" // Import the utils package
	"net/http"
	// Import the utils package if in the same directory
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

var users = map[string]Login{}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) < 8 || len(password) < 8 {
		http.Error(w, "Username and password must be at least 8 characters long", http.StatusNotAcceptable)
		return
	}

	if _, exists := users[username]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	users[username] = Login{
		HashedPassword: hashedPassword,
	}

	fmt.Fprintln(w, "User registered successfully")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login endpoint - Implement authentication here")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout endpoint - Implement session termination here")
}

func protected(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Protected endpoint - Implement authentication check here")
}
