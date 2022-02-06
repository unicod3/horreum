package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// List example
// @Tags orders
// @Summary Get all orders
// @Description Get all orders
// @ID list-orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /orders/ [get]
func (h *Handler) List(g *gin.Context) {
	orders, err := h.OrderService.GetAll()
	if err != nil {
		return
	}
	g.JSON(http.StatusOK, orders)
}

// Get example
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
func (h *Handler) Get(g *gin.Context) {
	var order Order

	if err := g.ShouldBindUri(&order); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	w, err := h.OrderService.GetById(order.ID)
	if err != nil {
		g.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	}
	g.JSON(http.StatusOK, w)
}

// Create example
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
func (h *Handler) Create(g *gin.Context) {
	var order Order

	if err := g.ShouldBindJSON(&order); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	err := h.OrderService.Create(&order)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, order)
}

// Update example
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
func (h *Handler) Update(g *gin.Context) {
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

	err := h.OrderService.Update(&order)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, order)
}

// Delete example
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
func (h *Handler) Delete(g *gin.Context) {
	var order Order

	if err := g.ShouldBindUri(&order); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	err := h.OrderService.Delete(&order)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	g.Status(http.StatusNoContent)
}
