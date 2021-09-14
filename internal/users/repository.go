package users

import "gorm.io/gorm"

type Repository struct {
	*gorm.DB
}

func NewRepository(environment string){
	if environment == {

	}
}

func (r Repository) Create(user User) (User, error) {
	tx := r.DB.Create(&user)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func (s Repository) Get(id int) (User, error) {
	return User{
		Id: id,
		Email: "algunotro@email.com",
	}, nil
}

func (s Repository) Update(id int, user User) (User, error) {
	return User{
		Email: "algun@email.com",
	}, nil
}

func (s Repository) Delete(id int) error {
	return nil
}