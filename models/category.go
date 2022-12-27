package models

import "gorm.io/plugin/soft_delete"

type CategoryPrimarKey struct {
	Id string `json:"category_id"`
}

type CreateCategory struct {
	CategoryName string `json:"category_name"`
	ParentId     string `json:"parent_id"`
}

type Category struct {
	Id           string                `json:"category_id"`
	ParentId     string                `json:"parent_id"`
	CategoryName string                `json:"category_name"`
	CreatedAt    string                `json:"created_at"`
	UpdatedAt    string                `json:"updated_at"`
	DeletedAt    soft_delete.DeletedAt `json:"gorm:"softDelete:milli""`
}

type UpdateCategory struct {
	Id           string `json:"category_id"`
	ParentId     string `json:"parent_id"`
	CategoryName string `json:"category_name"`
}

type GetListCategoryRequest struct {
	Limit  int32
	Offset int32
}

type GetListCategoryResponse struct {
	Count      int32       `json:"count"`
	Categories []*Category `json:"categories"`
}

type Cp struct {
	Count      int32            `json:"count"`
	Categories []*ChildCategory `json:"categories"`
}

type ChildCategory struct {
	ParentId       string              `json:"id"`
	Name           string              `json:"name"`
	ChildsCategory []*CategoryByParent `json:"child"`
}

type CategoryByParent struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}