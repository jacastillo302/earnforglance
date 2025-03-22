package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type DisplayDefaultFooterItemSettingsController struct {
	DisplayDefaultFooterItemSettingsUsecase common.DisplayDefaultFooterItemSettingsUsecase
	Env                                     *bootstrap.Env
}

func (tc *DisplayDefaultFooterItemSettingsController) Create(c *gin.Context) {
	var task common.DisplayDefaultFooterItemSettings
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

	err = tc.DisplayDefaultFooterItemSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DisplayDefaultFooterItemSettings created successfully",
	})
}

func (tc *DisplayDefaultFooterItemSettingsController) Update(c *gin.Context) {
	var task common.DisplayDefaultFooterItemSettings
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

	err = tc.DisplayDefaultFooterItemSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DisplayDefaultFooterItemSettings update successfully",
	})
}

func (tc *DisplayDefaultFooterItemSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.DisplayDefaultFooterItemSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *DisplayDefaultFooterItemSettingsController) FetchByID(c *gin.Context) {
	DisplayDefaultFooterItemSettingsID := c.Query("id")
	if DisplayDefaultFooterItemSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	DisplayDefaultFooterItemSettings, err := lc.DisplayDefaultFooterItemSettingsUsecase.FetchByID(c, DisplayDefaultFooterItemSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: DisplayDefaultFooterItemSettingsID})
		return
	}

	c.JSON(http.StatusOK, DisplayDefaultFooterItemSettings)
}

func (lc *DisplayDefaultFooterItemSettingsController) Fetch(c *gin.Context) {

	DisplayDefaultFooterItemSettings, err := lc.DisplayDefaultFooterItemSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DisplayDefaultFooterItemSettings)
}
