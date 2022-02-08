package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListProducts example
// @Tags products
// @Summary Get all products
// @Description Get all products
// @ID list-products
// @Accept  json
// @Produce  json
// @Success 200 {array} ProductArticle
// @Router /products/ [get]
func (service *ProductService) ListProducts(g *gin.Context) {
	products, err := service.GetAll()
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
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
func (service *ProductService) GetProduct(g *gin.Context) {
	var product Product

	if err := g.ShouldBindUri(&product); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	w, err := service.GetById(product.ID)
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
// @ID create-product
// @Accept  json
// @Produce  json
// @Param article body ProductRequestBody true "Product"
// @Success 200 {object} Article
// @Failure 400 {object} ErrorResponse
// @Router /products/ [post]
func (service *ProductService) CreateProduct(g *gin.Context) {
	var product Product

	if err := g.ShouldBindJSON(&product); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	p, err := service.Create(&product)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, p)
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
func (service *ProductService) UpdateProduct(g *gin.Context) {
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

	p, err := service.Update(&product)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, p)
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
func (service *ProductService) DeleteProduct(g *gin.Context) {
	var product Product

	if err := g.ShouldBindUri(&product); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	err := service.Delete(&product)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	g.Status(http.StatusNoContent)
}
