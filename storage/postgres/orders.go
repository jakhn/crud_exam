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
func (f *OrdersRepo) GetByPKey(ctx context.Context, pkey *models.OrdersPrimarKey) (*models.OrderList, error) { 

	var (
		productCategory models.ProductCategory
		productList     models.ProductList
		orderList       models.OrderList

		orderId          sql.NullString
		orderDescription sql.NullString
		productId        sql.NullString
		productName      sql.NullString
		categoryId       sql.NullString
		categoryName     sql.NullString
		categoryParentId sql.NullString
	)

	query := `
	SELECT
		orders.orders_id,
		orders.description,
		product.product_id,
		product.product_name,
		category.category_id,
		category.category_name,
		category.parent_id
	FROM
    	orders
	JOIN product ON orders.product_id = product.product_id
	JOIN category ON product.category_id = category.category_id
	WHERE orders.is_deleted = false AND product.is_deleted = false AND category.is_deleted = false AND orders.orders_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).Scan(
		&orderId,
		&orderDescription,
		&productId,
		&productName,
		&categoryId,
		&categoryName,
		&categoryParentId,
	)

	productCategory.Id = categoryId.String
	productCategory.Name = categoryName.String
	productCategory.ParentID = categoryParentId.String

	productList.Id = productId.String
	productList.Name = productName.String
	productList.Category = productCategory

	orderList.Id = orderId.String
	orderList.Description = orderDescription.String
	orderList.Product = productList

	return &orderList, err
}


func (f *OrdersRepo) GetList(ctx context.Context, req *models.GetListOrdersRequest) (*models.GetListOrdersResponse, error) {
	var (
		resp   = models.GetListOrdersResponse{}
		offset = ""
		limit  = ""
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
	
	SELECT
		COUNT(*) OVER(),
		orders.orders_id,
		orders.description,
		product.product_id,
		product.product_name,
		category.category_id,
		category.category_name,
		category.parent_id
	FROM
    	orders
	JOIN product ON orders.product_id = products.product_id
	JOIN category ON products.category_id = category.category_id
	WHERE orders.is_deleted = false AND products.is_deleted = false AND categories.is_deleted = false
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {
		var (
			productCategory models.ProductCategory
			productList     models.ProductList

			orderId          sql.NullString
			orderDescription sql.NullString
			productId        sql.NullString
			productName      sql.NullString
			categoryId       sql.NullString
			categoryName     sql.NullString
			categoryParentId sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&orderId,
			&orderDescription,
			&productId,
			&productName,
			&categoryId,
			&categoryName,
			&categoryParentId,
		)
		if err != nil {
			return nil, err
		}

		productCategory.Id = categoryId.String
		productCategory.Name = categoryName.String
		productCategory.ParentID = categoryParentId.String

		productList.Id = productId.String
		productList.Name = productName.String
		productList.Category = productCategory

		resp.Order = append(resp.Order, &models.OrderList{
			Id:          orderId.String,
			Description: orderDescription.String,
			Product:     productList,
		})

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


	