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

type MessageTemplatesSettingsController struct {
	MessageTemplatesSettingsUsecase domain.MessageTemplatesSettingsUsecase
	Env                             *bootstrap.Env
}

func (tc *MessageTemplatesSettingsController) Create(c *gin.Context) {
	var task domain.MessageTemplatesSettings
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

	err = tc.MessageTemplatesSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MessageTemplatesSettings created successfully",
	})
}

func (tc *MessageTemplatesSettingsController) Update(c *gin.Context) {
	var task domain.MessageTemplatesSettings
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

	err = tc.MessageTemplatesSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MessageTemplatesSettings update successfully",
	})
}

func (tc *MessageTemplatesSettingsController) Delete(c *gin.Context) {
	var task domain.MessageTemplatesSettings
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

	err = tc.MessageTemplatesSettingsUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MessageTemplatesSettings update successfully",
	})
}

func (lc *MessageTemplatesSettingsController) FetchByID(c *gin.Context) {
	MessageTemplatesSettingsID := c.Query("id")
	if MessageTemplatesSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	MessageTemplatesSettings, err := lc.MessageTemplatesSettingsUsecase.FetchByID(c, MessageTemplatesSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: MessageTemplatesSettingsID})
		return
	}

	c.JSON(http.StatusOK, MessageTemplatesSettings)
}

func (lc *MessageTemplatesSettingsController) Fetch(c *gin.Context) {

	MessageTemplatesSettings, err := lc.MessageTemplatesSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageTemplatesSettings)
}
