package repositories

//go:generate gotests -w -all .

import "sample/internal/models"

type UserRepository struct {
}

func (r UserRepository) Find(id int) models.User {
	return models.User{ID: id}
}
