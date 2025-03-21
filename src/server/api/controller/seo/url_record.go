package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/seo"

	"github.com/gin-gonic/gin"
)

type UrlRecordController struct {
	UrlRecordUsecase domain.UrlRecordUsecase
	Env              *bootstrap.Env
}

func (tc *UrlRecordController) Create(c *gin.Context) {
	var task domain.UrlRecord
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

	err = tc.UrlRecordUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "UrlRecord created successfully",
	})
}

func (tc *UrlRecordController) Update(c *gin.Context) {
	var task domain.UrlRecord
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

	err = tc.UrlRecordUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "UrlRecord update successfully",
	})
}

func (tc *UrlRecordController) Delete(c *gin.Context) {
	var task domain.UrlRecord
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

	err = tc.UrlRecordUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "UrlRecord update successfully",
	})
}

func (lc *UrlRecordController) FetchByID(c *gin.Context) {
	UrlRecordID := c.Query("id")
	if UrlRecordID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	UrlRecord, err := lc.UrlRecordUsecase.FetchByID(c, UrlRecordID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: UrlRecordID})
		return
	}

	c.JSON(http.StatusOK, UrlRecord)
}

func (lc *UrlRecordController) Fetch(c *gin.Context) {

	UrlRecord, err := lc.UrlRecordUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, UrlRecord)
}
