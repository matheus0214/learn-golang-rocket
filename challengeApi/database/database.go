package database

import (
	"challengeApi/domain"
	"errors"
)

type ID string

type DB interface {
	FindAll() map[ID]domain.User
	FindById(id string) (domain.User, error)
	Insert(domain.User) error
	Update(domain.User) error
	Delete(id string) error
}

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyRegister = errors.New("user already register")

type Database struct {
	data map[ID]domain.User
}

func (d Database) FindAll() map[ID]domain.User {
	return d.data
}

func (d Database) FindById(id string) (domain.User, error) {
	user, ok := d.data[ID(id)]

	if !ok {
		return domain.User{}, ErrUserNotFound
	}

	return user, nil
}

func (d Database) Insert(user domain.User) error {
	if _, ok := d.data[ID(user.ID.String())]; ok {
		return ErrUserAlreadyRegister
	}

	d.data[ID(user.ID.String())] = user

	return nil
}

func (d Database) Update(user domain.User) error {
	if _, ok := d.data[ID(user.ID.String())]; !ok {
		return ErrUserNotFound
	}

	d.data[ID(user.ID.String())] = user

	return nil
}

func (d Database) Delete(id string) error {
	if _, ok := d.data[ID(id)]; !ok {
		return ErrUserNotFound
	}

	delete(d.data, ID(id))

	return nil
}

func NewDatabase() DB {
	return Database{data: make(map[ID]domain.User)}
}
