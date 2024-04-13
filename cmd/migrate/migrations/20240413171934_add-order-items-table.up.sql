CREATE TABLE IF NOT EXISTS order_items (
	id INT CHECK (id > 0) NOT NULL GENERATED ALWAYS AS IDENTITY,
	orderId INT CHECK (orderId > 0) NOT NULL,
	productId INT CHECK (productId > 0) NOT NULL,
	quantity INT NOT NULL,
	price DECIMAL(10, 2) NOT NULL,

	PRIMARY KEY (id),
	FOREIGN KEY (orderId) REFERENCES orders(id),
	FOREIGN KEY (productId) REFERENCES products(id)
);