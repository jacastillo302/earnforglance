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

type RecurringPaymentController struct {
	RecurringPaymentUsecase domain.RecurringPaymentUsecase
	Env                     *bootstrap.Env
}

func (tc *RecurringPaymentController) Create(c *gin.Context) {
	var task domain.RecurringPayment
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

	err = tc.RecurringPaymentUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RecurringPayment created successfully",
	})
}

func (tc *RecurringPaymentController) Update(c *gin.Context) {
	var task domain.RecurringPayment
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

	err = tc.RecurringPaymentUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RecurringPayment update successfully",
	})
}

func (tc *RecurringPaymentController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.RecurringPaymentUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *RecurringPaymentController) FetchByID(c *gin.Context) {
	RecurringPaymentID := c.Query("id")
	if RecurringPaymentID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	RecurringPayment, err := lc.RecurringPaymentUsecase.FetchByID(c, RecurringPaymentID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: RecurringPaymentID})
		return
	}

	c.JSON(http.StatusOK, RecurringPayment)
}

func (lc *RecurringPaymentController) Fetch(c *gin.Context) {

	RecurringPayment, err := lc.RecurringPaymentUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RecurringPayment)
}
