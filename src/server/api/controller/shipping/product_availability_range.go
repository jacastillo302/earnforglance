package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/shipping"

	"github.com/gin-gonic/gin"
)

type ProductAvailabilityRangeController struct {
	ProductAvailabilityRangeUsecase domain.ProductAvailabilityRangeUsecase
	Env                             *bootstrap.Env
}

func (tc *ProductAvailabilityRangeController) Create(c *gin.Context) {
	var task domain.ProductAvailabilityRange
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

	err = tc.ProductAvailabilityRangeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAvailabilityRange created successfully",
	})
}

func (tc *ProductAvailabilityRangeController) Update(c *gin.Context) {
	var task domain.ProductAvailabilityRange
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

	err = tc.ProductAvailabilityRangeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAvailabilityRange update successfully",
	})
}

func (tc *ProductAvailabilityRangeController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ProductAvailabilityRangeUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ProductAvailabilityRangeController) FetchByID(c *gin.Context) {
	ProductAvailabilityRangeID := c.Query("id")
	if ProductAvailabilityRangeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductAvailabilityRange, err := lc.ProductAvailabilityRangeUsecase.FetchByID(c, ProductAvailabilityRangeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductAvailabilityRangeID})
		return
	}

	c.JSON(http.StatusOK, ProductAvailabilityRange)
}

func (lc *ProductAvailabilityRangeController) Fetch(c *gin.Context) {

	ProductAvailabilityRange, err := lc.ProductAvailabilityRangeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductAvailabilityRange)
}
