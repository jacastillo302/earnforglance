package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/discounts"

	"github.com/gin-gonic/gin"
)

type DiscountUsageHistoryController struct {
	DiscountUsageHistoryUsecase domain.DiscountUsageHistoryUsecase
	Env                         *bootstrap.Env
}

func (tc *DiscountUsageHistoryController) CreateMany(c *gin.Context) {
	var task []domain.DiscountUsageHistory
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

	err = tc.DiscountUsageHistoryUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountUsageHistory created successfully",
	})
}

func (tc *DiscountUsageHistoryController) Create(c *gin.Context) {
	var task domain.DiscountUsageHistory
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

	err = tc.DiscountUsageHistoryUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountUsageHistory created successfully",
	})
}

func (tc *DiscountUsageHistoryController) Update(c *gin.Context) {
	var task domain.DiscountUsageHistory
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

	err = tc.DiscountUsageHistoryUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountUsageHistory update successfully",
	})
}

func (tc *DiscountUsageHistoryController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.DiscountUsageHistoryUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *DiscountUsageHistoryController) FetchByID(c *gin.Context) {
	DiscountUsageHistoryID := c.Query("id")
	if DiscountUsageHistoryID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	DiscountUsageHistory, err := lc.DiscountUsageHistoryUsecase.FetchByID(c, DiscountUsageHistoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: DiscountUsageHistoryID})
		return
	}

	c.JSON(http.StatusOK, DiscountUsageHistory)
}

func (lc *DiscountUsageHistoryController) Fetch(c *gin.Context) {

	DiscountUsageHistory, err := lc.DiscountUsageHistoryUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DiscountUsageHistory)
}
