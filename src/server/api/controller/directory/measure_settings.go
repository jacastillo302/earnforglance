package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/directory"

	"github.com/gin-gonic/gin"
)

type MeasureSettingsController struct {
	MeasureSettingsUsecase domain.MeasureSettingsUsecase
	Env                    *bootstrap.Env
}

func (tc *MeasureSettingsController) Create(c *gin.Context) {
	var task domain.MeasureSettings
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

	err = tc.MeasureSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MeasureSettings created successfully",
	})
}

func (tc *MeasureSettingsController) Update(c *gin.Context) {
	var task domain.MeasureSettings
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

	err = tc.MeasureSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MeasureSettings update successfully",
	})
}

func (tc *MeasureSettingsController) Delete(c *gin.Context) {
	var task domain.MeasureSettings
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

	err = tc.MeasureSettingsUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MeasureSettings update successfully",
	})
}

func (lc *MeasureSettingsController) FetchByID(c *gin.Context) {
	MeasureSettingsID := c.Query("id")
	if MeasureSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	MeasureSettings, err := lc.MeasureSettingsUsecase.FetchByID(c, MeasureSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: MeasureSettingsID})
		return
	}

	c.JSON(http.StatusOK, MeasureSettings)
}

func (lc *MeasureSettingsController) Fetch(c *gin.Context) {

	MeasureSettings, err := lc.MeasureSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MeasureSettings)
}
