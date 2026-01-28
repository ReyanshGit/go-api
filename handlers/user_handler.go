package handlers

import (
	"fmt"
	"net/http"

	"go-api/config"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")

	if id == "" || name == "" || email == "" {
		http.Error(w, "Invalid input", 400)
		return
	}

	_, err := config.DB.Exec(
		"INSERT INTO users(id,name,email) VALUES(?,?,?)",
		id, name, email,
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintln(w, "User created")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id,name,email FROM users")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		found = true
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		fmt.Fprintf(w, "%d %s %s\n", id, name, email)
	}

	if !found {
		fmt.Fprintln(w, "No users found")
	}
}
