package users

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Create(user User) (User, error)
	Get(id int) (User, error)
	Update(user User) (User, error)
	Delete(id int) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) Create(user User) (User, error) {
	// Generate and set the new user password.
	hash, err := generatePassword(user.Password)
	if err != nil {
		return User{}, err
	}
	user.Password = hash

	user, err = s.repository.Create(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s Service) Get(id int) (User, error) {
	user, err := s.repository.Get(id)
	if err != nil {
		if err == ErrRecordNotFound {
			return User{}, ErrRecordNotFound
		}
		return User{}, err
	}

	return user, nil
}

func (s Service) Update(id int, user User) (User, error) {
	// If needed, generate and set the new user password.
	if user.Password != "" {
		hash, err := generatePassword(user.Password)
		if err != nil {
			return User{}, err
		}
		user.Password = hash
	}

	user.ID = id

	user, err := s.repository.Update(user)
	if err != nil {
		if err == ErrRecordNotFound {
			return User{}, ErrRecordNotFound
		}
		return User{}, err
	}

	return user, nil
}

func (s Service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		if err == ErrRecordNotFound {
			return ErrRecordNotFound
		}
		return err
	}

	return nil
}

func generatePassword(plainPassword string) (string, error) {
	// Generate "hash" to mysql from user password.
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(hash), nil
}
