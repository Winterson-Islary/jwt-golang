package user

import (
	"net/http"

	"github.com/Winterson-Islary/jwt-golang.git/types"
	"github.com/Winterson-Islary/jwt-golang.git/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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

}
