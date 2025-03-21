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

type MessageTemplateController struct {
	MessageTemplateUsecase domain.MessageTemplateUsecase
	Env                    *bootstrap.Env
}

func (tc *MessageTemplateController) Create(c *gin.Context) {
	var task domain.MessageTemplate
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

	err = tc.MessageTemplateUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MessageTemplate created successfully",
	})
}

func (tc *MessageTemplateController) Update(c *gin.Context) {
	var task domain.MessageTemplate
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

	err = tc.MessageTemplateUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MessageTemplate update successfully",
	})
}

func (tc *MessageTemplateController) Delete(c *gin.Context) {
	var task domain.MessageTemplate
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

	err = tc.MessageTemplateUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MessageTemplate update successfully",
	})
}

func (lc *MessageTemplateController) FetchByID(c *gin.Context) {
	MessageTemplateID := c.Query("id")
	if MessageTemplateID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	MessageTemplate, err := lc.MessageTemplateUsecase.FetchByID(c, MessageTemplateID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: MessageTemplateID})
		return
	}

	c.JSON(http.StatusOK, MessageTemplate)
}

func (lc *MessageTemplateController) Fetch(c *gin.Context) {

	MessageTemplate, err := lc.MessageTemplateUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageTemplate)
}
