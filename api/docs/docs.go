// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/category": {
            "get": {
                "description": "Get List Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get List Category",
                "operationId": "get_list_category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetCategoryBody",
                        "schema": {
                            "$ref": "#/definitions/models.GetListCategoryResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create Category",
                "operationId": "create_category",
                "parameters": [
                    {
                        "description": "CreateCategoryRequestBody",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateCategory"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "GetCategoryBody",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/category/{id}": {
            "get": {
                "description": "Get By Id Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get By Id Category",
                "operationId": "get_by_id_category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetCategoryBody",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Update Category",
                "operationId": "update_category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "CreateCategoryRequestBody",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateCategorySwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetCategorysBody",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete By Id Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Delete By Id Category",
                "operationId": "delete_by_id_category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetCategoryBody",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order": {
            "get": {
                "description": "Get List Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get List Order",
                "operationId": "get_list_order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetOrderBody",
                        "schema": {
                            "$ref": "#/definitions/models.GetListOrderResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Create Order",
                "operationId": "create_order",
                "parameters": [
                    {
                        "description": "CreateOrderRequestBody",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateOrder"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "GetorderBody",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order/{id}": {
            "get": {
                "description": "Get By Id Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get By Id Order",
                "operationId": "get_by_id_order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetOrderBody",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Update Order",
                "operationId": "update_order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "CreateOrderRequestBody",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateOrderSwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetordersBody",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete By Id Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Delete By Id Order",
                "operationId": "delete_by_id_order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetOrderBody",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "description": "Get List Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get List Product",
                "operationId": "get_list_product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "category_id",
                        "name": "category_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetProductBody",
                        "schema": {
                            "$ref": "#/definitions/models.GetListProductResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create Product",
                "operationId": "create_product",
                "parameters": [
                    {
                        "description": "CreateProductRequestBody",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateProduct"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "GetProductBody",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "description": "Get By Id Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get By Id Product",
                "operationId": "get_by_id_product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetProductBody",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update Product",
                "operationId": "update_product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "CreateProductRequestBody",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProductSwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetProductsBody",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete By Id Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete By Id Product",
                "operationId": "delete_by_id_product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetProductBody",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Categories": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "order_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OrderList"
                    }
                },
                "parent_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parent_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CategoryList": {
            "type": "object",
            "properties": {
                "childs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Category"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parent_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CreateCategory": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "parent_id": {
                    "type": "string"
                }
            }
        },
        "models.CreateOrder": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "models.CreateProduct": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "models.GetListCategoryResponse": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CategoryList"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "models.GetListOrderResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OrderList"
                    }
                }
            }
        },
        "models.GetListProductResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Product"
                    }
                }
            }
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.OrderList": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "product": {
                    "$ref": "#/definitions/models.ProductList"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.ProductList": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/models.Categories"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.UpdateCategorySwagger": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "parent_id": {
                    "type": "string"
                }
            }
        },
        "models.UpdateOrderSwagger": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "models.UpdateProductSwagger": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
