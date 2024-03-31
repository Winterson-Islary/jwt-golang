package user

import (
	"net/http"

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

}
