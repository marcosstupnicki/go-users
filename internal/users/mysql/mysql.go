package mysql

import (
	"errors"
	"github.com/marcosstupnicki/go-users/internal/platform/config"
	"github.com/marcosstupnicki/go-users/internal/platform/db"
	"github.com/marcosstupnicki/go-users/internal/users"
	"gorm.io/gorm"
)

var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = errors.New("record not found error")
)

type MySQL struct {
	DB *gorm.DB
}

func NewMySQL(cfg config.Database) (MySQL, error) {
	gdb, err := db.Connect(cfg)
	if err != nil {
		return MySQL{}, err
	}

	return MySQL{
		DB: gdb,
	}, nil
}

func (repository MySQL) Create(user users.User) (users.User, error) {
	tx := repository.DB.Create(&user)
	if tx.Error != nil {
		return users.User{}, tx.Error
	}

	return user, nil
}

func (repository MySQL) Get(id int) (users.User, error) {
	user := users.User{ID: id}
	tx := repository.DB.First(&user)

	if tx.RowsAffected == 0{
		return users.User{}, ErrRecordNotFound
	}
	if tx.Error != nil {
		return users.User{}, tx.Error
	}

	return user, nil
}

func (repository MySQL) Update(user users.User) (users.User, error) {
	tx := repository.DB.Model(&user).Updates(user)
	if tx.RowsAffected == 0 {
		return users.User{}, ErrRecordNotFound
	}
	if tx.Error != nil {
		return users.User{}, tx.Error
	}

	return user, nil
}

func (repository MySQL) Delete(id int) error {
	user := users.User{ID: id}

	tx := repository.DB.Delete(&user)
	if tx.RowsAffected == 0 {
		return ErrRecordNotFound
	}

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repository MySQL) AutoMigrate() error {
	err := repository.DB.AutoMigrate(&users.User{})
	if err != nil {
		return err
	}

	return nil
}