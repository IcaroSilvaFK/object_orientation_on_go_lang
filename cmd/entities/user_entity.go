package entities

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	name     string
	age      int8
	email    string
	password string
}

type UserEntityInterface interface {
	GetName() string
	GetAge() int8
	GetEmail() string
	GetPassword() string
	ensurePassword() error
	verifyEmail(string) error
	Save() (bool, error)
}

var users []UserEntity

func NewUserEntity(name, email, password string, age int8) UserEntityInterface {

	return &UserEntity{
		name, age, email, password,
	}

}

func (u *UserEntity) Save() (bool, error) {

	if err := u.verifyEmail(u.email); err != nil {
		return false, err
	}
	if err := u.ensurePassword(); err != nil {
		return false, err
	}

	users = append(users, *u)

	return false, nil
}

func (u *UserEntity) GetName() string {
	return u.name
}
func (u *UserEntity) GetAge() int8 {
	return u.age
}
func (u *UserEntity) GetEmail() string {
	return u.email
}
func (u *UserEntity) GetPassword() string {
	return u.password
}
func (u *UserEntity) ensurePassword() error {

	bt, err := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.password = string(bt)

	return nil
}
func (u *UserEntity) verifyEmail(e string) error {

	for _, v := range users {
		if e == v.email {
			return errors.New("this email already exists")
		}
	}
	return nil
}
