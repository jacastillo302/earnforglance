package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/localization"

	"github.com/gin-gonic/gin"
)

type LocalizationSettingsController struct {
	LocalizationSettingsUsecase domain.LocalizationSettingsUsecase
	Env                         *bootstrap.Env
}

func (tc *LocalizationSettingsController) Create(c *gin.Context) {
	var task domain.LocalizationSettings
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

	err = tc.LocalizationSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "LocalizationSettings created successfully",
	})
}

func (tc *LocalizationSettingsController) Update(c *gin.Context) {
	var task domain.LocalizationSettings
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

	err = tc.LocalizationSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "LocalizationSettings update successfully",
	})
}

func (tc *LocalizationSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.LocalizationSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *LocalizationSettingsController) FetchByID(c *gin.Context) {
	LocalizationSettingsID := c.Query("id")
	if LocalizationSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	LocalizationSettings, err := lc.LocalizationSettingsUsecase.FetchByID(c, LocalizationSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: LocalizationSettingsID})
		return
	}

	c.JSON(http.StatusOK, LocalizationSettings)
}

func (lc *LocalizationSettingsController) Fetch(c *gin.Context) {

	LocalizationSettings, err := lc.LocalizationSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, LocalizationSettings)
}
