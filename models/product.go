package models

import "gorm.io/plugin/soft_delete"

type ProductPrimarKey struct {
	Id string `json:"product_id"`
}

type CreateProduct struct {
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	CategoryId  string `json:"category_id"`
}

type Product struct {
	Id          string                `json:"product_id"`
	ProductName string                `json:"product_name"`
	Price       int64                 `json:"price"`
	CategoryId  string                `json:"category_id"`
	CreatedAt   string                `json:"created_at"`
	UpdatedAt   string                `json:"updated_at"`
	DeletedAt   soft_delete.DeletedAt `json:"gorm:"softDelete:milli""`
}

type UpdateProduct struct {
	Id          string `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	CategoryId  string `json:"category_id"`
}

type GetListProductRequest struct {
	Limit  int32
	Offset int32
}

type GetListProductResponse struct {
	Count    int32      `json:"count"`
	Products []*Product `json:"products"`
}
