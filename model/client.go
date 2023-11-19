package model

type Client struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	DateOfBirth string `json:"date_of_birth"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ClientCreate struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	DateOfBirth string `json:"date_of_birth"`
}

type ClientPrimaryKey struct {
	Id string `json:"id"`
}

type ClientUpdate struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	DateOfBirth string `json:"date_of_birth"`
}

type GetListClientRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListClientResponse struct {
	Count   int      `json:"count"`
	Clients []Client `json:"clients"`
}
