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
