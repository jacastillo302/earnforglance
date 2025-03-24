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

type BestSellersReportLineController struct {
	BestSellersReportLineUsecase domain.BestSellersReportLineUsecase
	Env                          *bootstrap.Env
}

func (tc *BestSellersReportLineController) CreateMany(c *gin.Context) {
	var task []domain.BestSellersReportLine
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

	err = tc.BestSellersReportLineUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BestSellersReportLine created successfully",
	})
}

func (tc *BestSellersReportLineController) Create(c *gin.Context) {
	var task domain.BestSellersReportLine
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

	err = tc.BestSellersReportLineUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BestSellersReportLine created successfully",
	})
}

func (tc *BestSellersReportLineController) Update(c *gin.Context) {
	var task domain.BestSellersReportLine
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

	err = tc.BestSellersReportLineUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BestSellersReportLine update successfully",
	})
}

func (tc *BestSellersReportLineController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.BestSellersReportLineUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *BestSellersReportLineController) FetchByID(c *gin.Context) {
	BestSellersReportLineID := c.Query("id")
	if BestSellersReportLineID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	BestSellersReportLine, err := lc.BestSellersReportLineUsecase.FetchByID(c, BestSellersReportLineID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: BestSellersReportLineID})
		return
	}

	c.JSON(http.StatusOK, BestSellersReportLine)
}

func (lc *BestSellersReportLineController) Fetch(c *gin.Context) {

	BestSellersReportLine, err := lc.BestSellersReportLineUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, BestSellersReportLine)
}
