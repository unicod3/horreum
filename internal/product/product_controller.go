package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListProduct example
// @Tags products
// @Summary Get all products
// @Description Get all products
// @ID list-products
// @Accept  json
// @Produce  json
// @Success 200 {array} Product
// @Router /products/ [get]
func (h *Handler) ListProduct(g *gin.Context) {
	products, err := h.ProductService.GetAll()
	if err != nil {
		return
	}
	g.JSON(http.StatusOK, products)
}

// GetProduct example
// @Tags products
// @Summary Get single product by id
// @Description Get single product by id
// @ID get-product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /products/{id} [get]
func (h *Handler) GetProduct(g *gin.Context) {
	var product Product

	if err := g.ShouldBindUri(&product); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	w, err := h.ProductService.GetById(product.ID)
	if err != nil {
		g.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	}
	g.JSON(http.StatusOK, w)
}

// CreateProduct example
// @Tags products
// @Summary Create a article with given data
// @Description Create a article with given data
// @ID create-article
// @Accept  json
// @Produce  json
// @Param article body ProductRequestBody true "Product"
// @Success 200 {object} Article
// @Failure 400 {object} ErrorResponse
// @Router /products/ [post]
func (h *Handler) CreateProduct(g *gin.Context) {
	var product Product

	if err := g.ShouldBindJSON(&product); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	err := h.ProductService.Create(&product)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, product)
}

// UpdateProduct example
// @Tags products
// @Summary Update a product with given data
// @Description Update a product with given data
// @ID update-product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param article body ProductRequestBody true "Product"
// @Success 200 {object} Product
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /products/{id} [put]
func (h *Handler) UpdateProduct(g *gin.Context) {
	var product Product

	if err := g.ShouldBindUri(&product); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	if err := g.ShouldBindJSON(&product); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	err := h.ProductService.Update(&product)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, product)
}

// DeleteProduct example
// @Tags products
// @Summary Delete a product by id
// @Description Delete a product by id
// @ID delete-product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 204 string string "NoContent"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /products/{id} [delete]
func (h *Handler) DeleteProduct(g *gin.Context) {
	var product Product

	if err := g.ShouldBindUri(&product); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	err := h.ProductService.Delete(&product)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	g.Status(http.StatusNoContent)
}
