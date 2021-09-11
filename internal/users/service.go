package users

type UserService struct {

}

func (s UserService) Create(userRequest UserRequest) (User, error) {
	return User{
		Id: 12,
		Email: "algun@email.com",
	}, nil
}

func (s UserService) Get(id int) (User, error) {
	return User{
		Email: "algunotro@email.com",
	}, nil
}

func (s UserService) Update(id int, userRequest UserRequest) (User, error) {
	return User{
		Email: "algun@email.com",
	}, nil
}

func (s UserService) Delete(id int) error {
	return nil
}