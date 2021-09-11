package handlers

import (
	"fmt"
	"github.com/go-chi/render"
	"github.com/marcosstupnicki/go-users/internal/users"
	"net/http"
)

type UserService interface {
	Create(user users.UserRequest) (users.User, error)
	Get(id int) (users.User, error)
	Update(id int, user users.UserRequest) (users.User, error)
	Delete(id int) error
}

type UserHandler struct {
	Service UserService
}

func NewUserHandler() UserHandler {
	return UserHandler{
		Service: users.UserService{},
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user, err := h.Service.Create(users.UserRequest{})
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	userResponse := buildUserResponseFromUser(user)

	render.JSON(w, r, userResponse)
}

func(h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, GET")
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, UPDATE")
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, DELETE")
}

func buildUserResponseFromUser(user users.User) users.UserResponse{
	return users.UserResponse{
		ID: user.Id,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}