package model

type OrderProducts struct {
	Id             string  `json:"id"`
	OrderId        string  `json:"order_id"`
	ProductId      string  `json:"product_id"`
	DiscountType   string  `json:"discount_type"`
	DiscountAmount string  `json:"discount_amount"`
	Quantity       int     `json:"quantity"`
	Price          float64 `json:"price"`
	Sum            float64 `json:"sum"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type AddOrderProduct struct {
	Id             string  `json:"id"`
	OrderId        string  `json:"order_id"`
	ProductId      string  `json:"product_id"`
	DiscountType   string  `json:"discount_type"`
	DiscountAmount string  `json:"discount_amount"`
	Quantity       int     `json:"quantity"`
	Price          float64 `json:"price"`
	Sum            float64 `json:"sum"`
}

type GetOrderProductByOrderId struct {
	Id string `json:"id"`
}
