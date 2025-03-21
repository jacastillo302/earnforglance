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

type ProductAttributeValueController struct {
	ProductAttributeValueUsecase domain.ProductAttributeValueUsecase
	Env                          *bootstrap.Env
}

func (tc *ProductAttributeValueController) Create(c *gin.Context) {
	var task domain.ProductAttributeValue
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

	err = tc.ProductAttributeValueUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAttributeValue created successfully",
	})
}

func (tc *ProductAttributeValueController) Update(c *gin.Context) {
	var task domain.ProductAttributeValue
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

	err = tc.ProductAttributeValueUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAttributeValue update successfully",
	})
}

func (tc *ProductAttributeValueController) Delete(c *gin.Context) {
	var task domain.ProductAttributeValue
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

	err = tc.ProductAttributeValueUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAttributeValue update successfully",
	})
}

func (lc *ProductAttributeValueController) FetchByID(c *gin.Context) {
	ProductAttributeValueID := c.Query("id")
	if ProductAttributeValueID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductAttributeValue, err := lc.ProductAttributeValueUsecase.FetchByID(c, ProductAttributeValueID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductAttributeValueID})
		return
	}

	c.JSON(http.StatusOK, ProductAttributeValue)
}

func (lc *ProductAttributeValueController) Fetch(c *gin.Context) {

	ProductAttributeValue, err := lc.ProductAttributeValueUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductAttributeValue)
}
