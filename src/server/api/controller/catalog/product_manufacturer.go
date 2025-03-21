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

type ProductManufacturerController struct {
	ProductManufacturerUsecase domain.ProductManufacturerUsecase
	Env                        *bootstrap.Env
}

func (tc *ProductManufacturerController) Create(c *gin.Context) {
	var task domain.ProductManufacturer
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

	err = tc.ProductManufacturerUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductManufacturer created successfully",
	})
}

func (tc *ProductManufacturerController) Update(c *gin.Context) {
	var task domain.ProductManufacturer
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

	err = tc.ProductManufacturerUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductManufacturer update successfully",
	})
}

func (tc *ProductManufacturerController) Delete(c *gin.Context) {
	var task domain.ProductManufacturer
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

	err = tc.ProductManufacturerUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductManufacturer update successfully",
	})
}

func (lc *ProductManufacturerController) FetchByID(c *gin.Context) {
	ProductManufacturerID := c.Query("id")
	if ProductManufacturerID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductManufacturer, err := lc.ProductManufacturerUsecase.FetchByID(c, ProductManufacturerID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductManufacturerID})
		return
	}

	c.JSON(http.StatusOK, ProductManufacturer)
}

func (lc *ProductManufacturerController) Fetch(c *gin.Context) {

	ProductManufacturer, err := lc.ProductManufacturerUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductManufacturer)
}
