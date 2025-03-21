package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/security"

	"github.com/gin-gonic/gin"
)

type PermissionRecordController struct {
	PermissionRecordUsecase domain.PermissionRecordUsecase
	Env                     *bootstrap.Env
}

func (tc *PermissionRecordController) Create(c *gin.Context) {
	var task domain.PermissionRecord
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

	err = tc.PermissionRecordUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermissionRecord created successfully",
	})
}

func (tc *PermissionRecordController) Update(c *gin.Context) {
	var task domain.PermissionRecord
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

	err = tc.PermissionRecordUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermissionRecord update successfully",
	})
}

func (tc *PermissionRecordController) Delete(c *gin.Context) {
	var task domain.PermissionRecord
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

	err = tc.PermissionRecordUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermissionRecord update successfully",
	})
}

func (lc *PermissionRecordController) FetchByID(c *gin.Context) {
	PermissionRecordID := c.Query("id")
	if PermissionRecordID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	PermissionRecord, err := lc.PermissionRecordUsecase.FetchByID(c, PermissionRecordID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: PermissionRecordID})
		return
	}

	c.JSON(http.StatusOK, PermissionRecord)
}

func (lc *PermissionRecordController) Fetch(c *gin.Context) {

	PermissionRecord, err := lc.PermissionRecordUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PermissionRecord)
}
