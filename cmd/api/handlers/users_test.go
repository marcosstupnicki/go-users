package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/marcosstupnicki/go-users/internal/users"
	gowebapp "github.com/marcosstupnicki/go-webapp/pkg"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) Create(_ users.User) (users.User, error) {
	args := s.Called()
	return args.Get(0).(users.User), args.Error(1)
}

func (s *ServiceMock) Get(_ int) (users.User, error) {
	args := s.Called()
	return args.Get(0).(users.User), args.Error(1)
}

func (s *ServiceMock) Update(_ int, _ users.User) (users.User, error) {
	args := s.Called()
	return args.Get(0).(users.User), args.Error(1)
}

func (s *ServiceMock) Delete(_ int) error {
	args := s.Called()
	return args.Error(0)
}

var ErrInternalErr = errors.New("internal error")

func TestUserHandler_Create(t *testing.T) {
	user := users.User{
		Email:    "dummy@email.com",
		Password: "dummypassword",
	}

	requestOk, err := json.Marshal(user)
	require.NoError(t, err)

	requestInvalid := []byte("request_invalid")

	user.ID = 5

	var tests = []struct {
		name               string
		service            *ServiceMock
		request            *bytes.Reader
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			name: "Ok - Create user success",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Create", mock.Anything).Return(user, nil)
				return &m
			}(),
			request:            bytes.NewReader(requestOk),
			expectedResponse:   "{\"id\":5,\"email\":\"dummy@email.com\"}",
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "Fail - Bad request",
			request:            bytes.NewReader(requestInvalid),
			expectedResponse:   "{\"message\":\"could not decode value from input\"}",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail - Internal error in user service",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Create", mock.Anything).Return(users.User{}, ErrInternalErr)
				return &m
			}(),
			request:            bytes.NewReader(requestOk),
			expectedResponse:   "{\"message\":\"Internal Server Error\"}",
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := gowebapp.NewWebApp("local")
			handler := NewHandler(tt.service)
			app.Post("/users", handler.Create)

			r := httptest.NewRequest(http.MethodPost, "/users", tt.request)

			rr := httptest.NewRecorder()
			app.Router.ServeHTTP(rr, r)

			res := rr.Result()
			resBody, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			require.Equal(t, tt.expectedStatusCode, res.StatusCode)
			require.Equal(t, tt.expectedResponse, string(resBody))
		})
	}
}

func TestUserHandler_Get(t *testing.T) {
	user := users.User{
		ID:       5,
		Email:    "dummy@email.com",
		Password: "dummypassword",
	}

	var tests = []struct {
		name               string
		service            *ServiceMock
		id                 int
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			name: "Ok - Get user success",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Get", mock.Anything).Return(user, nil)
				return &m
			}(),
			id:                 5,
			expectedResponse:   "{\"id\":5,\"email\":\"dummy@email.com\"}",
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Fail - User not found",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Get", mock.Anything).Return(users.User{}, users.ErrUserNotFound)
				return &m
			}(),
			id:                 6,
			expectedResponse:   "{\"message\":\"user not found\"}",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "Fail - Internal error in user service",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Get", mock.Anything).Return(users.User{}, ErrInternalErr)
				return &m
			}(),
			id:                 7,
			expectedResponse:   "{\"message\":\"Internal Server Error\"}",
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := gowebapp.NewWebApp("local")
			handler := NewHandler(tt.service)
			app.Get("/users/{id}", handler.Get)

			r := httptest.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(tt.id), nil)

			rr := httptest.NewRecorder()
			app.Router.ServeHTTP(rr, r)

			res := rr.Result()
			resBody, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			require.Equal(t, tt.expectedStatusCode, res.StatusCode)
			require.Equal(t, tt.expectedResponse, string(resBody))
		})
	}
}

func TestUserHandler_Update(t *testing.T) {
	user := users.User{
		Email:    "dummy@email.com",
		Password: "dummypassword",
	}

	requestOk, err := json.Marshal(user)
	require.NoError(t, err)

	requestInvalid := []byte("request_invalid")

	user.ID = 5

	var tests = []struct {
		name               string
		service            *ServiceMock
		id                 int
		request            *bytes.Reader
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			name: "Ok - Update user success",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Update", mock.Anything).Return(user, nil)
				return &m
			}(),
			id:                 5,
			request:            bytes.NewReader(requestOk),
			expectedResponse:   "{\"id\":5,\"email\":\"dummy@email.com\"}",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail - Bad request",
			id:                 5,
			request:            bytes.NewReader(requestInvalid),
			expectedResponse:   "{\"message\":\"could not decode value from input\"}",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail - User not found",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Update", mock.Anything).Return(users.User{}, users.ErrUserNotFound)
				return &m
			}(),
			id:                 6,
			request:            bytes.NewReader(requestOk),
			expectedResponse:   "{\"message\":\"user not found\"}",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "Fail - Internal error in user service",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Update", mock.Anything).Return(users.User{}, ErrInternalErr)
				return &m
			}(),
			id:                 6,
			request:            bytes.NewReader(requestOk),
			expectedResponse:   "{\"message\":\"Internal Server Error\"}",
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := gowebapp.NewWebApp("local")
			handler := NewHandler(tt.service)
			app.Put("/users/{id}", handler.Update)

			r := httptest.NewRequest(http.MethodPut, "/users/"+strconv.Itoa(tt.id), tt.request)

			rr := httptest.NewRecorder()
			app.Router.ServeHTTP(rr, r)

			res := rr.Result()
			resBody, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			require.Equal(t, tt.expectedStatusCode, res.StatusCode)
			require.Equal(t, tt.expectedResponse, string(resBody))
		})
	}
}

func TestUserHandler_Delete(t *testing.T) {
	var tests = []struct {
		name               string
		service            *ServiceMock
		id                 int
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			name: "Ok - Delete user success",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Delete", mock.Anything).Return(nil)
				return &m
			}(),
			id:                 5,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "Fail - User not found",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Delete", mock.Anything).Return(users.ErrUserNotFound)
				return &m
			}(),
			id:                 6,
			expectedResponse:   "{\"message\":\"user not found\"}",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "Fail - Internal error in user service",
			service: func() *ServiceMock {
				m := ServiceMock{}
				m.On("Delete", mock.Anything).Return(ErrInternalErr)
				return &m
			}(),
			id:                 7,
			expectedResponse:   "{\"message\":\"Internal Server Error\"}",
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := gowebapp.NewWebApp("local")
			handler := NewHandler(tt.service)
			app.Delete("/users/{id}", handler.Delete)

			r := httptest.NewRequest(http.MethodDelete, "/users/"+strconv.Itoa(tt.id), nil)

			rr := httptest.NewRecorder()
			app.Router.ServeHTTP(rr, r)

			res := rr.Result()
			resBody, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			require.Equal(t, tt.expectedStatusCode, res.StatusCode)
			require.Equal(t, tt.expectedResponse, string(resBody))
		})
	}
}
