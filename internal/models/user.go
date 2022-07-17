package models

//go:generate gotests -w -all .

type User struct {
	ID   int
	Name string
	Age  int
}
