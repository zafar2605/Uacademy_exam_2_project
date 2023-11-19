package postgres

import (
	"database/sql"
	"fmt"
	"strconv"

	"market/helpers"
	"market/model"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (p *ProductRepo) Create(req model.ProductCreate) (*model.Product, error) {
	resp, err := p.GetLast()
	if err != nil {
		return nil, err
	}
	last_number, _ := strconv.Atoi(resp)

	var product = model.Product{
		Id:          "P-" + helpers.GetSerialId(last_number),
		CategoryId:  req.CategoryId,
		Title:       req.Title,
		Photos:      req.Photos,
		Description: req.Description,
		Price:       req.Price,
	}

	var query = `INSERT INTO products(id, category_id, title, photos, description, price, updated_at) VALUES($1, $2, $3, $4, $5, $6, NOW())`

	_, err = p.db.Exec(query, product.Id, product.CategoryId, product.Title, product.Photos, product.Description, product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (c *ProductRepo) GetById(req model.ProductPrimaryKey) (*model.Product, error) {

	var product = model.Product{}
	var query = `
		SELECT 
			p.id, p.category_id, p.title, p.photos, p.description, p.price, p.created_at, p.updated_at,
			c.id, c.title, c.image, c.parent_id, c.created_at, c.updated_at	
		FROM 
			products AS p
		JOIN category AS c ON c.id = p.category_id
		WHERE p.id = $1
	`
	resp := c.db.QueryRow(query, req.Id)
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	err := resp.Scan(
		&product.Id,
		&product.CategoryId,
		&product.Title,
		&product.Photos,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,

		&product.Category.Id,
		&product.Category.Title,
		&product.Category.Image,
		&product.Category.ParentID,
		&product.Category.CreatedAt,
		&product.Category.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (c *ProductRepo) GetList(req model.GetListProductRequest) (*model.GetListProductResponse, error) {
	var resp = model.GetListProductResponse{}
	var offset string = " offset"
	var limit string = " limit"
	if req.Offset <= 0 {
		offset += " 0"
	} else if req.Offset > 0 {
		offset += fmt.Sprintf(" %d", req.Offset)
	}

	if req.Limit <= 0 {
		limit += " 10"
	} else if req.Limit > 0 {
		limit += fmt.Sprintf(" %d", req.Limit)
	}
	var query = `
		SELECT 
			COUNT(*) OVER(), p.id, p.category_id, p.title, p.photos, p.description, p.price, p.created_at, p.updated_at,
			c.id, c.title, c.image, c.parent_id, c.created_at, c.updated_at	
		FROM 
			products AS p
		JOIN category AS c ON c.id = p.category_id
	`
	query += offset + limit
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product = model.Product{}

		rows.Scan(
			&resp.Count,
			&product.Id,
			&product.CategoryId,
			&product.Title,
			&product.Photos,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,

			&product.Category.Id,
			&product.Category.Title,
			&product.Category.Image,
			&product.Category.ParentID,
			&product.Category.CreatedAt,
			&product.Category.UpdatedAt,
		)

		resp.Products = append(resp.Products, product)
	}

	rows.Close()
	return &resp, nil
}

func (c *ProductRepo) Update(req model.ProductUpdate) (string, error) {

	_, err := c.db.Exec(`
		UPDATE
			products
		SET
			category_id = $1, title = $2, photos = $3, description = $4, price = $5, updated_at = NOW()
	 	WHERE id = $6`, helpers.NewNullString(req.CategoryId), req.Title, req.Photos, req.Description, req.Price, req.Id)

	if err != nil {
		return "Can not update", err
	}

	return "Updated", nil
}

func (c *ProductRepo) Delete(req model.ProductPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM products WHERE id = $1`, req.Id)

	if err != nil {
		return "Can not delete", err
	}

	return "Deleted", nil
}

func (c *ProductRepo) GetLast() (string, error) {
	var count string
	resp := c.db.QueryRow(`SELECT id FROM products ORDER BY id DESC LIMIT 1`)

	if resp.Err() != nil {
		return "P-00000000", resp.Err()
	}
	err := resp.Scan(
		&count,
	)
	if err != nil {
		return "Can not scan from db", nil
	}

	return count[2:], nil
}
