package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListOrders example
// @Tags orders
// @Summary Get all orders
// @Description Get all orders
// @ID list-orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /orders/ [get]
func (service *OrderService) ListOrders(g *gin.Context) {
	orders, err := service.GetAll()
	if err != nil {
		return
	}
	g.JSON(http.StatusOK, orders)
}

// GetOrder example
// @Tags orders
// @Summary Get single order by id
// @Description Get single order by id
// @ID get-order
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} Order
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /orders/{id} [get]
func (service *OrderService) GetOrder(g *gin.Context) {
	var order Order

	if err := g.ShouldBindUri(&order); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	w, err := service.GetById(order.ID)
	if err != nil {
		g.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	}
	g.JSON(http.StatusOK, w)
}

// CreateOrder example
// @Tags orders
// @Summary Create a order with given data
// @Description Create a order with given data
// @ID create-order
// @Accept  json
// @Produce  json
// @Param order body RequestBody true "Order"
// @Success 200 {object} Order
// @Failure 400 {object} ErrorResponse
// @Router /orders/ [post]
func (service *OrderService) CreateOrder(g *gin.Context) {
	var order Order

	if err := g.ShouldBindJSON(&order); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	err := service.Create(&order)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, order)
}

// UpdateOrder example
// @Tags orders
// @Summary Update a order with given data
// @Description Update a order with given data
// @ID update-order
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param order body RequestBody true "Order"
// @Success 200 {object} Order
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /orders/{id} [put]
func (service *OrderService) UpdateOrder(g *gin.Context) {
	var order Order

	if err := g.ShouldBindUri(&order); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	if err := g.ShouldBindJSON(&order); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	err := service.Update(&order)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, order)
}

// DeleteOrder example
// @Tags orders
// @Summary Delete a order by id
// @Description Delete a order by id
// @ID delete-order
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 204 string string "NoContent"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /orders/{id} [delete]
func (service *OrderService) DeleteOrder(g *gin.Context) {
	var order Order

	if err := g.ShouldBindUri(&order); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	err := service.Delete(&order)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	g.Status(http.StatusNoContent)
}
