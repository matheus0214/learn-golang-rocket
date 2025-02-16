package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Biography string    `json:"biography"`
}

func NewUser(firstName, lastName, biography string) (User, error) {
	if firstName == "" {
		return User{}, requiredField("first name")
	}

	if lastName == "" {
		return User{}, requiredField("last name")
	}

	if biography == "" {
		return User{}, requiredField("biography")
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return User{}, errors.New("error to create user id")
	}

	return User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Biography: biography,
	}, nil
}

func requiredField(field string) error {
	return fmt.Errorf("the field %s is required", field)
}
