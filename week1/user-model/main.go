package main

import (
	"errors"
	"fmt"
	"goLabs/week1/user-model/utils"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func NewUser(firstname, lastname, email string) (*User, error) {
	if firstname == "" {
		return nil, errors.New("First name can't be empty")
	}
	if lastname == "" {
		return nil, errors.New("Last name can't be empty")
	}
	if email == "" {
		return nil, errors.New("Email can't be empty")
	}

	ValidEmail := utils.IsValidEmail(email)

	if !ValidEmail {
		return nil, errors.New("Email is not valid")
	}

	return &User{
		ID:        utils.GenerateUUID(),
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		IsActive:  true,
	}, nil
}

func (u *User) Deactivate() {
	u.IsActive = false
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	newUser, err := NewUser("Vasu", "Jain", "vasujain275@gmail.com")

	if err != nil {
		fmt.Printf("Error : %v \n", err)
		return
	}

	fmt.Printf("User Created! Details - \n %v \n", newUser)

	fmt.Println("Full Name:", newUser.FullName())

	newUser.Deactivate()
	fmt.Println("User active status after deactivation:", newUser.IsActive)
}
