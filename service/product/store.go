package product

import (
	"database/sql"
	"fmt"

	"github.com/Winterson-Islary/jwt-golang.git/types"
	"github.com/Winterson-Islary/jwt-golang.git/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) GetProducts() ([]types.Product, error) {
	rows, err := store.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		prod, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *prod)
	}
	return products, nil
}

func (store *Store) GetProductByIDs(PIDs []int) ([]types.Product, error) {
	placeholders := utils.GetIDPlaceholders(PIDs)
	query := fmt.Sprintf("SELECT * FROM product WHERE id IN (%s)", placeholders)

	// Converting PIDs (Product IDs) to []interface{}, ig to give each product some methods
	args := make([]interface{}, len(PIDs))
	for index, item := range PIDs {
		args[index] = item
	}

	rows, err := store.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		prod, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *prod)
	}

	return products, nil
}

func (store *Store) UpdateProduct(product types.Product) error {
	_, err := store.db.Exec("UPDATE products SET name = $1, price = $2, image = $3, description = $4, quantity = $5 WHERE id = $6", product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
