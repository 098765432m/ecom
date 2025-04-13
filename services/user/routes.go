package user

import (
	"fmt"
	"net/http"

	"github.com/098765432m/ecom/services/auth"
	"github.com/098765432m/ecom/types"
	"github.com/098765432m/ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin)
	router.HandleFunc("/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", validationErrors))
		// utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload &v", validationErrors))
		return 
	}

	// check user exist
	_, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	hashedPassword, err := auth.HashedPassword(payload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FistName,
		LastName:  payload.LastName,
		Password:  hashedPassword,
		Email:     payload.Email,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)

		return
	}

	utils.WtiteJSON(w, http.StatusCreated, nil)
}
