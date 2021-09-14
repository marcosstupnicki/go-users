package users

type Service struct {
	repository Repository
}

func (s Service) Create(userRequest UserRequest) (User, error) {
	user := buildUserFromUserRequest(userRequest)

	user, err := s.repository.Create(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s Service) Get(id int) (User, error) {
	return User{
		Id:    id,
		Email: "algunotro@email.com",
	}, nil
}

func (s Service) Update(id int, userRequest UserRequest) (User, error) {
	return User{
		Email: "algun@email.com",
	}, nil
}

func (s Service) Delete(id int) error {
	return nil
}

func buildUserFromUserRequest(user UserRequest) User {
	return User {
		Email:    user.Email,
		Password: user.Password,
	}
}
