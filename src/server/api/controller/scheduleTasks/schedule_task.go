package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/scheduleTasks"

	"github.com/gin-gonic/gin"
)

type ScheduleTaskController struct {
	ScheduleTaskUsecase domain.ScheduleTaskUsecase
	Env                 *bootstrap.Env
}

func (tc *ScheduleTaskController) Create(c *gin.Context) {
	var task domain.ScheduleTask
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

	err = tc.ScheduleTaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ScheduleTask created successfully",
	})
}

func (tc *ScheduleTaskController) Update(c *gin.Context) {
	var task domain.ScheduleTask
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

	err = tc.ScheduleTaskUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ScheduleTask update successfully",
	})
}

func (tc *ScheduleTaskController) Delete(c *gin.Context) {
	var task domain.ScheduleTask
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

	err = tc.ScheduleTaskUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ScheduleTask update successfully",
	})
}

func (lc *ScheduleTaskController) FetchByID(c *gin.Context) {
	ScheduleTaskID := c.Query("id")
	if ScheduleTaskID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ScheduleTask, err := lc.ScheduleTaskUsecase.FetchByID(c, ScheduleTaskID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ScheduleTaskID})
		return
	}

	c.JSON(http.StatusOK, ScheduleTask)
}

func (lc *ScheduleTaskController) Fetch(c *gin.Context) {

	ScheduleTask, err := lc.ScheduleTaskUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ScheduleTask)
}
