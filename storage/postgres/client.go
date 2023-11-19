package postgres

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"market/model"
)

type ClientRepo struct {
	db *sql.DB
}

func NewClientRepo(db *sql.DB) *ClientRepo {
	return &ClientRepo{
		db: db,
	}
}

func (c *ClientRepo) Create(req model.ClientCreate) (*model.Client, error) {

	var client = model.Client{
		Id:          uuid.New().String(),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Phone:       req.Phone,
		Photo:       req.Photo,
		DateOfBirth: req.DateOfBirth,
	}

	var query = `INSERT INTO client(id, firstname, lastname, phone, photo, date_of_birth, updated_at) VALUES($1, $2, $3, $4, $5, $6, NOW())`

	_, err := c.db.Exec(query, client.Id, client.FirstName, client.LastName, client.Phone, client.Photo, client.DateOfBirth)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (c *ClientRepo) GetById(req model.ClientPrimaryKey) (*model.Client, error) {

	var client = model.Client{}

	var query = `
		SELECT id, firstname, lastname, phone, photo, date_of_birth, created_at, updated_at
		FROM client
		WHERE id = $1
	`

	resp := c.db.QueryRow(query, req.Id)

	if resp.Err() != nil {
		return nil, resp.Err()
	}
	err := resp.Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.Phone,
		&client.Photo,
		&client.CreatedAt,
		&client.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *ClientRepo) GetList(req model.GetListClientRequest) (*model.GetListClientResponse, error) {

	var resp = model.GetListClientResponse{}
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
			COUNT(*) OVER(), id, firstname, lastname, phone, photo, date_of_birth, created_at,  updated_at
		FROM 
			client
	`
	query += offset + limit
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var client = model.Client{}
		rows.Scan(
			&resp.Count,
			&client.Id,
			&client.FirstName,
			&client.LastName,
			&client.Phone,
			&client.Photo,
			&client.DateOfBirth,
			&client.CreatedAt,
			&client.UpdatedAt,
		)

		resp.Clients = append(resp.Clients, client)
	}
	rows.Close()

	return &resp, nil
}

func (c *ClientRepo) Update(req model.ClientUpdate) (string, error) {

	_, err := c.db.Exec(`
		UPDATE 
			client 
		SET
			firstname = $1, lastname = $2, phone = $3, photo = $4, date_of_birth = $5, updated_at = NOW()
	 	WHERE id = $6`, req.FirstName, req.LastName, req.Phone, req.Photo, req.DateOfBirth, req.Id)

	if err != nil {
		return "Can not update", err
	}

	return "Updated", nil
}

func (c *ClientRepo) Delete(req model.ClientPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM client WHERE id = $1`, req.Id)

	if err != nil {
		return "Can not delete", err
	}

	return "Deleted", nil

}
