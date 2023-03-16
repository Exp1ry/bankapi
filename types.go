package main

import (
	"time"
)

type CreateAccountRequest struct {
	Type                string `json:"type"`
	Name                string `json:"name"`
	Location            string `json:"location"`
	City                string `json:"city"`
	Phone               string `json:"phone"`
	ProductsAndServices string `json:"product_and_service"`
}

// type Account struct {
// 	ID        int       `json:"id"`
// 	FirstName string    `json:"first_name"`
// 	LastName  string    `json:"last_name"`
// 	Number    uuid.UUID `json:"number"`
// 	Balance   int64     `json:"balnce"`
// 	CreatedAt time.Time `json:"created_at"`
// }

type FurnitureStore struct {
	ID                  int       `json:"id"`
	Type                string    `json:"type"`
	Name                string    `json:"name"`
	Location            string    `json:"location"`
	City                string    `json:"city"`
	Phone               string    `json:"phone"`
	ProductsAndServices string    `json:"product_and_service"`
	CreatedAt           time.Time `json:"created_at"`
}

func NewAccount(name, typee, location, city, phone, productAndService string) *FurnitureStore {
	return &FurnitureStore{
		Type:                typee,
		Name:                name,
		Location:            location,
		City:                city,
		Phone:               phone,
		ProductsAndServices: productAndService,
		CreatedAt:           time.Now().UTC(),
	}
}
