package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type SearchTermReportLineController struct {
	SearchTermReportLineUsecase common.SearchTermReportLineUsecase
	Env                         *bootstrap.Env
}

func (tc *SearchTermReportLineController) CreateMany(c *gin.Context) {
	var task []common.SearchTermReportLine
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

	err = tc.SearchTermReportLineUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SearchTermReportLine created successfully",
	})
}

func (tc *SearchTermReportLineController) Create(c *gin.Context) {
	var task common.SearchTermReportLine
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

	err = tc.SearchTermReportLineUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SearchTermReportLine created successfully",
	})
}

func (tc *SearchTermReportLineController) Update(c *gin.Context) {
	var task common.SearchTermReportLine
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

	err = tc.SearchTermReportLineUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SearchTermReportLine update successfully",
	})
}

func (tc *SearchTermReportLineController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.SearchTermReportLineUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *SearchTermReportLineController) FetchByID(c *gin.Context) {
	SearchTermReportLineID := c.Query("id")
	if SearchTermReportLineID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	SearchTermReportLine, err := lc.SearchTermReportLineUsecase.FetchByID(c, SearchTermReportLineID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: SearchTermReportLineID})
		return
	}

	c.JSON(http.StatusOK, SearchTermReportLine)
}

func (lc *SearchTermReportLineController) Fetch(c *gin.Context) {

	SearchTermReportLine, err := lc.SearchTermReportLineUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SearchTermReportLine)
}
