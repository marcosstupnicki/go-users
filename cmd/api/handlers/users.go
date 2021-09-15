package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/marcosstupnicki/go-users/internal/users"
	"net/http"
	"strconv"
)

const (
	_ErrorMessageInvalidIDParam      = "invalid param ID. ID must be an integer."
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
		respondWithError(w, 400, _ErrorMessageCouldNotDecodeInput)
		return
	}

	user, err := h.Service.Create(userRequest)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	userResponse := buildUserResponseFromUser(user)
	respondwithJSON(w, 201, userResponse)
	return
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, _ErrorMessageInvalidIDParam)
		return
	}

	user, err := h.Service.Get(id)
	if err != nil {
		if err == users.ErrRecordNotFound {
			respondWithError(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	userResponse := buildUserResponseFromUser(user)

	respondwithJSON(w, 200, userResponse)
	return
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, _ErrorMessageInvalidIDParam)
		return
	}

	var userRequest users.UserRequest
	err = json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		respondWithError(w, 400, _ErrorMessageCouldNotDecodeInput)
		return
	}

	user, err := h.Service.Update(id, userRequest)
	if err != nil {
		if err == users.ErrRecordNotFound {
			respondWithError(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	userResponse := buildUserResponseFromUser(user)

	respondwithJSON(w, 200, userResponse)
	return
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	err = h.Service.Delete(id)
	if err != nil {
		if err == users.ErrRecordNotFound {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	respondwithJSON(w, 200, nil)
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

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
