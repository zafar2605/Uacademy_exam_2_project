package model

type Order struct {
	Id            string          `json:"id"`
	ClientId      string          `json:"client_id"`
	BranchId      string          `json:"branch_id"`
	Address       string          `json:"address"`
	DeliveryPrice float64         `json:"delivery_price"`
	TotalCount    int             `json:"total_count"`
	TotalPrice    float64         `json:"total_price"`
	Status        string          `json:"status"`
	CreatedAt     string          `json:"created_at"`
	UpdatedAt     string          `json:"updated_at"`
	OrderProducts []OrderProducts `json:"order_products"`
}

type OrderCreate struct {
	Client  string `json:"order"`
	Branch  string `json:"branch"`
	Address string `json:"address"`
}

type OrderStatusUpdate struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type OrderUpdate struct {
	Id       string `json:"id"`
	ClientId string `json:"client_id"`
	BranchId string `json:"branch_id"`
	Address  string `json:"address"`
	Status   string `json:"status"`
}

type GetListOrderRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListOrderResponse struct {
	Count  int     `json:"count"`
	Orders []Order `json:"orders"`
}
