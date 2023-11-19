package postgres

import (
	"database/sql"
	"time"

	"market/model"

	"github.com/google/uuid"
)

type BranchRepo struct {
	db *sql.DB
}

func NewBranchRepo(db *sql.DB) *BranchRepo {
	return &BranchRepo{
		db: db,
	}
}

func (p *BranchRepo) Create(req model.BranchCreate) (*model.Branch, error) {
	var branch = model.Branch{
		Id:            uuid.New().String(),
		Name:          req.Name,
		Phone:         req.Phone,
		Photo:         req.Photo,
		Adress:        req.Adress,
		Active:        req.Active,
		WorkStartHour: req.WorkStartHour,
		WorkEndHour:   req.WorkEndHour,
	}
	var query = `INSERT INTO branch (id, name, phone, photo, address, active, work_start_hour, work_end_hour, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, NOW())`

	_, err := p.db.Exec(query, branch.Id, branch.Name, branch.Phone, branch.Photo, branch.Adress, branch.Active, branch.WorkStartHour, branch.WorkEndHour)
	if err != nil {
		return nil, err
	}

	return &branch, nil
}

func (c *BranchRepo) GetById(req model.BranchPrimaryKey) (*model.Branch, error) {

	var Branch = model.Branch{}

	var query = `
		SELECT 
			id, name, phone, photo, address, active, work_start_hour, work_start_hour, updated_at) VALUES($1, $2, $3, $4, $5, $6, NOW()),
		FROM 
			branch
		WHERE id = $1
	`
	resp := c.db.QueryRow(query, req.Id)
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	err := resp.Scan(
		&Branch.Id,
		&Branch.Name,
		&Branch.Phone,
		&Branch.Photo,
		&Branch.Adress,
		&Branch.Active,
		&Branch.WorkStartHour,
		&Branch.WorkEndHour,
		&Branch.CreatedAt,
		&Branch.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &Branch, nil
}

func (c *BranchRepo) GetList(req model.GetListBranchRequest) (*model.GetListBranchResponse, error) {

	var resp = model.GetListBranchResponse{}

	var query = `
		SELECT 
			id, name, phone, photo, address, active, work_start_hour, work_end_hour, created_at, updated_at
		FROM 
			branch
	`
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var branch = model.Branch{}
		rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Phone,
			&branch.Photo,
			&branch.Adress,
			&branch.Active,
			&branch.WorkStartHour,
			&branch.WorkEndHour,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		)
		branch.WorkStartHour = branch.WorkStartHour[12:17]
		branch.WorkEndHour = branch.WorkEndHour[12:17]
		var now_time = time.Now().Format("15:04:05")

		if branch.WorkStartHour[11:19] <= now_time && now_time <= branch.WorkEndHour[11:19] {
			resp.Branches = append(resp.Branches, branch)
		}
		resp.Count = len(resp.Branches)
	}
	rows.Close()

	return &resp, nil
}

func (c *BranchRepo) Update(req model.BranchUpdate) (string, error) {

	_, err := c.db.Exec(`
		UPDATE branch
		SET name = $1, phone = $2, photo = $3, address = $4, active = $5, work_start_hour = $6, work_start_hour = $7, updated_at = NOW()
	 	WHERE id = $8`, req.Name, req.Phone, req.Photo, req.Adress, req.Active, req.WorkStartHour, req.WorkEndHour, req.Id)

	if err != nil {
		return "Can not update", err
	}

	return "Updated", nil
}

func (c *BranchRepo) Delete(req model.BranchPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM branch WHERE id = $1`, req.Id)

	if err != nil {
		return "Can not delete", err
	}

	return "Deleted", nil
}
