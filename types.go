package main

import "math/rand/v2"

type Account struct {
	ID        int `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64 `json:"number"`
	Balance   int64`json:"balance"`
}

// Snake case for the json serialization of the Account properties is also possible.
// ex: `json:"first_name"`

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID: rand.IntN(10000),
		FirstName: firstName,
		LastName: lastName,
		Number: int64(rand.IntN(1000000)),
	}
}