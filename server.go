package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Memory storage for usernames
var usernames []string

// Handler to add username
func addUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read JSON input
	var req struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Store username
	usernames = append(usernames, req.Username)

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "User added successfully",
		"username": req.Username,
	})
}

// Handler to list all usernames
func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usernames)
}

func main() {
	// Routes
	http.HandleFunc("/add", addUserHandler)
	http.HandleFunc("/list", listUsersHandler)

	port := "9090" // using port 9090
	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, nil)
}
