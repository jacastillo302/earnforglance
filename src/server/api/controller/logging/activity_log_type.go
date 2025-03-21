package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/logging"

	"github.com/gin-gonic/gin"
)

type ActivityLogTypeController struct {
	ActivityLogTypeUsecase domain.ActivityLogTypeUsecase
	Env                    *bootstrap.Env
}

func (tc *ActivityLogTypeController) Create(c *gin.Context) {
	var task domain.ActivityLogType
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

	err = tc.ActivityLogTypeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ActivityLogType created successfully",
	})
}

func (tc *ActivityLogTypeController) Update(c *gin.Context) {
	var task domain.ActivityLogType
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

	err = tc.ActivityLogTypeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ActivityLogType update successfully",
	})
}

func (tc *ActivityLogTypeController) Delete(c *gin.Context) {
	var task domain.ActivityLogType
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

	err = tc.ActivityLogTypeUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ActivityLogType update successfully",
	})
}

func (lc *ActivityLogTypeController) FetchByID(c *gin.Context) {
	ActivityLogTypeID := c.Query("id")
	if ActivityLogTypeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ActivityLogType, err := lc.ActivityLogTypeUsecase.FetchByID(c, ActivityLogTypeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ActivityLogTypeID})
		return
	}

	c.JSON(http.StatusOK, ActivityLogType)
}

func (lc *ActivityLogTypeController) Fetch(c *gin.Context) {

	ActivityLogType, err := lc.ActivityLogTypeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ActivityLogType)
}
