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

type SecuritySettingsController struct {
	SecuritySettingsUsecase domain.SecuritySettingsUsecase
	Env                     *bootstrap.Env
}

func (tc *SecuritySettingsController) Create(c *gin.Context) {
	var task domain.SecuritySettings
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

	err = tc.SecuritySettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SecuritySettings created successfully",
	})
}

func (tc *SecuritySettingsController) Update(c *gin.Context) {
	var task domain.SecuritySettings
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

	err = tc.SecuritySettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SecuritySettings update successfully",
	})
}

func (tc *SecuritySettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.SecuritySettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *SecuritySettingsController) FetchByID(c *gin.Context) {
	SecuritySettingsID := c.Query("id")
	if SecuritySettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	SecuritySettings, err := lc.SecuritySettingsUsecase.FetchByID(c, SecuritySettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: SecuritySettingsID})
		return
	}

	c.JSON(http.StatusOK, SecuritySettings)
}

func (lc *SecuritySettingsController) Fetch(c *gin.Context) {

	SecuritySettings, err := lc.SecuritySettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SecuritySettings)
}
