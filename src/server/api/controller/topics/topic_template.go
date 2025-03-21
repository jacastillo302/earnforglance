package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/topics"

	"github.com/gin-gonic/gin"
)

type TopicTemplateController struct {
	TopicTemplateUsecase domain.TopicTemplateUsecase
	Env                  *bootstrap.Env
}

func (tc *TopicTemplateController) Create(c *gin.Context) {
	var task domain.TopicTemplate
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

	err = tc.TopicTemplateUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "TopicTemplate created successfully",
	})
}

func (tc *TopicTemplateController) Update(c *gin.Context) {
	var task domain.TopicTemplate
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

	err = tc.TopicTemplateUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "TopicTemplate update successfully",
	})
}

func (tc *TopicTemplateController) Delete(c *gin.Context) {
	var task domain.TopicTemplate
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

	err = tc.TopicTemplateUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "TopicTemplate update successfully",
	})
}

func (lc *TopicTemplateController) FetchByID(c *gin.Context) {
	TopicTemplateID := c.Query("id")
	if TopicTemplateID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	TopicTemplate, err := lc.TopicTemplateUsecase.FetchByID(c, TopicTemplateID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: TopicTemplateID})
		return
	}

	c.JSON(http.StatusOK, TopicTemplate)
}

func (lc *TopicTemplateController) Fetch(c *gin.Context) {

	TopicTemplate, err := lc.TopicTemplateUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, TopicTemplate)
}
