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

type WarehouseController struct {
	WarehouseUsecase domain.WarehouseUsecase
	Env              *bootstrap.Env
}

func (tc *WarehouseController) Create(c *gin.Context) {
	var task domain.Warehouse
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

	err = tc.WarehouseUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Warehouse created successfully",
	})
}

func (tc *WarehouseController) Update(c *gin.Context) {
	var task domain.Warehouse
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

	err = tc.WarehouseUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Warehouse update successfully",
	})
}

func (tc *WarehouseController) Delete(c *gin.Context) {
	var task domain.Warehouse
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

	err = tc.WarehouseUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Warehouse update successfully",
	})
}

func (lc *WarehouseController) FetchByID(c *gin.Context) {
	WarehouseID := c.Query("id")
	if WarehouseID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	Warehouse, err := lc.WarehouseUsecase.FetchByID(c, WarehouseID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: WarehouseID})
		return
	}

	c.JSON(http.StatusOK, Warehouse)
}

func (lc *WarehouseController) Fetch(c *gin.Context) {

	Warehouse, err := lc.WarehouseUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Warehouse)
}
