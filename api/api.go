package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "crud/api/docs"
	"crud/api/handler"
	"crud/storage"
)

func SetUpApi(r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandlerV1(storage)

	r.Use(customCORSMiddleware())

	v1 := r.Group("/v1")

	v1.Use(checkPassword())
	

	r.POST("/product", handlerV1.CreateProduct)
	r.GET("/product/:id", handlerV1.GetProductById)
	r.GET("/product", handlerV1.GetProductList)
	r.GET("/sortby", handlerV1.GetCategoryList)
	r.PUT("/product/:id", handlerV1.UpdateProduct)
	r.DELETE("/product/:id", handlerV1.DeleteProduct)

	r.POST("/category", handlerV1.CreateCategory)
	r.GET("/category/:id", handlerV1.GetCategoryById)
	r.GET("/category", handlerV1.GetCategoryList)
	r.PUT("/category/:id", handlerV1.UpdateCategory)
	r.DELETE("/category/:id", handlerV1.DeleteCategory)

	r.POST("/order", handlerV1.CreateCategory)
	r.GET("/order/:id", handlerV1.GetCategoryById)
	r.GET("/order", handlerV1.GetCategoryList)
	r.PUT("/order/:id", handlerV1.UpdateCategory)
	r.DELETE("/order/:id", handlerV1.DeleteCategory)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func checkPassword() gin.HandlerFunc {

	return func(c *gin.Context) {

		if _, ok := c.Request.Header["Password"]; ok {
			if c.Request.Header["Password"][0] != "1234" {
				c.AbortWithError(http.StatusForbidden, errors.New("not found password"))
				return
			}
		} else {
			c.AbortWithError(http.StatusForbidden, errors.New("not found password"))
			return
		}

		c.Next()
	}
}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
