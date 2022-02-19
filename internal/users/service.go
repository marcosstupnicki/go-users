package users

import "github.com/marcosstupnicki/go-users/internal/config"

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

func (s Service) Create(userRequest UserRequest) (User, error) {
	user := buildUserFromUserRequest(userRequest)

	user, err := s.repository.Create(user)
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

func (s Service) Update(id int, userRequest UserRequest) (User, error) {
	user := buildUserFromUserRequest(userRequest)
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

func buildUserFromUserRequest(user UserRequest) User {
	return User {
		Email:    user.Email,
		Password: user.Password,
	}
}
