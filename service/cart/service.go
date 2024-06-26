package cart

import (
	"fmt"

	"github.com/Winterson-Islary/jwt-golang.git/types"
)

func GetCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))
	for index, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}
		productIds[index] = item.ProductID
	}

	return productIds, nil
}

func (handler *Handler) CreateOrder(products []types.Product, items []types.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}
	if err := CheckItemInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	totalPrice := CalculateTotalPrice(items, productMap)

	//! NAIVE IMPLEMENTATION: Updating Stock (CAN CAUSE DATA INCONSISTENCIES)
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity
		handler.productStore.UpdateProduct(product)
	}
	//? Marked For Refactoring
	orderID, err := handler.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "HardCoded Address: Improvements Required",
	})
	if err != nil {
		return 0, 0, err
	}

	for _, item := range items {
		handler.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}

func CheckItemInStock(items []types.CartItem, products map[int]types.Product) error {

	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range items {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}
		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}
	return nil
}

func CalculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var total float64
	for _, item := range cartItems {
		product := products[item.ProductID]
		total += float64(item.Quantity) * product.Price
	}
	return total
}
