package main

import "fmt"

// UserService depends on the UserRepository interface
type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) RegisterUser(user User) error {
	fmt.Println("Registering user:", user.Name)
	return s.repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}

// ✅ The service layer doesn’t know or care about how users are stored.
// It just calls repo.Create() or repo.GetAll().

// This is dependency injection using an interface.
