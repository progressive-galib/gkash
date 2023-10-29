package main

import (
	"math/rand"
)

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int
	Balance   float32
}

func NewAccount(firstName string, LastName string, num int) *Account {
	return &Account{
		ID:        rand.Intn(100000),
		FirstName: firstName,
		LastName:  LastName,
		Number:    num,
	}
}
