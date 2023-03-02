package main

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    uuid.UUID `json:"number"`
	Balance   int64     `json:"balnce"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    uuid.New(),
		CreatedAt: time.Now().UTC(),
	}
}
