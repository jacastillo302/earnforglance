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

type ExternalAuthenticationSettingsController struct {
	ExternalAuthenticationSettingsUsecase domain.ExternalAuthenticationSettingsUsecase
	Env                                   *bootstrap.Env
}

func (tc *ExternalAuthenticationSettingsController) CreateMany(c *gin.Context) {
	var task []domain.ExternalAuthenticationSettings
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

	err = tc.ExternalAuthenticationSettingsUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ExternalAuthenticationSettings created successfully",
	})
}

func (tc *ExternalAuthenticationSettingsController) Create(c *gin.Context) {
	var task domain.ExternalAuthenticationSettings
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

	err = tc.ExternalAuthenticationSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ExternalAuthenticationSettings created successfully",
	})
}

func (tc *ExternalAuthenticationSettingsController) Update(c *gin.Context) {
	var task domain.ExternalAuthenticationSettings
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

	err = tc.ExternalAuthenticationSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ExternalAuthenticationSettings updated successfully",
	})
}

func (tc *ExternalAuthenticationSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ExternalAuthenticationSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ExternalAuthenticationSettingsController) FetchByID(c *gin.Context) {
	ExternalAuthenticationSettingsID := c.Query("id")
	if ExternalAuthenticationSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ExternalAuthenticationSettings, err := lc.ExternalAuthenticationSettingsUsecase.FetchByID(c, ExternalAuthenticationSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ExternalAuthenticationSettingsID})
		return
	}

	c.JSON(http.StatusOK, ExternalAuthenticationSettings)
}

func (lc *ExternalAuthenticationSettingsController) Fetch(c *gin.Context) {
	ExternalAuthenticationSettings, err := lc.ExternalAuthenticationSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ExternalAuthenticationSettings)
}
