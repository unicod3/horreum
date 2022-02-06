package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListArticle example
// @Tags articles
// @Summary Get all articles
// @Description Get all articles
// @ID list-articles
// @Accept  json
// @Produce  json
// @Success 200 {array} Article
// @Router /articles/ [get]
func (h *Handler) ListArticle(g *gin.Context) {
	articles, err := h.ArticleService.GetAll()
	if err != nil {
		return
	}
	g.JSON(http.StatusOK, articles)
}

// GetArticle example
// @Tags articles
// @Summary Get single article by id
// @Description Get single article by id
// @ID get-article
// @Accept  json
// @Produce  json
// @Param id path int true "Article ID"
// @Success 200 {object} Article
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /articles/{id} [get]
func (h *Handler) GetArticle(g *gin.Context) {
	var article Article

	if err := g.ShouldBindUri(&article); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	w, err := h.ArticleService.GetById(article.ID)
	if err != nil {
		g.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	}
	g.JSON(http.StatusOK, w)
}

// CreateArticle example
// @Tags articles
// @Summary Create a article with given data
// @Description Create a article with given data
// @ID create-article
// @Accept  json
// @Produce  json
// @Param article body ArticleRequestBody true "Article"
// @Success 200 {object} Article
// @Failure 400 {object} ErrorResponse
// @Router /articles/ [post]
func (h *Handler) CreateArticle(g *gin.Context) {
	var article Article

	if err := g.ShouldBindJSON(&article); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	err := h.ArticleService.Create(&article)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, article)
}

// UpdateArticle example
// @Tags articles
// @Summary Update a article with given data
// @Description Update a article with given data
// @ID update-article
// @Accept  json
// @Produce  json
// @Param id path int true "Article ID"
// @Param article body ArticleRequestBody true "Article"
// @Success 200 {object} Article
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /articles/{id} [put]
func (h *Handler) UpdateArticle(g *gin.Context) {
	var article Article

	if err := g.ShouldBindUri(&article); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	if err := g.ShouldBindJSON(&article); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the body",
		})
		return
	}

	err := h.ArticleService.Update(&article)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, article)
}

// DeleteArticle example
// @Tags articles
// @Summary Delete a article by id
// @Description Delete a article by id
// @ID delete-article
// @Accept  json
// @Produce  json
// @Param id path int true "Article ID"
// @Success 204 string string "NoContent"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /articles/{id} [delete]
func (h *Handler) DeleteArticle(g *gin.Context) {
	var article Article

	if err := g.ShouldBindUri(&article); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Couldn't resolve the params",
		})
		return
	}

	err := h.ArticleService.Delete(&article)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	g.Status(http.StatusNoContent)
}
