package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IsActive  bool   `json:"is_active"`
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

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := HealthResponse{Status: "ok"}

	jsonBytes, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Write(jsonBytes)
}

func UserHandler(store DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		memStore, ok := store.(*MemoryStore)
		if !ok {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Convert map to slice
		users := []*User{}
		for _, user := range memStore.Users {
			users = append(users, user)
		}

		jsonBytes, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)

	}
}

func main() {
	m := http.NewServeMux()

	store := &MemoryStore{
		Users: map[string]*User{
			"1": {
				ID:        "1",
				FirstName: "Vasu",
				LastName:  "Jain",
				Email:     "vasu@example.com",
				IsActive:  true,
			},
		},
	}

	m.HandleFunc("/health", HealthHandler)
	m.HandleFunc("/users", UserHandler(store))

	http.ListenAndServe(":8080", m)
}
