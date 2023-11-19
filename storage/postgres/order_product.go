package postgres

import (
	"database/sql"

	"market/model"

	"github.com/google/uuid"
)

type OrderProductRepo struct {
	db *sql.DB
}

func NewOrderProductRepo(db *sql.DB) *OrderProductRepo {
	return &OrderProductRepo{
		db: db,
	}
}

func (p *OrderProductRepo) Create(req model.AddOrderProduct) (*model.OrderProducts, error) {

	var orderProduct = model.OrderProducts{
		Id:           uuid.New().String(),
		OrderId:      req.OrderId,
		ProductId:    req.ProductId,
		DiscountType: req.DiscountType,
		Quantity:     req.Quantity,
	}

	var query = `INSERT INTO order_products(id, order_id, product_id, discount_type, quantity, updated_at) VALUES($1, $2, $3, $4, $5, NOW())`

	_, err := p.db.Exec(query, orderProduct.Id, orderProduct.OrderId, orderProduct.ProductId, orderProduct.DiscountType, orderProduct.Quantity)
	if err != nil {
		return nil, err
	}

	return &orderProduct, nil
}

func (c *OrderProductRepo) Delete(req model.GetOrderProductByOrderId) (string, error) {

	_, err := c.db.Exec(`DELETE FROM order_products WHERE id = $1`, req.Id)

	if err != nil {
		return "Can not delete", err
	}

	return "OrderProduct deleted", nil
}

func (p *OrderProductRepo) GetByOrderId(req model.GetOrderProductByOrderId) (*[]model.OrderProducts, error) {
	var order_product []model.OrderProducts
	var query = `
		SELECT
			id, order_id, product_id, discount_type, discount_amount, quantity, price, sum, created_at, updated_at 
		FROM order_products
		WHERE order_id = $1
	`
	rows, err := p.db.Query(query, req.Id)
	if err != nil {
		return &[]model.OrderProducts{}, err
	}

	for rows.Next() {
		var order_p = model.OrderProducts{}
		rows.Scan(
			&order_p.Id,
			&order_p.OrderId,
			&order_p.ProductId,
			&order_p.DiscountType,
			&order_p.DiscountAmount,
			&order_p.Quantity,
			&order_p.Price,
			&order_p.Sum,
			&order_p.CreatedAt,
			&order_p.UpdatedAt,
		)
		order_product = append(order_product, order_p)
	}
	rows.Close()

	return &order_product, nil
}
