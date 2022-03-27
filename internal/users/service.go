package users

import (
	"github.com/marcosstupnicki/go-users/internal/config"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Service struct {
	repository Repository

}

func NewService(cfg config.Database) (Service, error) {
	repository, err := newRepository(cfg)
	if err != nil {
		return Service{}, err
	}

	return Service{
		repository: repository,
	}, nil
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

func generatePassword(plainPassword string) (string, error){
	// Generate "hash" to store from user password.
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(hash), nil
}