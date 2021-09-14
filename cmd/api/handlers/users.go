package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/marcosstupnicki/go-users/internal/users"
	"net/http"
	"strconv"
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
		Service: users.Service{},
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {

	var userRequest users.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)

	user, err := h.Service.Create(userRequest)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	userResponse := buildUserResponseFromUser(user)

	render.JSON(w, r, userResponse)
}

func(h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r,"id")
	if idParam != "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	user, err := h.Service.Get(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	userResponse := buildUserResponseFromUser(user)

	render.JSON(w, r, userResponse)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	user, err := h.Service.Update(12, users.UserRequest{})
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	userResponse := buildUserResponseFromUser(user)

	render.JSON(w, r, userResponse)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Delete(12)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func buildUserResponseFromUser(user users.User) users.UserResponse{
	return users.UserResponse{
		ID: user.Id,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}