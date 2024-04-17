package product

import (
	"net/http"

	"github.com/Winterson-Islary/jwt-golang.git/types"
	"github.com/Winterson-Islary/jwt-golang.git/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (handler Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", handler.handleCreateProduct).Methods(http.MethodGet)

}

func (handler *Handler) handleCreateProduct(res http.ResponseWriter, req *http.Request) {
	pStore, err := handler.store.GetProducts()
	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(res, http.StatusOK, pStore)
}
