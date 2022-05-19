package users

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/marcosstupnicki/go-users/internal/platform/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// ErrUserNotFound users not found error
	ErrUserNotFound = errors.New("user not found")
)

type MySQL struct {
	DB *gorm.DB
}

func NewMySQL(cfg config.Database) (MySQL, error) {
	gdb, err := connect(cfg)
	if err != nil {
		return MySQL{}, err
	}

	return MySQL{
		DB: gdb,
	}, nil
}

func (repository MySQL) Create(user User) (User, error) {
	tx := repository.DB.Create(&user)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func (repository MySQL) Get(id int) (User, error) {
	user := User{ID: id}
	tx := repository.DB.First(&user)

	if tx.RowsAffected == 0 {
		return User{}, ErrUserNotFound
	}
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func (repository MySQL) Update(user User) (User, error) {
	tx := repository.DB.Model(&user).Updates(user)
	if tx.RowsAffected == 0 {
		return User{}, ErrUserNotFound
	}
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func (repository MySQL) Delete(id int) error {
	user := User{ID: id}

	tx := repository.DB.Delete(&user)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return ErrUserNotFound
	}


	return nil
}

func (repository MySQL) AutoMigrate() error {
	err := repository.DB.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	return nil
}

func connect(cfg config.Database) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: cfg.LogLevel,
		},
	)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
