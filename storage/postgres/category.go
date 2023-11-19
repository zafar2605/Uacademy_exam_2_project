package postgres

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"market/helpers"
	"market/model"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (c *CategoryRepo) Create(req model.CategoryCreate) (*model.Category, error) {

	var category = model.Category{
		Id:       uuid.New().String(),
		Title:    req.Title,
		Image:    req.Image,
		ParentID: req.ParentID,
	}

	var query = `INSERT INTO category(id, title, image, parent_id ,updated_at) VALUES($1, $2, $3, $4, NOW())`

	_, err := c.db.Exec(query, category.Id, category.Title, category.Image, helpers.NewNullString(category.ParentID))
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *CategoryRepo) GetById(req model.CategoryPrimaryKey) (*model.Category, error) {

	var category = model.Category{}

	var query = `
		SELECT 
			id, 
			title,
			COALESCE(CAST(parent_id AS VARCHAR), ''),
			created_at, 
			updated_at
		FROM 
			category
		WHERE id = $1
	`

	resp := c.db.QueryRow(query, req.Id)

	if resp.Err() != nil {
		return nil, resp.Err()
	}
	err := resp.Scan(
		&category.Id,
		&category.Title,
		&category.ParentID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *CategoryRepo) GetList(req model.GetListCategoryRequest) (*model.GetListCategoryResponse, error) {

	var resp = model.GetListCategoryResponse{}
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
			COUNT(*) OVER(), id, title, parent_id, created_at,  updated_at
		FROM 
			category
	`
	query += offset + limit
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category = model.Category{}
		rows.Scan(
			&resp.Count,
			&category.Id,
			&category.Title,
			&category.ParentID,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		resp.Categories = append(resp.Categories, category)
	}
	rows.Close()

	return &resp, nil
}

func (c *CategoryRepo) Update(req model.CategoryUpdate) (string, error) {

	_, err := c.db.Exec(`
		UPDATE 
			category 
		SET 
			title = $1, image = $2, parent_id = $3, updated_at = NOW()
	 	WHERE id = $4`, req.Title, req.Image, helpers.NewNullString(req.ParentID), req.Id)

	if err != nil {
		return "Can not update", err
	}

	return "Updated", nil
}

func (c *CategoryRepo) Delete(req model.CategoryPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM category WHERE id = $1`, req.Id)

	if err != nil {
		return "Can not delete", err
	}

	return "Deleted", nil

}
