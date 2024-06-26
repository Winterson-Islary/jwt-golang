CREATE TABLE IF NOT EXISTS products (
	id INT CHECK (id > 0) NOT NULL GENERATED ALWAYS AS IDENTITY,
	firstName VARCHAR(255) NOT NULL,
	description TEXT NOT NULL,
	image VARCHAR(255) NOT NULL,
	price DECIMAL(10, 2) NOT NULL,
	quantity INT CHECK (quantity > 0) NOT NULL,
	createdAt TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,

	PRIMARY KEY (id)
);