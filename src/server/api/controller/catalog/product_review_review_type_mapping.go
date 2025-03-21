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

type ProductReviewReviewTypeMappingController struct {
	ProductReviewReviewTypeMappingUsecase domain.ProductReviewReviewTypeMappingUsecase
	Env                                   *bootstrap.Env
}

func (tc *ProductReviewReviewTypeMappingController) Create(c *gin.Context) {
	var task domain.ProductReviewReviewTypeMapping
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

	err = tc.ProductReviewReviewTypeMappingUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductReviewReviewTypeMapping created successfully",
	})
}

func (tc *ProductReviewReviewTypeMappingController) Update(c *gin.Context) {
	var task domain.ProductReviewReviewTypeMapping
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

	err = tc.ProductReviewReviewTypeMappingUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductReviewReviewTypeMapping update successfully",
	})
}

func (tc *ProductReviewReviewTypeMappingController) Delete(c *gin.Context) {
	var task domain.ProductReviewReviewTypeMapping
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

	err = tc.ProductReviewReviewTypeMappingUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductReviewReviewTypeMapping update successfully",
	})
}

func (lc *ProductReviewReviewTypeMappingController) FetchByID(c *gin.Context) {
	ProductReviewReviewTypeMappingID := c.Query("id")
	if ProductReviewReviewTypeMappingID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductReviewReviewTypeMapping, err := lc.ProductReviewReviewTypeMappingUsecase.FetchByID(c, ProductReviewReviewTypeMappingID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductReviewReviewTypeMappingID})
		return
	}

	c.JSON(http.StatusOK, ProductReviewReviewTypeMapping)
}

func (lc *ProductReviewReviewTypeMappingController) Fetch(c *gin.Context) {

	ProductReviewReviewTypeMapping, err := lc.ProductReviewReviewTypeMappingUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductReviewReviewTypeMapping)
}
