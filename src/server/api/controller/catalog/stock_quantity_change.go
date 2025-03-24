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

type StockQuantityChangeController struct {
	StockQuantityChangeUsecase domain.StockQuantityChangeUsecase
	Env                        *bootstrap.Env
}

func (tc *StockQuantityChangeController) CreateMany(c *gin.Context) {
	var task []domain.StockQuantityChange
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

	err = tc.StockQuantityChangeUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "StockQuantityChange created successfully",
	})
}

func (tc *StockQuantityChangeController) Create(c *gin.Context) {
	var task domain.StockQuantityChange
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

	err = tc.StockQuantityChangeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "StockQuantityChange created successfully",
	})
}

func (tc *StockQuantityChangeController) Update(c *gin.Context) {
	var task domain.StockQuantityChange
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

	err = tc.StockQuantityChangeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "StockQuantityChange update successfully",
	})
}

func (tc *StockQuantityChangeController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.StockQuantityChangeUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *StockQuantityChangeController) FetchByID(c *gin.Context) {
	StockQuantityChangeID := c.Query("id")
	if StockQuantityChangeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	StockQuantityChange, err := lc.StockQuantityChangeUsecase.FetchByID(c, StockQuantityChangeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: StockQuantityChangeID})
		return
	}

	c.JSON(http.StatusOK, StockQuantityChange)
}

func (lc *StockQuantityChangeController) Fetch(c *gin.Context) {

	StockQuantityChange, err := lc.StockQuantityChangeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, StockQuantityChange)
}
