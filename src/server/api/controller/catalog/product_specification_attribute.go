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

type ProductSpecificationAttributeController struct {
	ProductSpecificationAttributeUsecase domain.ProductSpecificationAttributeUsecase
	Env                                  *bootstrap.Env
}

func (tc *ProductSpecificationAttributeController) Create(c *gin.Context) {
	var task domain.ProductSpecificationAttribute
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

	err = tc.ProductSpecificationAttributeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductSpecificationAttribute created successfully",
	})
}

func (tc *ProductSpecificationAttributeController) Update(c *gin.Context) {
	var task domain.ProductSpecificationAttribute
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

	err = tc.ProductSpecificationAttributeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductSpecificationAttribute update successfully",
	})
}

func (tc *ProductSpecificationAttributeController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ProductSpecificationAttributeUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ProductSpecificationAttributeController) FetchByID(c *gin.Context) {
	ProductSpecificationAttributeID := c.Query("id")
	if ProductSpecificationAttributeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductSpecificationAttribute, err := lc.ProductSpecificationAttributeUsecase.FetchByID(c, ProductSpecificationAttributeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductSpecificationAttributeID})
		return
	}

	c.JSON(http.StatusOK, ProductSpecificationAttribute)
}

func (lc *ProductSpecificationAttributeController) Fetch(c *gin.Context) {

	ProductSpecificationAttribute, err := lc.ProductSpecificationAttributeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductSpecificationAttribute)
}
