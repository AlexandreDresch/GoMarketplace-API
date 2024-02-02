package database

import (
	"database/sql"

	"github.com/AlexandreDresch/GoMarketplace-API/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{
		db: db,
	}
}

func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query(`SELECT id, name, description, price, category_id, image_url FROM products`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := pd.db.QueryRow(`SELECT id, name, price, category_id, image_url, description FROM products WHERE id = $1`, id).
		Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL, &product.Description)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query(`SELECT id, name, price, category_id, image_url, description FROM products WHERE category_id = $1`, categoryID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (pd *ProductDB) CreateProduct(product *entity.Product) (string, error) {
	_, err := pd.db.Exec(`INSERT INTO products (id, name, description, price, category_id, image_url) VALUES ($1,?$2,$3,$4,$5,$6)`, product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		return "", err
	}

	return product.ID, nil
}
