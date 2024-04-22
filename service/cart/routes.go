package cart

import (
	"fmt"
	"net/http"

	"github.com/Winterson-Islary/jwt-golang.git/types"
	"github.com/Winterson-Islary/jwt-golang.git/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
}

func (handler *Handler) NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, productStore: productStore}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", handler.HandleCheckout).Methods(http.MethodPost)
}

func (handler *Handler) HandleCheckout(res http.ResponseWriter, req *http.Request) {
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(req, &cart); err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(res, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	//TODO: Implement
	prodStore, err := handler.productStore.GetProductByIDs(productIds)
}
