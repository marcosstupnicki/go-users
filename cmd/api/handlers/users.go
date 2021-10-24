package handlers

import (
	"encoding/json"
	"github.com/marcosstupnicki/go-users/internal/users"
	gowebapp "github.com/marcosstupnicki/go-webapp/pkg"
	"net/http"
	"strconv"
)

const (
	_ErrorMessageInvalidIDParam      = "invalid param ID. ID must be a integer."
	_ErrorMessageCouldNotDecodeInput = "could not decode value from input"
)

type Service interface {
	Create(user users.UserRequest) (users.User, error)
	Get(id int) (users.User, error)
	Update(id int, user users.UserRequest) (users.User, error)
	Delete(id int) error
}

type UserHandler struct {
	Service Service
}

func NewHandler(service Service) UserHandler {
	return UserHandler{
		Service: service,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest users.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		gowebapp.RespondWithError(w, http.StatusBadRequest, _ErrorMessageCouldNotDecodeInput)
		return
	}

	user, err := h.Service.Create(userRequest)
	if err != nil {
		gowebapp.RespondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	userResponse := buildUserResponseFromUser(user)
	gowebapp.RespondWithJSON(w, http.StatusCreated, userResponse)
	return
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	idParam := gowebapp.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		gowebapp.RespondWithError(w, http.StatusBadRequest, _ErrorMessageInvalidIDParam)
		return
	}

	user, err := h.Service.Get(id)
	if err != nil {
		if err == users.ErrRecordNotFound {
			gowebapp.RespondWithError(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		gowebapp.RespondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	userResponse := buildUserResponseFromUser(user)

	gowebapp.RespondWithJSON(w, http.StatusCreated, userResponse)
	return
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := gowebapp.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		gowebapp.RespondWithError(w, http.StatusBadRequest, _ErrorMessageInvalidIDParam)
		return
	}

	var userRequest users.UserRequest
	err = json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		gowebapp.RespondWithError(w, 400, _ErrorMessageCouldNotDecodeInput)
		return
	}

	user, err := h.Service.Update(id, userRequest)
	if err != nil {
		if err == users.ErrRecordNotFound {
			gowebapp.RespondWithError(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		gowebapp.RespondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	userResponse := buildUserResponseFromUser(user)

	gowebapp.RespondWithJSON(w, http.StatusOK, userResponse)
	return
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := gowebapp.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	err = h.Service.Delete(id)
	if err != nil {
		if err == users.ErrRecordNotFound {
			gowebapp.RespondWithError(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		gowebapp.RespondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	gowebapp.RespondWithJSON(w, http.StatusNoContent, nil)
	return
}

func buildUserResponseFromUser(user users.User) users.UserResponse {
	return users.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}