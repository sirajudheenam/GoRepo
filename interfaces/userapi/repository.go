package main

import "fmt"

// InMemoryUserRepo implements the interface UserRepository because it defines all required methods.

// UserRepository defines an interface for user data operations
type UserRepository interface {
	Create(user User) error
	GetByID(id int) (User, error)
	GetAll() ([]User, error)
}

// Implementation: InMemoryUserRepo
type InMemoryUserRepo struct {
	data map[int]User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{data: make(map[int]User)}
}

func (repo *InMemoryUserRepo) Create(user User) error {
	repo.data[user.ID] = user
	fmt.Println("User added:", user.Name)
	return nil
}

func (repo *InMemoryUserRepo) GetByID(id int) (User, error) {
	user, exists := repo.data[id]
	if !exists {
		return User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (repo *InMemoryUserRepo) GetAll() ([]User, error) {
	users := []User{}
	for _, u := range repo.data {
		users = append(users, u)
	}
	return users, nil
}

// You could later add:
// type PostgresUserRepo struct { db *sql.DB }
