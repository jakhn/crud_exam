package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"crud/models"

)

// CreateOrders godoc
// @ID create_orders
// @Router /orders [POST]
// @Summary Create Orders
// @Description Create Orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param orders body models.CreateOrders true "CreateOrdersRequestBody"
// @Success 201 {object} models.Orders "GetOrdersBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) CreateOrders(c *gin.Context) {
	var orders models.CreateOrders

	err := c.ShouldBindJSON(&orders)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Orders().Create(context.Background(), &orders)
	if err != nil {
		log.Printf("error whiling Create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling Create").Error())
		return
	}

	resp, err := h.storage.Orders().GetByPKey(
		context.Background(),
		&models.OrdersPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIdOrders godoc
// @ID get_by_id_orders
// @Router /orders/{id} [GET]
// @Summary Get By Id Orders
// @Description Get By Id Orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Orders "GetOrdersBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetOrdersById(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.Orders().GetByPKey(
		context.Background(),
		&models.OrdersPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetListOrders godoc
// @ID get_list_orders
// @Router /orders [GET]
// @Summary Get List Orders
// @Description Get List Orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.Cp "GetOrdersBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetOrdersList(c *gin.Context) {
	var (
		limit  int
		offset int
		err    error
	)

	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Printf("error whiling limit: %v\n", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Printf("error whiling limit: %v\n", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	resp, err := h.storage.Orders().GetList(
		context.Background(),
		&models.GetListOrdersRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateOrders godoc
// @ID update_orders
// @Router /orders/{id} [PUT]
// @Summary Update Orders
// @Description Update Orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param orders body models.CreateOrders true "CreateOrdersRequestBody"
// @Success 200 {object} models.Orders "GetOrdersBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) UpdateOrders(c *gin.Context) {

	var (
		orders models.UpdateOrders
	)

	orders.Id = c.Param("id")

	if orders.Id == "" {
		log.Printf("error whiling update: %v\n", errors.New("required orders id").Error())
		c.JSON(http.StatusBadRequest, errors.New("required orders id").Error())
		return
	}

	err := c.ShouldBindJSON(&orders)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storage.Orders().Update(
		context.Background(),
		&orders,
	)

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	if rowsAffected == 0 {
		log.Printf("error whiling update rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update rows affected").Error())
		return
	}

	resp, err := h.storage.Orders().GetByPKey(
		context.Background(),
		&models.OrdersPrimarKey{Id: orders.Id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteByOrders godoc
// @ID delete_by_id_orders
// @Router /orders/{id} [DELETE]
// @Summary Delete By Id Orders
// @Description Delete By Id Orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Orders "GetOrdersBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) DeleteOrders(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		log.Printf("error whiling update: %v\n", errors.New("required orders id").Error())
		c.JSON(http.StatusBadRequest, errors.New("required orders id").Error())
		return
	}

	err := h.storage.Orders().Delete(
		context.Background(),
		&models.OrdersPrimarKey{
			Id: id,
		},
	)

	if err != nil {
		log.Printf("error whiling delete: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete").Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
