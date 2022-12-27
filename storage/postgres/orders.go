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

type OrdersRepo struct {
	db *pgxpool.Pool
}

func NewOrdersRepo(db *pgxpool.Pool) *OrdersRepo {
	return &OrdersRepo{
		db: db,
	}
}

func (f *OrdersRepo) Create(ctx context.Context, orders *models.CreateOrders) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO orders(
			orders_id,
			description, 
			poduct_id, 
			updated_at
		) VALUES ( $1, $2 , $3, now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		orders.Description,
		orders.ProductId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *OrdersRepo) GetByPKey(ctx context.Context, pkey *models.OrdersPrimarKey) (*models.Orders, error) {

	
	var (
		respMain       = models.Ords{}
		respChildOrd   = models.ChildOrder{}
		respProducts   = models.Products{} 
		offset         = " OFFSET 0"
		limit          = " LIMIT 20"
		categoryId     string
	) 

	query := `
	SELECT
	orders_id,
	description,
	product_id, 
	FROM
	orders WHERE deleted_at = 0;
	`
	queryPrentId := `select parent_id from category where orders_id = $1`

	queryAll := `select count(*) over() from orders where orders_id = $1`
	queryProduct := `select product_id from product where product_id = $1`
	

	rows, err := f.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		res := &models.Categories{}

		err := rows.Scan(
			&res.CategoryId,
			&res.CategoryName,
			&res.ParentId,
		)

		if err != nil {
			return nil, err
		}
		categoryId = res.CategoryId
		respProducts.ProductId = res.ProductId
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
}

func (f *OrdersRepo) GetList(ctx context.Context, req *models.GetListOrdersRequest) (*models.Cp, error) {

	var (
		respMain       = models.Ords{}
		respChildOrd   = models.ChildOrder{}
		respProducts   = models.Products{} 
		offset         = " OFFSET 0"
		limit          = " LIMIT 20"
		categoryId     string
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
	SELECT
	orders_id,
	description,
	product_id, 
	FROM
	orders WHERE deleted_at = 0;
	`
	queryPrentId := `select parent_id from category where orders_id = $1`

	queryAll := `select count(*) over() from orders where orders_id = $1`
	queryProduct := `select product_id from product where product_id = $1`
	query += offset + limit

	rows, err := f.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		res := &models.Categories{}

		err := rows.Scan(
			&res.CategoryId,
			&res.CategoryName,
			&res.ParentId,
		)

		if err != nil {
			return nil, err
		}
		categoryId = res.CategoryId
		respProducts.ProductId = res.ProductId
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

func (f *OrdersRepo) Update(ctx context.Context, req *models.UpdateOrders) (int64, error) {

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
		WHERE orders_id = :orders_id
	`

	params = map[string]interface{}{
		"orders_id":   req.Id,
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

func (f *OrdersRepo) Delete(ctx context.Context, req *models.OrdersPrimarKey) error {

	_, err := f.db.Exec(ctx, "UPDATE orders SET deleted_at = now(), is_deleted = true WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}


	