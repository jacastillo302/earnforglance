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

type ReturnRequestController struct {
	ReturnRequestUsecase domain.ReturnRequestUsecase
	Env                  *bootstrap.Env
}

func (tc *ReturnRequestController) Create(c *gin.Context) {
	var task domain.ReturnRequest
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

	err = tc.ReturnRequestUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ReturnRequest created successfully",
	})
}

func (tc *ReturnRequestController) Update(c *gin.Context) {
	var task domain.ReturnRequest
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

	err = tc.ReturnRequestUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ReturnRequest update successfully",
	})
}

func (tc *ReturnRequestController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ReturnRequestUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ReturnRequestController) FetchByID(c *gin.Context) {
	ReturnRequestID := c.Query("id")
	if ReturnRequestID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ReturnRequest, err := lc.ReturnRequestUsecase.FetchByID(c, ReturnRequestID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ReturnRequestID})
		return
	}

	c.JSON(http.StatusOK, ReturnRequest)
}

func (lc *ReturnRequestController) Fetch(c *gin.Context) {

	ReturnRequest, err := lc.ReturnRequestUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ReturnRequest)
}
