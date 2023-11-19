package model

type Branch struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Photo         string `json:"photo"`
	Adress        string `json:"adress"`
	Active        string `json:"active"`
	WorkStartHour string `json:"work_start_hour"`
	WorkEndHour   string `json:"work_end_hour"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type BranchPrimaryKey struct {
	Id string `json:"id"`
}

type BranchCreate struct {
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Photo         string `json:"photo"`
	Adress        string `json:"adress"`
	Active        string `json:"active"`
	WorkStartHour string `json:"work_start_hour"`
	WorkEndHour   string `json:"work_end_hour"`
}

type BranchUpdate struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Photo         string `json:"photo"`
	Adress        string `json:"adress"`
	Active        string `json:"active"`
	WorkStartHour string `json:"work_start_hour"`
	WorkEndHour   string `json:"work_end_hour"`
}

type GetListBranchRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListBranchResponse struct {
	Count    int      `json:"count"`
	Branches []Branch `json:"branches"`
}
