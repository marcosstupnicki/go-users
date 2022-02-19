package users

import (
	"errors"
	"github.com/marcosstupnicki/go-users/internal/config"
	"github.com/marcosstupnicki/go-users/internal/db"
	"gorm.io/gorm"
)

var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = errors.New("record not found error")
)

type Repository struct {
	DB *gorm.DB
}

func newRepository(cfg config.Database) (Repository, error) {
	gdb, err := db.Connect(cfg)
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		DB: gdb,
	}, nil
}

func (repository Repository) Create(user User) (User, error) {
	tx := repository.DB.Create(&user)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func (repository Repository) Get(id int) (User, error) {
	user := User{ID: id}
	tx := repository.DB.First(&user)

	if tx.RowsAffected == 0{
		return User{}, ErrRecordNotFound
	}
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func (repository Repository) Update(user User) (User, error) {
	tx := repository.DB.Model(&user).Updates(user)
	if tx.RowsAffected == 0 {
		return User{}, ErrRecordNotFound
	}
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func (repository Repository) Delete(id int) error {
	user := User{ID: id}

	tx := repository.DB.Delete(&user)
	if tx.RowsAffected == 0 {
		return ErrRecordNotFound
	}

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}