package postgres

import (
	"database/sql"
	"errors"
	"strconv"

	"market/helpers"
	"market/model"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (p *OrderRepo) Create(req model.OrderCreate) (*model.Order, error) {
	resp, err := p.GetLast()
	if err != nil {
		return &model.Order{}, err
	}
	last_number, _ := strconv.Atoi(resp)

	var order = model.Order{
		Id:            "O-" + helpers.GetSerialId(last_number),
		ClientId:      req.Client,
		BranchId:      req.Branch,
		Address:       req.Address,
		DeliveryPrice: 0,
		TotalCount:    0,
		TotalPrice:    0,
		Status:        "New",
	}

	var query = `INSERT INTO orders(id, client_id, branch_id, address, delivery_price, total_count, total_price, status, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, NOW())`

	_, err = p.db.Exec(query, order.Id, order.ClientId, order.BranchId, order.Address, order.DeliveryPrice, order.TotalCount, order.TotalPrice, order.Status)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *OrderRepo) GetById(req model.OrderPrimaryKey) (*model.Order, error) {

	var Order = model.Order{}

	var query = `
		SELECT 
			id, client, branch, address, status, created_at, updated_at
		FROM 
			orders
		WHERE id = $1
	`
	resp := c.db.QueryRow(query, req.Id)
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	err := resp.Scan(
		&Order.Id,
		&Order.ClientId,
		&Order.BranchId,
		&Order.Address,
		&Order.Status,
		&Order.CreatedAt,
		&Order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &Order, nil
}

func (c *OrderRepo) GetList(req model.GetListOrderRequest) (*model.GetListOrderResponse, error) {

	var resp = model.GetListOrderResponse{}
	// var query = `
	// 	SELECT
	// 		COUNT(*) OVER(), o.id, o.client_id, o.branch_id, o.address, o.status, o.delivery_price, o.total_count, o.total_price,  o.created_at, o.updated_at,
	// 		op.id, op.order_id, op.product_id, op.discount_type, op.discount_amount, op.quantity, op.price, op.sum, op.created_at, op.updated_at
	// 	FROM
	// 		orders AS o
	// 	JOIN order_products AS op
	// `
	var query = `
		SELECT 
			COUNT(*) OVER(),id,client_id,branch_id,address,status,delivery_price,total_count,total_price, created_at,updated_at
		FROM 
			orders
	`
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order = model.Order{}
		rows.Scan(
			&resp.Count,
			&order.Id,
			&order.ClientId,
			&order.BranchId,
			&order.Address,
			&order.Status,
			&order.DeliveryPrice,
			&order.TotalCount,
			&order.TotalPrice,
			&order.CreatedAt,
			&order.UpdatedAt,

			// &order.OrderProducts,

		)
		resp.Orders = append(resp.Orders, order)
	}
	rows.Close()

	return &resp, nil
}

func (c *OrderRepo) Update(req model.OrderUpdate) (string, error) {

	_, err := c.db.Exec(`
		UPDATE
			orders
		SET
			id = $1, client_id = $2, branch_id = $3, address = $4, created_at = $5, updated_at = NOW()
	 	WHERE id = $6`, req.ClientId, req.BranchId, req.Address)

	if err != nil {
		return "Can not update", err
	}

	return "Updated", nil
}

func (c *OrderRepo) StatusUpdate(req model.OrderStatusUpdate) (string, error) {

	var last_status string
	resp := c.db.QueryRow(`SELECT status FROM orders WHERE id = $1`, req.Id)
	if resp.Err() != nil {
		return "Can not update status", resp.Err()
	}
	err := resp.Scan(&last_status)

	if err != nil {
		return "Can not update status", err
	}

	if req.Status == "In_proccess" && last_status != "New" {
		return "Do not true request", errors.New("do not true request")
	} else if req.Status == "Success" && last_status != "In_proccess" {
		return "Do not true request", errors.New("do not true request")
	} else if req.Status == "Cencel" && last_status != "New" {
		return "Do not true request", errors.New("do not true request")
	}
	_, err = c.db.Exec(`
		UPDATE
			orders
		SET
			status = $1, updated_at = NOW()
	 	WHERE id = $2`, req.Status, req.Id)

	if err != nil {
		return "Can not updated status", err
	}

	return "Status updated", nil
}

func (c *OrderRepo) Delete(req model.OrderPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM orders WHERE id = $1`, req.Id)

	if err != nil {
		return "Can not delete", err
	}

	return "Deleted", nil
}

func (c *OrderRepo) GetLast() (string, error) {
	var count string
	resp := c.db.QueryRow(`SELECT id FROM orders ORDER BY id DESC LIMIT 1`)

	if resp.Err() != nil {
		return "O-00000000", resp.Err()
	}
	err := resp.Scan(
		&count,
	)
	if err != nil {
		return "Can not scan from db", nil
	}

	return count[2:], nil
}
