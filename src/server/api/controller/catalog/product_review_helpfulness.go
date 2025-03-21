package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type ProductReviewHelpfulnessController struct {
	ProductReviewHelpfulnessUsecase domain.ProductReviewHelpfulnessUsecase
	Env                             *bootstrap.Env
}

func (tc *ProductReviewHelpfulnessController) Create(c *gin.Context) {
	var task domain.ProductReviewHelpfulness
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

	err = tc.ProductReviewHelpfulnessUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductReviewHelpfulness created successfully",
	})
}

func (tc *ProductReviewHelpfulnessController) Update(c *gin.Context) {
	var task domain.ProductReviewHelpfulness
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

	err = tc.ProductReviewHelpfulnessUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductReviewHelpfulness update successfully",
	})
}

func (tc *ProductReviewHelpfulnessController) Delete(c *gin.Context) {
	var task domain.ProductReviewHelpfulness
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

	err = tc.ProductReviewHelpfulnessUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductReviewHelpfulness update successfully",
	})
}

func (lc *ProductReviewHelpfulnessController) FetchByID(c *gin.Context) {
	ProductReviewHelpfulnessID := c.Query("id")
	if ProductReviewHelpfulnessID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductReviewHelpfulness, err := lc.ProductReviewHelpfulnessUsecase.FetchByID(c, ProductReviewHelpfulnessID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductReviewHelpfulnessID})
		return
	}

	c.JSON(http.StatusOK, ProductReviewHelpfulness)
}

func (lc *ProductReviewHelpfulnessController) Fetch(c *gin.Context) {

	ProductReviewHelpfulness, err := lc.ProductReviewHelpfulnessUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductReviewHelpfulness)
}
