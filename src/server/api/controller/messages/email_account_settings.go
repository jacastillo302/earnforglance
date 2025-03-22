package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/messages"

	"github.com/gin-gonic/gin"
)

type EmailAccountSettingsController struct {
	EmailAccountSettingsUsecase domain.EmailAccountSettingsUsecase
	Env                         *bootstrap.Env
}

func (tc *EmailAccountSettingsController) Create(c *gin.Context) {
	var task domain.EmailAccountSettings
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

	err = tc.EmailAccountSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "EmailAccountSettings created successfully",
	})
}

func (tc *EmailAccountSettingsController) Update(c *gin.Context) {
	var task domain.EmailAccountSettings
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

	err = tc.EmailAccountSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "EmailAccountSettings update successfully",
	})
}

func (tc *EmailAccountSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.EmailAccountSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *EmailAccountSettingsController) FetchByID(c *gin.Context) {
	EmailAccountSettingsID := c.Query("id")
	if EmailAccountSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	EmailAccountSettings, err := lc.EmailAccountSettingsUsecase.FetchByID(c, EmailAccountSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: EmailAccountSettingsID})
		return
	}

	c.JSON(http.StatusOK, EmailAccountSettings)
}

func (lc *EmailAccountSettingsController) Fetch(c *gin.Context) {

	EmailAccountSettings, err := lc.EmailAccountSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, EmailAccountSettings)
}
