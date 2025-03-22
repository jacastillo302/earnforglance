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

type RobotsTxtSettingsController struct {
	RobotsTxtSettingsUsecase domain.RobotsTxtSettingsUsecase
	Env                      *bootstrap.Env
}

func (tc *RobotsTxtSettingsController) Create(c *gin.Context) {
	var task domain.RobotsTxtSettings
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

	err = tc.RobotsTxtSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RobotsTxtSettings created successfully",
	})
}

func (tc *RobotsTxtSettingsController) Update(c *gin.Context) {
	var task domain.RobotsTxtSettings
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

	err = tc.RobotsTxtSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RobotsTxtSettings update successfully",
	})
}

func (tc *RobotsTxtSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.RobotsTxtSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *RobotsTxtSettingsController) FetchByID(c *gin.Context) {
	RobotsTxtSettingsID := c.Query("id")
	if RobotsTxtSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	RobotsTxtSettings, err := lc.RobotsTxtSettingsUsecase.FetchByID(c, RobotsTxtSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: RobotsTxtSettingsID})
		return
	}

	c.JSON(http.StatusOK, RobotsTxtSettings)
}

func (lc *RobotsTxtSettingsController) Fetch(c *gin.Context) {

	RobotsTxtSettings, err := lc.RobotsTxtSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RobotsTxtSettings)
}
