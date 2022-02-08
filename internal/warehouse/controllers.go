package warehouse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListWarehouses example
// @Tags warehouses
// @Summary Get all warehouses
// @Description Get all warehouses
// @ID list-warehouse
// @Accept  json
// @Produce  json
// @Success 200 {array} Warehouse
// @Router /warehouses/ [get]
func (service *WarehouseService) ListWarehouses(g *gin.Context) {
	warehouses, err := service.GetAll()
	if err != nil {
		return
	}
	g.JSON(http.StatusOK, warehouses)
}

// GetWarehouse example
// @Tags warehouses
// @Summary Get single warehouse by id
// @Description Get single warehouse by id
// @ID get-warehouse
// @Accept  json
// @Produce  json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} Warehouse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /warehouses/{id} [get]
func (service *WarehouseService) GetWarehouse(g *gin.Context) {
	warehouse := Warehouse{}

	if err := g.ShouldBindUri(&warehouse); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	w, err := service.GetById(warehouse.ID)
	if err != nil {
		g.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	}
	g.JSON(http.StatusOK, w)
}

// CreateWarehouse example
// @Tags warehouses
// @Summary Create a warehouse with given data
// @Description Create a warehouse with given data
// @ID create-warehouse
// @Accept  json
// @Produce  json
// @Param warehouse body RequestBody true "Warehouse"
// @Success 200 {object} Warehouse
// @Failure 400 {object} ErrorResponse
// @Router /warehouses/ [post]
func (service *WarehouseService) CreateWarehouse(g *gin.Context) {
	var warehouse Warehouse
	if err := g.ShouldBindJSON(&warehouse); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	err := service.Create(&warehouse)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, warehouse)
}

// UpdateWarehouse example
// @Tags warehouses
// @Summary Update a warehouse with given data
// @Description Update a warehouse with given data
// @ID update-warehouse
// @Accept  json
// @Produce  json
// @Param id path int true "Warehouse ID"
// @Param warehouse body RequestBody true "Warehouse"
// @Success 200 {object} Warehouse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /warehouses/{id} [put]
func (service *WarehouseService) UpdateWarehouse(g *gin.Context) {
	var warehouse Warehouse

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

	err := service.Update(&warehouse)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, warehouse)
}

// DeleteWarehouse example
// @Tags warehouses
// @Summary Delete a warehouse by id
// @Description Delete a warehouse by id
// @ID delete-warehouse
// @Accept  json
// @Produce  json
// @Param id path int true "Warehouse ID"
// @Success 204 string string "NoContent"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /warehouses/{id} [delete]
func (service *WarehouseService) DeleteWarehouse(g *gin.Context) {
	warehouse := Warehouse{}

	if err := g.ShouldBindUri(&warehouse); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	err := service.Delete(&warehouse)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	g.Status(http.StatusNoContent)
}
