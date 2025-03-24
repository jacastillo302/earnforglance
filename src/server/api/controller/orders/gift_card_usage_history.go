package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/orders"

	"github.com/gin-gonic/gin"
)

type GiftCardUsageHistoryController struct {
	GiftCardUsageHistoryUsecase domain.GiftCardUsageHistoryUsecase
	Env                         *bootstrap.Env
}

func (tc *GiftCardUsageHistoryController) CreateMany(c *gin.Context) {
	var task []domain.GiftCardUsageHistory
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

	err = tc.GiftCardUsageHistoryUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "GiftCardUsageHistory created successfully",
	})
}

func (tc *GiftCardUsageHistoryController) Create(c *gin.Context) {
	var task domain.GiftCardUsageHistory
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

	err = tc.GiftCardUsageHistoryUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "GiftCardUsageHistory created successfully",
	})
}

func (tc *GiftCardUsageHistoryController) Update(c *gin.Context) {
	var task domain.GiftCardUsageHistory
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

	err = tc.GiftCardUsageHistoryUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "GiftCardUsageHistory update successfully",
	})
}

func (tc *GiftCardUsageHistoryController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.GiftCardUsageHistoryUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *GiftCardUsageHistoryController) FetchByID(c *gin.Context) {
	GiftCardUsageHistoryID := c.Query("id")
	if GiftCardUsageHistoryID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	GiftCardUsageHistory, err := lc.GiftCardUsageHistoryUsecase.FetchByID(c, GiftCardUsageHistoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: GiftCardUsageHistoryID})
		return
	}

	c.JSON(http.StatusOK, GiftCardUsageHistory)
}

func (lc *GiftCardUsageHistoryController) Fetch(c *gin.Context) {

	GiftCardUsageHistory, err := lc.GiftCardUsageHistoryUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GiftCardUsageHistory)
}
