package user

import (
	"fmt"
	"net/http"

	"github.com/Winterson-Islary/jwt-golang.git/service/auth"
	"github.com/Winterson-Islary/jwt-golang.git/types"
	"github.com/Winterson-Islary/jwt-golang.git/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// * Route Registration
func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", handler.HandleLogin).Methods("POST")
	router.HandleFunc("/register", handler.HandleRegister).Methods("POST")
}

// * Handling User Login
func (handler *Handler) HandleLogin(res http.ResponseWriter, req *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(res, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	user, err := handler.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}
	if !auth.ComparePasswords(user.Password, []byte(payload.Password)) {
		utils.WriteError(res, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
	}

	utils.WriteJSON(res, http.StatusOK, map[string]string{"token": ""})
}

// * Handling User Registration
func (handler *Handler) HandleRegister(res http.ResponseWriter, req *http.Request) {

	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(res, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	_, err := handler.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(res, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}
	err = handler.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(res, http.StatusCreated, nil)
}
