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

type RecurringPaymentHistoryController struct {
	RecurringPaymentHistoryUsecase domain.RecurringPaymentHistoryUsecase
	Env                            *bootstrap.Env
}

func (tc *RecurringPaymentHistoryController) Create(c *gin.Context) {
	var task domain.RecurringPaymentHistory
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

	err = tc.RecurringPaymentHistoryUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RecurringPaymentHistory created successfully",
	})
}

func (tc *RecurringPaymentHistoryController) Update(c *gin.Context) {
	var task domain.RecurringPaymentHistory
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

	err = tc.RecurringPaymentHistoryUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RecurringPaymentHistory update successfully",
	})
}

func (tc *RecurringPaymentHistoryController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.RecurringPaymentHistoryUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *RecurringPaymentHistoryController) FetchByID(c *gin.Context) {
	RecurringPaymentHistoryID := c.Query("id")
	if RecurringPaymentHistoryID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	RecurringPaymentHistory, err := lc.RecurringPaymentHistoryUsecase.FetchByID(c, RecurringPaymentHistoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: RecurringPaymentHistoryID})
		return
	}

	c.JSON(http.StatusOK, RecurringPaymentHistory)
}

func (lc *RecurringPaymentHistoryController) Fetch(c *gin.Context) {

	RecurringPaymentHistory, err := lc.RecurringPaymentHistoryUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RecurringPaymentHistory)
}
