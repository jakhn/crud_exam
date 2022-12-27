package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"crud/models"
	"crud/pkg/helper"
)

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (f *CategoryRepo) Create(ctx context.Context, category *models.CreateCategory) (string, error) {

	var (
		id     = uuid.New().String()
		query  string
		nullId sql.NullString
	)

	query = `
		INSERT INTO category(
			category_id,
			parent_id,
			category_name, 
			updated_at
		) VALUES ( $1, $2 , $3, now())
	`

	if category.ParentId == "" {
		_, err := f.db.Exec(ctx, query,
			id,
			nullId,
			category.CategoryName,
		)

		if err != nil {
			return "", err
		}
	} else {

		_, err := f.db.Exec(ctx, query,
			id,
			category.ParentId,
			category.CategoryName,
		)

		if err != nil {
			return "", err
		}

	}

	return id, nil
}

func (f *CategoryRepo) GetByPKey(ctx context.Context, pkey *models.CategoryPrimarKey) (*models.Cp, error) {

	var (
		resp              = models.Cp{}
		respCategoryChild = models.ChildCategory{}
		categoryId        string
	)

	query := `
		SELECT
		category_id,
		category_name,
		parent_id
		FROM
		category WHERE id = $1 and deleted_at = null;
		`
	queryName := `select product_name from product where category_id = $1`

	queryAll := `select count(*) over() from category where parent_id=$1`

	rows, err := f.db.Query(ctx, query, pkey.Id)

	for rows.Next() {
		res := &models.CategoryByParent{}

		err := rows.Scan(
			&res.Id,
			&res.Name,
			&res.ParentId,
		)

		if err != nil {
			return nil, err
		}
		categoryId = res.Id
		respCategoryChild.ParentId = res.ParentId
		respCategoryChild.ChildsCategory = append(respCategoryChild.ChildsCategory, res)
	}
	err = f.db.QueryRow(ctx, queryName, categoryId).Scan(&respCategoryChild.Name)
	if err != nil {
		return nil, err
	}
	rows, err = f.db.Query(ctx, queryAll, respCategoryChild.ParentId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(
			&resp.Count,
		)
		resp.Categories = append(resp.Categories, &respCategoryChild)
	}

	return &resp, err
}

func (f *CategoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.Cp, error) {

	var (
		resp              = models.Cp{}
		respCategoryChild = models.ChildCategory{}
		offset            = " OFFSET 0"
		limit             = " LIMIT 20"
		categoryId        string
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
  SELECT
  category_id,
  category_name,
  parent_id
  FROM
  category WHERE deleted_at = 0;
  `
	queryName := `select product_name from product where category_id = $1`

	queryAll := `select count(*) over() from category where parent_id=$1`
	query += offset + limit

	rows, err := f.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		res := &models.CategoryByParent{}

		err := rows.Scan(
			&res.Id,
			&res.Name,
			&res.ParentId,
		)

		if err != nil {
			return nil, err
		}
		categoryId = res.Id
		respCategoryChild.ParentId = res.ParentId
		respCategoryChild.ChildsCategory = append(respCategoryChild.ChildsCategory, res)
	}
	err = f.db.QueryRow(ctx, queryName, categoryId).Scan(&respCategoryChild.Name)
	if err != nil {
		return nil, err
	}
	rows, err = f.db.Query(ctx, queryAll, respCategoryChild.ParentId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(
			&resp.Count,
		)
		resp.Categories = append(resp.Categories, &respCategoryChild)
	}

	return &resp, err
}

func (f *CategoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			category
		SET
			parent_id = :parent_id,
			category_name = :category_name, 
			updated_at = now()
		WHERE category_id = :category_id
	`

	params = map[string]interface{}{
		"category_id":   req.Id,
		"parent_id":     req.ParentId,
		"category_name": req.CategoryName,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *CategoryRepo) Delete(ctx context.Context, req *models.CategoryPrimarKey) error {

	_, err := f.db.Exec(ctx, "UPDATE category SET deleted_at = now(), is_deleted = true WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
