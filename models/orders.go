package models

type OrdersPrimarKey struct {
	Id string `json:"orders_id"`
}

type CreateOrders struct {
	Description string `json:"description"`
	ProductId   string `json:"product_id"`
}

type Orders struct {
	Id          string `json:"orders_id"`
	Description string `json:"description"`
	ProductId   string `json:"product_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type UpdateOrders struct {
	Id          string `json:"orders_id"`
	Description string `json:"description"`
	ProductId   string `json:"product_id"`
}

type GetListOrdersRequest struct {
	Limit  int32
	Offset int32
}

type GetListOrdersResponse struct {
	Count int32     `json:"count"`
	Order []*OrderList `json:"orderss"`
}

type OrderList struct {
	Id          string      `json:"orders_id"`
	Description string      `json:"description"`
	Product     ProductList `json:"product"`
}
type ProductList struct {
	Id       string          `json:"product_id"`
	Name     string          `json:"name"`
	Category ProductCategory `json:"category"`
}
type ProductCategory struct {
	Id       string `json:"category_id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}
