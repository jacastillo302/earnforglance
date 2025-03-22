package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type DisplayDefaultMenuItemSettingsController struct {
	DisplayDefaultMenuItemSettingsUsecase common.DisplayDefaultMenuItemSettingsUsecase
	Env                                   *bootstrap.Env
}

func (tc *DisplayDefaultMenuItemSettingsController) Create(c *gin.Context) {
	var task common.DisplayDefaultMenuItemSettings
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

	err = tc.DisplayDefaultMenuItemSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DisplayDefaultMenuItemSettings created successfully",
	})
}

func (tc *DisplayDefaultMenuItemSettingsController) Update(c *gin.Context) {
	var task common.DisplayDefaultMenuItemSettings
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

	err = tc.DisplayDefaultMenuItemSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DisplayDefaultMenuItemSettings update successfully",
	})
}

func (tc *DisplayDefaultMenuItemSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.DisplayDefaultMenuItemSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *DisplayDefaultMenuItemSettingsController) FetchByID(c *gin.Context) {
	DisplayDefaultMenuItemSettingsID := c.Query("id")
	if DisplayDefaultMenuItemSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	DisplayDefaultMenuItemSettings, err := lc.DisplayDefaultMenuItemSettingsUsecase.FetchByID(c, DisplayDefaultMenuItemSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: DisplayDefaultMenuItemSettingsID})
		return
	}

	c.JSON(http.StatusOK, DisplayDefaultMenuItemSettings)
}

func (lc *DisplayDefaultMenuItemSettingsController) Fetch(c *gin.Context) {

	DisplayDefaultMenuItemSettings, err := lc.DisplayDefaultMenuItemSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DisplayDefaultMenuItemSettings)
}
