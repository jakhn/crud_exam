package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"crud/models"
	"crud/pkg/helper"
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (f *OrderRepo) Create(ctx context.Context, order *models.CreateOrder) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO orders(
			id,
			description,
			product_id,
			updated_at
		) VALUES ( $1, $2, $3, now() )
	`
	_, err := f.db.Exec(ctx, query,
		id,
		order.Description,
		order.ProductId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *OrderRepo) GetByPKey(ctx context.Context, pkey *models.OrderPrimarKey) (*models.OrderList, error) {

	var (
		id          sql.NullString
		description sql.NullString
		productId   sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
		orderResp   models.Categories
	)

	query := `
		SELECT
			id,
			description,
			product_id,
			created_at,
			updated_at
		FROM orders
		WHERE id = $1 AND deleted_at IS NULL
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).Scan(
		&id,
		&description,
		&productId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	resp := &models.OrderList{
		Id:          id.String,
		Description: description.String,
		Product:     &models.ProductList{Id: productId.String},
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}

	var (
		name       sql.NullString
		categoryId sql.NullString
	)

	queryProduct := `
		SELECT
			id,
			name,
			category_id,
			created_at,
			updated_at
		FROM products
		WHERE id = $1 AND deleted_at IS NULL
	`

	err = f.db.QueryRow(ctx, queryProduct, resp.Product.Id).Scan(
		&productId,
		&name,
		&categoryId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	response := &models.ProductList{
		Id:        id.String,
		Name:      name.String,
		Category:  &models.Categories{Id: categoryId.String},
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	var (
		parentId sql.NullString
	)
	queryCategory := `
	SELECT
		id,
		name,
		parent_id,
		created_at,
		updated_at
	FROM categories
	WHERE id = $1 AND deleted_at IS NULL
`

	rows, err := f.db.Query(ctx, queryCategory, response.Category.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return resp, nil
		}

		return nil, err
	}

	for rows.Next() {

		err = rows.Scan(
			&id,
			&name,
			&parentId,
			&createdAt,
			&updatedAt,
		)

		orderResp.Order = append(orderResp.Order, &models.OrderList{
			Id:          id.String,
			Description: description.String,
			Product:     resp.Product,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return orderResp.Order[0], nil
}

func (f *OrderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {

	var (
		filter string
		params = make(map[string]interface{})
	)
	params["offset"] = req.Offset
	params["limit"] = req.Limit

	var (
		id          sql.NullString
		description sql.NullString
		productId   sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
		orderResp   models.Categories
		count       int
	)

	query := `
		SELECT
			count(*) over(),
			id,
			description,
			product_id,
			created_at,
			updated_at
		FROM orders
		WHERE deleted_at IS NULL
	` + filter

	err := f.db.QueryRow(ctx, query).Scan(
		&count,
		&id,
		&description,
		&productId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	resp := &models.OrderList{
		Id:          id.String,
		Description: description.String,
		Product:     &models.ProductList{Id: productId.String},
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}

	var (
		name       sql.NullString
		categoryId sql.NullString
	)

	queryProduct := `
		SELECT
			id,
			name,
			category_id,
			created_at,
			updated_at
		FROM products
		WHERE id = $1 AND deleted_at IS NULL
	`

	err = f.db.QueryRow(ctx, queryProduct, resp.Product.Id).Scan(
		&productId,
		&name,
		&categoryId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	response := &models.ProductList{
		Id:        id.String,
		Name:      name.String,
		Category:  &models.Categories{Id: categoryId.String},
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	var (
		parentId sql.NullString
	)
	queryCategory := `
	SELECT
		id,
		name,
		parent_id,
		created_at,
		updated_at
	FROM categories
	WHERE id = $1 AND deleted_at IS NULL
`

	rows, err := f.db.Query(ctx, queryCategory, response.Category.Id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err = rows.Scan(
			&id,
			&name,
			&parentId,
			&createdAt,
			&updatedAt,
		)

		orderResp.Order = append(orderResp.Order, &models.OrderList{
			Id:          id.String,
			Description: description.String,
			Product:     resp.Product,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return &models.GetListOrderResponse{
		Count:  count,
		Orders: orderResp.Order,
	}, nil
}

func (f *OrderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			orders
		SET
			description = :description,
			product_id = :product_id,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"description": req.Description,
		"product_id":  req.ProductId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *OrderRepo) Delete(ctx context.Context, req *models.OrderPrimarKey) error {

	_, err := f.db.Exec(ctx, "UPDATE orders SET deleted_at = now() WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
