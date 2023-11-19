package model

type Category struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	ParentID  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryCreate struct {
	Title    string `json:"title"`
	Image    string `json:"image"`
	ParentID string `json:"parent_id"`
}

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}

type CategoryUpdate struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	ParentID string `json:"parent_id"`
}

type GetListCategoryRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListCategoryResponse struct {
	Count      int        `json:"count"`
	Categories []Category `json:"categories"`
}
