package warehouse

import (
	"github.com/gin-gonic/gin"
	intWarehouse "github.com/unicod3/horreum/internal/warehouse"
	"net/http"
)

// List example
// @Tags warehouses
// @Summary Get all warehouses
// @Description Get all warehouses
// @ID list-warehouse
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Router /warehouses/ [get]
func List(g *gin.Context) {
	w := intWarehouse.GetDummyWarehouses()
	g.JSON(http.StatusOK, SuccessResponse{w})
}

// Get example
// @Tags warehouses
// @Summary Get single warehouse by id
// @Description Get single warehouse by id
// @ID get-warehouse
// @Accept  json
// @Produce  json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /warehouses/{id} [get]
func Get(g *gin.Context) {
	warehouse := intWarehouse.Warehouse{}

	if err := g.ShouldBindUri(&warehouse); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	w := intWarehouse.GetWarehouse(warehouse.ID)
	g.JSON(http.StatusOK, SuccessResponse{w})
}

// Create example
// @Tags warehouses
// @Summary Create a warehouse with given data
// @Description Create a warehouse with given data
// @ID create-warehouse
// @Accept  json
// @Produce  json
// @Param warehouse body RequestBody true "Warehouse"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /warehouses/ [post]
func Create(g *gin.Context) {
	warehouse := intWarehouse.Warehouse{}

	if err := g.ShouldBindJSON(&warehouse); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	w := intWarehouse.CreateWarehouse(warehouse)
	g.JSON(http.StatusOK, SuccessResponse{w})
}

// Update example
// @Tags warehouses
// @Summary Update a warehouse with given data
// @Description Update a warehouse with given data
// @ID update-warehouse
// @Accept  json
// @Produce  json
// @Param id path int true "Warehouse ID"
// @Param warehouse body RequestBody true "Warehouse"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /warehouses/{id} [put]
func Update(g *gin.Context) {
	warehouse := intWarehouse.Warehouse{}

	if err := g.ShouldBindUri(&warehouse); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	if err := g.ShouldBindJSON(&warehouse); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	w := intWarehouse.UpdateWarehouse(warehouse)
	g.JSON(http.StatusOK, SuccessResponse{w})
}

// Delete example
// @Tags warehouses
// @Summary Delete a warehouse by id
// @Description Delete a warehouse by id
// @ID delete-warehouse
// @Accept  json
// @Produce  json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /warehouses/{id} [delete]
func Delete(g *gin.Context) {
	warehouse := intWarehouse.Warehouse{}

	if err := g.ShouldBindUri(&warehouse); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	w := intWarehouse.DeleteWarehouse(warehouse)
	g.JSON(http.StatusOK, SuccessResponse{w})
}
