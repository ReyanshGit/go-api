package main

import (
	"fmt"
	"net/http"

	"go-api/config"
	"go-api/handlers"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/create", handlers.CreateUser)
	http.HandleFunc("/users", handlers.GetUsers)

	fmt.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)
}
