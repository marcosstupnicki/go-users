package users

import (
	"errors"
	"gorm.io/gorm"
)

var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = errors.New("record not found error")
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		DB: db,
	}
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

func (repository Repository) Update(id int, user User) (User, error) {
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