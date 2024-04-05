package user

import (
	"fmt"
	"net/http"

	"github.com/Winterson-Islary/jwt-golang.git/types"
	"github.com/Winterson-Islary/jwt-golang.git/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", handler.HandleLogin).Methods("POST")
	router.HandleFunc("/register", handler.HandleRegister).Methods("POST")
}

func (handler *Handler) HandleLogin(res http.ResponseWriter, req *http.Request) {

}

func (handler *Handler) HandleRegister(res http.ResponseWriter, req *http.Request) {

	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(req, payload); err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
	}

	_, err := handler.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	err = handler.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	})
	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(res, http.StatusCreated, nil)
}
