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

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (f *ProductRepo) Create(ctx context.Context, product *models.CreateProduct) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO product(
			product_id,
			product_name, 
			price, 
			category_id, 
			updated_at
		) VALUES ( $1, $2 , $3, $4, now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		product.ProductName,
		product.Price,
		product.CategoryId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *ProductRepo) GetByPKey(ctx context.Context, pkey *models.ProductPrimarKey) (*models.Product, error) {

	var (
		id                sql.NullString
		productName       sql.NullString
		productPrice      sql.NullInt64
		productCategoryId sql.NullString
		createdAt         sql.NullString
		updatedAt         sql.NullString
	)

	query := `
		SELECT
			product_id,
			product_name,
			price,
			category_id, 
			created_at,
			updated_at
		FROM
			product
		WHERE product_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&productName,
			&productPrice,
			&productCategoryId,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Product{
		Id:          id.String,
		ProductName: productName.String,
		Price:       productPrice.Int64,
		CategoryId:  productCategoryId.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (f *ProductRepo) GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error) {

	var (
		resp   = models.GetListProductResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 2"
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
		product_id,
		product_name,
		price, 
		category_id,
		created_at,
		updated_at
	FROM
		product WHERE deleted_at = 0;
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id          sql.NullString
			productName sql.NullString
			price       sql.NullInt64
			categoryId  sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&productName,
			&price,
			&categoryId,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &models.Product{
			Id:          id.String,
			ProductName: productName.String,
			Price:       price.Int64,
			CategoryId:  categoryId.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})

	}

	return &resp, err
}

func (f *ProductRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			product
		SET
			product_name = :product_name,
			price = :price,
			category_id = :category_id, 
			updated_at = now()
		WHERE product_id = :product_id
	`

	params = map[string]interface{}{
		"product_id":   req.Id,
		"product_name": req.ProductName,
		"price":        req.Price,
		"category_id":  req.CategoryId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *ProductRepo) Delete(ctx context.Context, req *models.ProductPrimarKey) error {

	_, err := f.db.Exec(ctx, "UPDATE orders SET deleted_at = now(), is_deleted = true WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil 
}

func (uh userHandler) FilterProduct (w http.ResponseWriter, r *http.Request) {
    users := []User{}
    // sortBy is expected to look like field.orderdirection i. e. id.asc
    sortBy := r.URL.Query().Get("sortBy")
    if sortBy == "" {
        // id.asc is the default sort query
        sortBy = "id.asc"
    }
    sortQuery, err := validateAndReturnSortQuery(sortBy)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    strLimit := r.URL.Query().Get("limit")
    // with a value as -1 for gorms Limit method, we'll get a request without limit as default
    limit := -1
    if strLimit != "" {
        limit, err = strconv.Atoi(strLimit)
        if err != nil || limit < -1 {
            http.Error(w, "limit query parameter is no valid number", http.StatusBadRequest)
            return
        }
    }
    strOffset := r.URL.Query().Get("offset")
    offset := -1
    if strOffset != "" {
        offset, err = strconv.Atoi(strOffset)
        if err != nil || offset < -1 {
            http.Error(w, "offset query parameter is no valid number", http.StatusBadRequest)
            return
        }
    }
    filter := r.URL.Query().Get("filter")
    filterMap := map[string]string{}
    if filter != "" {
        filterMap, err = validateAndReturnFilterMap(filter)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
    }
    if err := uh.db.Where(filterMap).Limit(limit).Offset(offset).Order(sortQuery).Find(&users).Error; err != nil {
        fmt.Println(err)
        http.Error(w, "Error on DB find for all users", http.StatusInternalServerError)
        return
    }
    w.Header().Add("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(users); err != nil {
        fmt.Println(err)
        http.Error(w, "Error encoding response object", http.StatusInternalServerError)
    }
}
func validateAndReturnFilterMap(filter string) (map[string]string, error) {
    splits := strings.Split(filter, ".")
    if len(splits) != 2 {
        return nil, errors.New("malformed sortBy query parameter, should be field.orderdirection")
    }
    field, value := splits[0], splits[1]
    if !stringInSlice(userFields, field) {
        return nil, errors.New("unknown field in filter query parameter")
    }
    return map[string]string{field: value}, nil
}


    filter := r.URL.Query().Get("filter")
    filterMap := map[string]string{}
    if filter != "" {
        filterMap, err = validateAndReturnFilterMap(filter)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
    }

