package models

type Account struct {
	AccNo        int
	Name         string
	Email        string
	Password     string
	Balance      float32
	Transactions []string
}
