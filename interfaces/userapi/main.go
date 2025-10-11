package main

import (
	"fmt"
)

func main() {
	// Inject the repository implementation
	repo := NewInMemoryUserRepo()
	fmt.Println("Repository created:", repo)
	service := NewUserService(repo)

	// Use the service
	service.RegisterUser(User{ID: 1, Name: "Alice", Email: "alice@example.com"})
	service.RegisterUser(User{ID: 2, Name: "Bob", Email: "bob@example.com"})

	users, _ := service.GetAllUsers()
	fmt.Println("All Users:", users)
}

// run this with `go run .` in the userapi directory
