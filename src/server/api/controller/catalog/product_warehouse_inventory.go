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

type ProductWarehouseInventoryController struct {
	ProductWarehouseInventoryUsecase domain.ProductWarehouseInventoryUsecase
	Env                              *bootstrap.Env
}

func (tc *ProductWarehouseInventoryController) Create(c *gin.Context) {
	var task domain.ProductWarehouseInventory
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

	err = tc.ProductWarehouseInventoryUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductWarehouseInventory created successfully",
	})
}

func (tc *ProductWarehouseInventoryController) Update(c *gin.Context) {
	var task domain.ProductWarehouseInventory
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

	err = tc.ProductWarehouseInventoryUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductWarehouseInventory update successfully",
	})
}

func (tc *ProductWarehouseInventoryController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ProductWarehouseInventoryUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ProductWarehouseInventoryController) FetchByID(c *gin.Context) {
	ProductWarehouseInventoryID := c.Query("id")
	if ProductWarehouseInventoryID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductWarehouseInventory, err := lc.ProductWarehouseInventoryUsecase.FetchByID(c, ProductWarehouseInventoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductWarehouseInventoryID})
		return
	}

	c.JSON(http.StatusOK, ProductWarehouseInventory)
}

func (lc *ProductWarehouseInventoryController) Fetch(c *gin.Context) {

	ProductWarehouseInventory, err := lc.ProductWarehouseInventoryUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductWarehouseInventory)
}
