package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/customers"

	"github.com/gin-gonic/gin"
)

type MultifactorAuthenticationSettingsController struct {
	MultifactorAuthenticationSettingsUsecase domain.MultiFactorAuthenticationSettingsUsecase
	Env                                      *bootstrap.Env
}

func (tc *MultifactorAuthenticationSettingsController) CreateMany(c *gin.Context) {
	var task []domain.MultiFactorAuthenticationSettings
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

	err = tc.MultifactorAuthenticationSettingsUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MultifactorAuthenticationSettings created successfully",
	})
}

func (tc *MultifactorAuthenticationSettingsController) Create(c *gin.Context) {
	var task domain.MultiFactorAuthenticationSettings
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

	err = tc.MultifactorAuthenticationSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MultifactorAuthenticationSettings created successfully",
	})
}

func (tc *MultifactorAuthenticationSettingsController) Update(c *gin.Context) {
	var task domain.MultiFactorAuthenticationSettings
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

	err = tc.MultifactorAuthenticationSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MultifactorAuthenticationSettings updated successfully",
	})
}

func (tc *MultifactorAuthenticationSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.MultifactorAuthenticationSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *MultifactorAuthenticationSettingsController) FetchByID(c *gin.Context) {
	MultifactorAuthenticationSettingsID := c.Query("id")
	if MultifactorAuthenticationSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	MultifactorAuthenticationSettings, err := lc.MultifactorAuthenticationSettingsUsecase.FetchByID(c, MultifactorAuthenticationSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: MultifactorAuthenticationSettingsID})
		return
	}

	c.JSON(http.StatusOK, MultifactorAuthenticationSettings)
}

func (lc *MultifactorAuthenticationSettingsController) Fetch(c *gin.Context) {
	MultifactorAuthenticationSettings, err := lc.MultifactorAuthenticationSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MultifactorAuthenticationSettings)
}
