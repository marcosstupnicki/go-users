package users

import (
	"errors"
	"github.com/marcosstupnicki/go-users/internal/users/mysql"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type RepositoryMock struct {
	mock.Mock
}

func (s *RepositoryMock) Create(user User) (User, error) {
	args := s.Called()
	return args.Get(0).(User), args.Error(1)
}

func (s *RepositoryMock) Get(id int) (User, error) {
	args := s.Called()
	return args.Get(0).(User), args.Error(1)
}

func (s *RepositoryMock) Update(user User) (User, error) {
	args := s.Called()
	return args.Get(0).(User), args.Error(1)
}

func (s *RepositoryMock) Delete(id int) error {
	args := s.Called()
	return args.Error(0)
}

func TestService_Create(t *testing.T) {
	user := User{
		ID:        1,
		Email:     "some@email.com",
		Password:  "$2a$10$i8u5FgiJXRui/p.ZDXnDO.kVq3H6rbqrQp6rInFX.IeEO0zN/2F5G",
		CreatedAt: 1651422724,
		UpdatedAt: 1651422724,
	}

	var tests = []struct {
		name           string
		repo           *RepositoryMock
		user           User
		expectedResult User
		expectedError  error
	}{
		{
			name: "Ok",
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Create", mock.Anything).Return(user, nil)
				return &m
			}(),
			user: User{
				Email:    "some@email.com",
				Password: "some-password",
			},
			expectedResult: user,
		},
		{
			name: "Fail - Internal error",
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Create", mock.Anything).Return(User{}, errors.New("internal error"))
				return &m
			}(),
			user: User{
				Email:    "some@email.com",
				Password: "some-password",
			},
			expectedError: errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(tt.repo)
			result, err := service.Create(tt.user)
			require.Equal(t, tt.expectedError, err)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestService_Get(t *testing.T) {
	user := User{
		ID:        1,
		Email:     "some@email.com",
		Password:  "$2a$10$i8u5FgiJXRui/p.ZDXnDO.kVq3H6rbqrQp6rInFX.IeEO0zN/2F5G",
		CreatedAt: 1651422724,
		UpdatedAt: 1651422724,
	}

	var tests = []struct {
		name           string
		repo           *RepositoryMock
		id             int
		expectedResult User
		expectedError  error
	}{
		{
			name: "Ok",
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Get", mock.Anything).Return(user, nil)
				return &m
			}(),
			id:             1,
			expectedResult: user,
		},
		{
			name: "Fail - User not found",
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Get", mock.Anything).Return(User{}, mysql.ErrRecordNotFound)
				return &m
			}(),
			expectedError: mysql.ErrRecordNotFound,
		},
		{
			name: "Fail - Internal error",
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Get", mock.Anything).Return(User{}, errors.New("internal error"))
				return &m
			}(),
			expectedError: errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(tt.repo)
			result, err := service.Get(tt.id)
			require.Equal(t, tt.expectedError, err)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestService_Update(t *testing.T) {
	user := User{
		ID:        1,
		Email:     "some@email.com",
		Password:  "$2a$10$lG1aALcjSRwQ8zAKZmcxBOX3fnZ5dMPN9zTy58crosLdtZ8XQooBC",
		CreatedAt: 1651422724,
		UpdatedAt: 1651422724,
	}

	var tests = []struct {
		name           string
		repo           *RepositoryMock
		id             int
		user           User
		expectedResult User
		expectedError  error
	}{
		{
			name: "Ok",
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Update", mock.Anything).Return(user, nil)
				return &m
			}(),
			id: 1,
			user: User{
				Email:    "some2@email.com",
				Password: "some2-password",
			},
			expectedResult: user,
		},
		{
			name: "Fail - User not found",
			id:   1,
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Update", mock.Anything).Return(User{}, mysql.ErrRecordNotFound)
				return &m
			}(),
			expectedError: mysql.ErrRecordNotFound,
		},
		{
			name: "Fail - Internal error",
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Update", mock.Anything).Return(User{}, errors.New("internal error"))
				return &m
			}(),
			id: 1,
			user: User{
				Email:    "some@email.com",
				Password: "some-password",
			},
			expectedError: errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(tt.repo)
			result, err := service.Update(tt.id, tt.user)
			require.Equal(t, tt.expectedError, err)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestService_Delete(t *testing.T) {
	var tests = []struct {
		name          string
		repo          *RepositoryMock
		id            int
		expectedError error
	}{
		{
			name: "Ok",
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Delete", mock.Anything).Return(nil)
				return &m
			}(),
			id: 1,
		},
		{
			name: "Fail - User not found",
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Delete", mock.Anything).Return(mysql.ErrRecordNotFound)
				return &m
			}(),
			expectedError: mysql.ErrRecordNotFound,
		},
		{
			name: "Fail - Internal error",
			repo: func() *RepositoryMock {
				m := RepositoryMock{}
				m.On("Delete", mock.Anything).Return(errors.New("internal error"))
				return &m
			}(),
			expectedError: errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(tt.repo)
			err := service.Delete(tt.id)
			require.Equal(t, tt.expectedError, err)
		})
	}
}
