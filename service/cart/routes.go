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
	userStore    types.UserStore
}

func (handler *Handler) NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, productStore: productStore}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(handler.HandleCheckout, handler.userStore)).Methods(http.MethodPost)
}

func (handler *Handler) HandleCheckout(res http.ResponseWriter, req *http.Request) {
	userID := 0 //TODO: WILL GET FROM JWT

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

	productIds, err := GetCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}

	prodStore, err := handler.productStore.GetProductByIDs(productIds)
	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := handler.CreateOrder(prodStore, cart.Items, userID)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(res, http.StatusOK, map[string]any{
		"total_price": totalPrice,
		"order_id":    orderID,
	})

}
