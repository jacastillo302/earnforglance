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

type SalesSummaryReportLineController struct {
	SalesSummaryReportLineUsecase domain.SalesSummaryReportLineUsecase
	Env                           *bootstrap.Env
}

func (tc *SalesSummaryReportLineController) Create(c *gin.Context) {
	var task domain.SalesSummaryReportLine
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

	err = tc.SalesSummaryReportLineUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SalesSummaryReportLine created successfully",
	})
}

func (tc *SalesSummaryReportLineController) Update(c *gin.Context) {
	var task domain.SalesSummaryReportLine
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

	err = tc.SalesSummaryReportLineUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SalesSummaryReportLine update successfully",
	})
}

func (tc *SalesSummaryReportLineController) Delete(c *gin.Context) {
	var task domain.SalesSummaryReportLine
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

	err = tc.SalesSummaryReportLineUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SalesSummaryReportLine update successfully",
	})
}

func (lc *SalesSummaryReportLineController) FetchByID(c *gin.Context) {
	SalesSummaryReportLineID := c.Query("id")
	if SalesSummaryReportLineID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	SalesSummaryReportLine, err := lc.SalesSummaryReportLineUsecase.FetchByID(c, SalesSummaryReportLineID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: SalesSummaryReportLineID})
		return
	}

	c.JSON(http.StatusOK, SalesSummaryReportLine)
}

func (lc *SalesSummaryReportLineController) Fetch(c *gin.Context) {

	SalesSummaryReportLine, err := lc.SalesSummaryReportLineUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SalesSummaryReportLine)
}
