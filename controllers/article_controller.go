package controllers

import (
	"net/http"
	"strconv"

	"sharingvision-be-test/dto"
	"sharingvision-be-test/services"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	service services.ArticleService
}

func NewArticleController(service services.ArticleService) *ArticleController {
	return &ArticleController{service: service}
}

func (c *ArticleController) CreateArticle(ctx *gin.Context) {
	var req dto.ArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid JSON payload: "+err.Error()))
		return
	}

	if err := c.service.CreateArticle(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.SuccessResponse("Article created successfully", map[string]interface{}{}))
}

func (c *ArticleController) GetArticles(ctx *gin.Context) {
	limitStr := ctx.Param("limit")
	if limitStr == "" {
		limitStr = ctx.Param("param1")
	}

	offsetStr := ctx.Param("offset")
	if offsetStr == "" {
		offsetStr = ctx.Param("param2")
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	articles, err := c.service.GetArticles(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to fetch articles: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.SuccessResponse("Articles fetched successfully", articles))
}

func (c *ArticleController) GetArticleByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		idStr = ctx.Param("param1")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid article ID"))
		return
	}

	article, err := c.service.GetArticleByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.SuccessResponse("Article fetched successfully", article))
}

func (c *ArticleController) UpdateArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		idStr = ctx.Param("param1")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid article ID"))
		return
	}

	var req dto.ArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid JSON payload: "+err.Error()))
		return
	}

	if err := c.service.UpdateArticle(id, &req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.SuccessResponse("Article updated successfully", map[string]interface{}{}))
}

func (c *ArticleController) DeleteArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		idStr = ctx.Param("param1")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid article ID"))
		return
	}

	if err := c.service.DeleteArticle(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to delete article: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.SuccessResponse("Article deleted successfully", map[string]interface{}{}))
}
