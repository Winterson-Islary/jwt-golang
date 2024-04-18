package cart

import (
	"net/http"

	"github.com/Winterson-Islary/jwt-golang.git/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.CartStore
}

func (handler *Handler) NewHandler(store types.CartStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", handler.HandleCheckout).Methods(http.MethodPost)
}

func (handler *Handler) HandleCheckout(res http.ResponseWriter, req *http.Request) {

}
