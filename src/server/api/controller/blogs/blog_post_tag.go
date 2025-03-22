package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/blogs"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type BlogPostTagController struct {
	BlogPostTagUsecase domain.BlogPostTagUsecase
	Env                *bootstrap.Env
}

func (tc *BlogPostTagController) Create(c *gin.Context) {
	var task domain.BlogPostTag
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid request body"})
		return
	}

	err = tc.BlogPostTagUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BlogPostTag created successfully",
	})
}

func (tc *BlogPostTagController) Update(c *gin.Context) {
	var task domain.BlogPostTag
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid request body"})
		return
	}

	err = tc.BlogPostTagUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BlogPostTag update successfully",
	})
}

func (tc *BlogPostTagController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.BlogPostTagUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *BlogPostTagController) FetchByID(c *gin.Context) {
	BlogPostTagID := c.Query("id")
	if BlogPostTagID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	BlogPostTag, err := lc.BlogPostTagUsecase.FetchByID(c, BlogPostTagID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: BlogPostTagID})
		return
	}

	c.JSON(http.StatusOK, BlogPostTag)
}

func (lc *BlogPostTagController) Fetch(c *gin.Context) {

	BlogPostTag, err := lc.BlogPostTagUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, BlogPostTag)
}
