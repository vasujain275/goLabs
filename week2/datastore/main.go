package main

import (
	"errors"
	"fmt"
)

// Task: Define a DataStore Interface.
// DoD: A DataStore interface is defined with methods like GetUser(id string) (*User, error) and CreateUser(user *User) error.

//  Task: Implement the DataStore Interface with a MemoryStore.
// DoD: A MemoryStore struct (using a map) is created that correctly implements the DataStore interface. The methods handle cases like a user not being found by returning a custom error.

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type DataStore interface {
	GetUser(id string) (*User, error)
	CreateUser(user *User) error
}

type MemoryStore struct {
	Users map[string]*User
}

func (m *MemoryStore) GetUser(id string) (*User, error) {
	user, ok := m.Users[id]

	if ok {
		return user, nil
	} else {
		return nil, errors.New("Can't Find the User with id " + id)
	}
}

func (m *MemoryStore) CreateUser(user *User) error {
	m.Users[user.ID] = user
	return nil
}

func main() {
	// 1. Create a MemoryStore, initializing its map
	store := &MemoryStore{
		Users: make(map[string]*User),
	}

	// 2. Create a User pointer and save it
	u := &User{
		ID:        "1",
		FirstName: "Vasu",
		LastName:  "Jain",
		Email:     "vasu@example.com",
		IsActive:  true,
	}
	if err := store.CreateUser(u); err != nil {
		fmt.Println("Create error:", err)
		return
	}

	// 3a. Retrieve the existing user
	retrieved, err := store.GetUser("1")
	if err != nil {
		fmt.Println("GetUser error:", err)
	} else {
		fmt.Printf("Found user: %+v\n", retrieved)
	}

	// 3b. Attempt to retrieve a non‑existent user
	_, err = store.GetUser("2")
	if err != nil {
		fmt.Println("Expected error:", err)
	}
}
