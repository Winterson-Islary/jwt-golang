package cart

import (
	"database/sql"

	"github.com/Winterson-Islary/jwt-golang.git/types"
)

type Store struct {
	db *sql.DB
}

func (store *Store) NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) CreateOrder(order types.Order) (int, error) {
	result, err := store.db.Exec("INSERT INTO orders (userID, total, status, address) VALUES ($1,$2,$3,$4)", order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (store *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := store.db.Exec("INSERT INTO order_items(orderId, productId, quantity, price) VALUES ($1,$2,$3,$4)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
