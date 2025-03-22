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

type ReturnRequestReasonController struct {
	ReturnRequestReasonUsecase domain.ReturnRequestReasonUsecase
	Env                        *bootstrap.Env
}

func (tc *ReturnRequestReasonController) Create(c *gin.Context) {
	var task domain.ReturnRequestReason
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

	err = tc.ReturnRequestReasonUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ReturnRequestReason created successfully",
	})
}

func (tc *ReturnRequestReasonController) Update(c *gin.Context) {
	var task domain.ReturnRequestReason
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

	err = tc.ReturnRequestReasonUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ReturnRequestReason update successfully",
	})
}

func (tc *ReturnRequestReasonController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ReturnRequestReasonUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ReturnRequestReasonController) FetchByID(c *gin.Context) {
	ReturnRequestReasonID := c.Query("id")
	if ReturnRequestReasonID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ReturnRequestReason, err := lc.ReturnRequestReasonUsecase.FetchByID(c, ReturnRequestReasonID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ReturnRequestReasonID})
		return
	}

	c.JSON(http.StatusOK, ReturnRequestReason)
}

func (lc *ReturnRequestReasonController) Fetch(c *gin.Context) {

	ReturnRequestReason, err := lc.ReturnRequestReasonUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ReturnRequestReason)
}
