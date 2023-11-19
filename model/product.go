package model

type Product struct {
	Id          string   `json:"id"`
	CategoryId  string   `json:"category_id"`
	Title       string   `json:"title"`
	Photos      string   `json:"photos"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Category    Category `json:"category"`
}

type ProductCreate struct {
	CategoryId  string  `json:"category_id"`
	Title       string  `json:"title"`
	Photos      string  `json:"photos"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type ProductUpdate struct {
	Id          string  `json:"id"`
	CategoryId  string  `json:"category_id"`
	Title       string  `json:"title"`
	Photos      string  `json:"photos"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type GetListProductRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListProductResponse struct {
	Count    int       `json:"count"`
	Products []Product `json:"products"`
}
