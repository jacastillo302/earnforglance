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

type ProductCategoryController struct {
	ProductCategoryUsecase domain.ProductCategoryUsecase
	Env                    *bootstrap.Env
}

func (tc *ProductCategoryController) Create(c *gin.Context) {
	var task domain.ProductCategory
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

	err = tc.ProductCategoryUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductCategory created successfully",
	})
}

func (tc *ProductCategoryController) Update(c *gin.Context) {
	var task domain.ProductCategory
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

	err = tc.ProductCategoryUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductCategory update successfully",
	})
}

func (tc *ProductCategoryController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ProductCategoryUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ProductCategoryController) FetchByID(c *gin.Context) {
	ProductCategoryID := c.Query("id")
	if ProductCategoryID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductCategory, err := lc.ProductCategoryUsecase.FetchByID(c, ProductCategoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductCategoryID})
		return
	}

	c.JSON(http.StatusOK, ProductCategory)
}

func (lc *ProductCategoryController) Fetch(c *gin.Context) {

	ProductCategory, err := lc.ProductCategoryUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductCategory)
}
