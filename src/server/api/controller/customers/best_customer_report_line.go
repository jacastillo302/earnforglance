package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/customers"

	"github.com/gin-gonic/gin"
)

type BestCustomerReportLineController struct {
	BestCustomerReportLineUsecase domain.BestCustomerReportLineUsecase
	Env                           *bootstrap.Env
}

func (tc *BestCustomerReportLineController) CreateMany(c *gin.Context) {
	var task []domain.BestCustomerReportLine
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

	err = tc.BestCustomerReportLineUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BestCustomerReportLine created successfully",
	})
}

func (tc *BestCustomerReportLineController) Create(c *gin.Context) {
	var task domain.BestCustomerReportLine
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

	err = tc.BestCustomerReportLineUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BestCustomerReportLine created successfully",
	})
}

func (tc *BestCustomerReportLineController) Update(c *gin.Context) {
	var task domain.BestCustomerReportLine
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

	err = tc.BestCustomerReportLineUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BestCustomerReportLine update successfully",
	})
}

func (tc *BestCustomerReportLineController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.BestCustomerReportLineUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *BestCustomerReportLineController) FetchByID(c *gin.Context) {
	BestCustomerReportLineID := c.Query("id")
	if BestCustomerReportLineID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	BestCustomerReportLine, err := lc.BestCustomerReportLineUsecase.FetchByID(c, BestCustomerReportLineID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: BestCustomerReportLineID})
		return
	}

	c.JSON(http.StatusOK, BestCustomerReportLine)
}

func (lc *BestCustomerReportLineController) Fetch(c *gin.Context) {

	BestCustomerReportLine, err := lc.BestCustomerReportLineUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, BestCustomerReportLine)
}
